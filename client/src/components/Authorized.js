import React from 'react'
import {Redirect} from 'react-router-dom'

class Login extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            userID: '',
            password: '',
        }
        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
    }

    handleChange(e) {
        const name = e.target.name
        if (name === 'user-id') {
            this.setState({
                userID: e.target.value,
            })
        }
        if (name === 'password') {
            this.setState({
                password: e.target.value
            })
        }
    }

    handleSubmit(e) {
        e.preventDefault()
        const userID = this.state.userID
        const password = this.state.password
        this.setState({
            userID: '',
            password: ''
        })
        this.props.handleSubmit(userID, password)
    }
    
    render(){
        if (this.props.isLogin){
            return <Redirect to="/" />
        }
        return (
            <div>
                <h1>Login</h1>
                <form method="POST" onSubmit={this.handleSubmit}>
                    <div className="mb-10">
                        <input type="text" name="user-id" value={this.state.userID} onChange={this.handleChange} className="input-text" placeholder="userID" />
                    </div>
                    <div className="mb-10">
                        <input type="password" name="password" value={this.state.password} onChange={this.handleChange} className="input-text" placeholder="password" />
                    </div>
                    <div>
                        <button type="submit" className="btn btn-primary">Login</button>
                    </div>
                </form>
            </div>
        )
    }
}

export default Login