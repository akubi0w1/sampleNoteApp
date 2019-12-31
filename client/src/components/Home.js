import React from 'react'
import { Link } from 'react-router-dom'

class Home extends React.Component{
    constructor(props) {
        super(props)
    }

    componentDidMount(){
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
            <div>
                <h1 className="h1">NoteList</h1>
                <Link to='/notes/new'><button className="btn btn-primary ml-10">new</button></Link>
            </div>
            {
                props.isFetching
                    ? <p>Now Loading...</p>
                    : <div>
                        {
                            !props.notes
                            ? <p>No note</p>
                            : props.notes.map(note => (
                                <NoteRow key={note.id} note={note} />
                            ))
                        }
                    </div>
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