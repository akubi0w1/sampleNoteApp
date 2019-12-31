import React from 'react'
import { Link } from 'react-router-dom'

class Home extends React.Component{
    constructor(props) {
        super(props)
    }

    componentDidMount(){
        console.log("did maou")
        this.props.fetchData()
    }

    render() {
        return (
            <div>
                <NoteList
                    isFetching={this.props.isFetching}
                    notes={this.props.notes}
                    />
            </div>
        )
    }
}

function NoteList(props) {
    return (
        <div>
            <h1>NoteList</h1>
            {
                props.isFetching
                    ? <p>Now Loading...</p>
                    : props.notes.map(note => (
                        <NoteRow key={note.id} note={note}/>
                    ))
                }
        </div>
    )

}

function NoteRow(props) {
    return (
        <div className="note-list-item">
            <h2><Link to={`/notes/item/${props.note.id}`}>{props.note.title}</Link></h2>
        </div>
    )
}

export default Home