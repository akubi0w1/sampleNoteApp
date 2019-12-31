import React from 'react'
import {Link} from 'react-router-dom'

class NoteDetail extends React.Component {
    constructor(props){
        super(props)
        this.state = {
            id: props.match.params.id
        }
    }

    componentDidMount(){
        this.props.fetchData(this.state.id)
    }

    render(){
        return(
            <div>
                {
                    this.props.isFetching 
                    ? <p>Now Loading...</p>
                    : <div>
                        <h1 className="h1">{this.props.note.title}</h1>
                        <Link to={`/notes/item/${this.props.note.id}/edit`}><button className="btn btn-primary ml-10" data-id={this.props.note.id}>edit</button></Link>
                        <p>{this.props.note.content}</p>
                    </div>
                }
            </div>
        )
    }
}

export default NoteDetail