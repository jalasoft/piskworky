
(function() {

    let playground;

    const whoStartsCheckbox = document.getElementById("whostarts");

    document.getElementById("start_btn").addEventListener("click", e => {
        e.preventDefault();
        startGame(whoStartsCheckbox.checked);
    });

    document.getElementById("next_pos_btn").addEventListener('click', e=> {
        e.preventDefault()
        fetch('/game/next')
        .then(r => r.json())
        .then(pos => {
            for (let p of pos) {
                console.log(p);   
                playground.markField('E', 'next-field', p.row, p.column)
            }
        })
    });

    function startGame(meStart) {
        const body = {
            "rows": 10,
            "columns": 10
        }

        fetch('/game/start', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(body) 
        }).then(resp => resp.json())
        .then(json => {
            initPlayground(json.rows, json.columns);
        })
        .then(r => {
            if (!meStart) {
                return letAiMove()
            }
        })
        .catch(err => console.log(err))
    }

    function initPlayground(rows, columns) {

        playground = new Playground(rows, columns, (row, column) => {
            humanMove(row, column)
            .then(r => letAiMove())
            .catch(err => console.log(err));
        });

        document.getElementById("content").appendChild(playground.elm);
    }

    function humanMove(row, column) {
        const payload = {
            row: row,
            column: column
        }

        return fetch("/game/oponent/move", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(payload)  
        })
        .then(resp => resp.json())
        .then(json => console.log(`Oponent move [${row}, ${column}] with ${json.status}`));
    }

    function letAiMove() {
        return fetch("/game/ai/move")
        .then(r => r.json())
        .then(json => playground.markAIField(json.position.row, json.positionm.column));
    }
})()


function Playground(rows, columns, fieldClickHandler) {
    const [elm, fields ] = buildPlayground(rows, columns, (row, column) => {
        this.markField('H', 'selected-field-human', row, column);
        fieldClickHandler(row, column);
    })

    this.fields = fields;
    this.elm = elm;

    function buildPlayground(rows, columns, handler) {

        const fields = new Map()

        const playgroundElm = document.createElement("div");
        playgroundElm.classList.add("playground");

        for(let row=0; row<rows; row++) {
            let rowElm = document.createElement("div");
            rowElm.classList.add("row");

            const rowMap = new Map()
            fields.set(row, rowMap);        

            for(let column=0;column<columns;column++) {
                let fieldElm = document.createElement("div");
                fieldElm.classList.add("field");
                fieldElm.addEventListener('click', e=>handler(row, column));

                rowElm.appendChild(fieldElm);
            
                rowMap.set(column, {
                    elm: fieldElm,
                    status: 'E',
                    row: row,
                    column: column
                });
            }

            playgroundElm.appendChild(rowElm);
        }

        return [ playgroundElm, fields ];
    }

    
}

Playground.prototype.markAIField = function(row, column) {
    console.log(`AI move: [${row}, ${column}]`)
    this.markField('M', 'selected-field-machine', row, column);
}

Playground.prototype.markField = function(type, styleClass, row, column) {
        const descriptor = this.fields.get(row).get(column);
        
        if (descriptor.status !== 'E') {
            console.log("Field at position [" + descriptor.row + ", " + descriptor.column + "] is already occupied ");
            return;
        }

        descriptor.elm.classList.add(styleClass)
        descriptor.status = type
    }

Playground.prototype.oponnentWon = function() {
    for(let row of this.fields) {
        for (let field in this.fields.get(row[0])) {
            field.elm.classList.add("human-won");            
        }
    }
}

Playground.prototype.aiWon = function() {
    for(let row of this.fields) {
        for (let field in this.fields.get(row[0])) {
            field.elm.classList.add("machine-won");            
        }
    }
}




