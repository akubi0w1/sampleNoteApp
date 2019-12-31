import React from 'react'
import { BrowserRouter, Route, Link, Redirect } from 'react-router-dom'

import axios from 'axios'

import Header from './components/Header'
import Account from './components/Account'
import Login from './components/Authorized'
import Home from './components/Home'
import NoteDefail from './components/Detail'
import NoteEdit from './components/Edit'

const BASE_URL = 'http://localhost:8080'

class App extends React.Component {
    constructor(props){
        super(props)

        this.state = {
            user: {},
            notes: [],
            note: {},
            isLogin: false,
            isFetching: false,
        }

        this.fetchAccountData = this.fetchAccountData.bind(this)
        this.updateAccount = this.updateAccount.bind(this)
        this.requestLogin = this.requestLogin.bind(this)
        this.fetchNotesData = this.fetchNotesData.bind(this)
        this.fetchNoteData = this.fetchNoteData.bind(this)
        this.updateNote = this.updateNote.bind(this)
        this.deleteNote = this.deleteNote.bind(this)
    }

    fetchAccountData(){
        this.setState({
            isFetching: true,
        })

        axios.get(BASE_URL + '/accounts', {
            withCredentials: true,
        })
        .then(res => {
            this.setState({
                isFetching: false,
            })

            const _user = res.data
            this.setState({
                user: _user
            })
        })
        .catch(err => {
            console.log(err)
        })
    }

    updateAccount(field, value){
        axios.put(BASE_URL + '/accounts', {
            [field]: value,
        }, {
            withCredentials: true,
        })
        .then(res => {
            console.log(res)
            this.fetchAccountData()
        })
        .catch(err => {
            console.log(err)
        })
    }

    fetchNotesData(){
        this.setState({
            isFetching: true,
        })

        axios.get(BASE_URL + '/notes', {
            withCredentials: true,
        })
            .then(res => {
                this.setState({
                    isFetching: false,
                })

                const _notes = res.data.notes
                this.setState({
                    notes: _notes
                })
            })
            .catch(err => {
                console.log(err)
            })
    }

    fetchNoteData(id) {
        this.setState({
            isFetching: true,
        })

        axios.get(BASE_URL + '/notes/' + id, {
            withCredentials: true,
        })
        .then(res => {
            console.log(res)

            const _note = res.data
            this.setState({
                note: _note
            })
            this.setState({
                isFetching: false,
            })
        })
        .catch(err => {
            console.log(err)
        })
    }

    updateNote(id, title, content) {
        if(id){
            axios.put(BASE_URL + '/notes/' + id, {
                title: title,
                content: content,
            }, {
                withCredentials: true,
            })
            .then(res => {
                console.log(res)
            })
            .catch(err => {
                console.log(err)
            })
        } else {
            axios.post(BASE_URL + '/notes', {
                title: title,
                content: content,
            }, {
                withCredentials: true,
            })
            .then(res => {
                console.log(res)
            })
            .catch(err => {
                console.log(err)
            })
        }
        this.setState({
            note: {}
        }) 
    }

    deleteNote(id){
        axios.delete( BASE_URL + '/notes/' + id, {
            withCredentials: true,
        })
        .then(res => {
            console.log(res)
            window.location = '/'
        })
        .catch(err => {
            console.log(err)
        })
    }

    requestLogin(userID, password){
        axios.post(BASE_URL + '/login', {
            id: userID,
            password: password,
        }, {
            withCredentials: true,
        })
        .then(res => {
            console.log(res.headers)
            this.setState({
                isLogin: true
            })
        })
        .catch(err => {
            console.log(err)
        })
    }

    render(){
        return(
            <BrowserRouter>
                <Header/>
                <div className="content">
                    <Route
                        path='/users'
                        render={props =>
                            <Users
                                {...props} />} />
                    <Route
                        exact
                        path='/'
                        render={props =>
                            <Home 
                                fetchData={this.fetchNotesData}
                                notes={this.state.notes}
                                isFetching={this.state.isFetching}
                                {...props}
                                />}
                        />
                    <Route
                        exact
                        path='/notes/item/:id'
                        render={props => 
                            <NoteDefail
                                fetchData={this.fetchNoteData}
                                deleteData={this.deleteNote}
                                note={this.state.note}
                                isFetching={this.state.isFetching}
                                {...props}
                            />
                        }
                    />
                    <Route
                        exact
                        path='/notes/item/:id/edit'
                        render={props =>
                            <NoteEdit
                                fetchData={this.fetchNoteData}
                                note={this.state.note}
                                isFetching={this.state.isFetching}
                                handleSubmit={this.updateNote}
                                {...props}
                            />
                        }
                    />
                    <Route
                        exact
                        path='/notes/new'
                        render={ props =>
                            <NoteEdit
                                fetchData={this.fetchNoteData}
                                note={this.state.note}
                                isFetching={this.state.isFetching}
                                handleSubmit={this.updateNote}
                                {...props}
                            />
                        }
                        />
                    <Route
                        path='/login'
                        render={props => 
                            <Login
                                handleSubmit={this.requestLogin}
                                isLogin={this.state.isLogin}
                                {...props}
                            />
                        }
                    />
                    <Route
                        exact
                        path='/account'
                        render={props =>
                            <Account
                                isFetching={this.state.isFetching}
                                fetchData={this.fetchAccountData}
                                updateData={this.updateAccount}
                                user={this.state.user}
                                {...props}
                            />}
                    />
                </div>
            </BrowserRouter>
        )
    }
}

class Users extends React.Component {
    constructor(props){
        super(props)

        this.state = {
            isFetching: true,
            users: [],
        }
        
        this.componentDidMount = this.componentDidMount.bind(this)
        this.fetchData = this.fetchData.bind(this)
        
    }

    componentDidMount = () => {
        this.fetchData()
    }

    fetchData = () => {
        axios.get(BASE_URL + '/users')
            .then(res => {
                const users = res.data.users
                this.setState({
                    isFetching: false,
                    users: users,
                })
            })
            .catch(err => {
                console.log(err)
            });
    }
    render(){
        return(
            <div>
                <h1>Users</h1>
                {
                    this.state.isFetching
                        ? <p>Now loading...</p>
                        : <div>
                            {this.state.users.map(user => (
                                <li key={user.id}>
                                    {user.id}
                                </li>
                            ))}
                          </div>
                }
            </div>
            
        )
    }
}

export default App