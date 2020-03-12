import React from 'react'

class NoteEdit extends React.Component {
    constructor(props){
        super(props)

        this.state = {
            id: '',
            title: '',
            content: '',
            isInit: false,
        }

        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
    }

    componentDidMount(){
        if (this.props.match.params.id) {
            this.props.fetchData(this.props.match.params.id)
            this.setState({
                id: this.props.match.params.id,
            })
        }
    }

    componentDidUpdate(){
        if(this.props.note.title){
            if (!this.state.isInit) {
                this.setState({
                    title: this.props.note.title,
                    content: this.props.note.content,
                    isInit: true
                })
            }
        }
    }

    handleChange(e){
        const field = e.target.name
        this.setState({
            [field]: e.target.value
        })
    }

    handleSubmit(e){
        e.preventDefault()
        const id = this.state.id
        const title = this.state.title
        const content = this.state.content
        this.setState({
            id: '',
            title: '',
            content: '',
        })
        this.props.handleSubmit(id, title, content)
    }

    render(){
        return(
            <div>
                {
                    this.props.isFetching
                    ? <p>Now Loading...</p>
                    : <form onSubmit={this.handleSubmit}>
                        <div className="mb-10">
                            <input name="title" type="text" value={this.state.title} onChange={this.handleChange} className="input-title" size="40" placeholder="note title..." />
                            <button type="submit" className="btn btn-primary ml-10">save</button>
                        </div>
                        <textarea value={this.state.content} name="content" onChange={this.handleChange} className="input-content" placeholder="note content..."></textarea>
                    </form>
                }
            </div>
        )
    }
}

export default NoteEdit