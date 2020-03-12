import React from 'react'
import axios from 'axios'

class Account extends React.Component {
    constructor(props){
        super(props)

        this.state = {
            isEditID: false,
            isEditName: false,
            isEditMail: false,

            id: '',
            name: '',
            mail: '',
        }

        this.handleClickEdit = this.handleClickEdit.bind(this)
        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
    }

    componentDidMount(){
        this.props.fetchData()
    }

    handleClickEdit(e) {
        const field = e.target.dataset.field
        let value
        let isEdit
        if (field === 'id') {
            value = this.props.user.id
            isEdit = "isEditID"
        }
        if (field === 'name'){
            value = this.props.user.name
            isEdit = "isEditName"
        }
        if (field === 'mail'){
            value = this.props.user.mail
            isEdit = "isEditMail"
        }
        this.setState({
            [field]: value,
            [isEdit]: true
        })
    }

    handleChange(e) {
        const field = e.target.name
        this.setState({
            [field]: e.target.value
        })
    }

    handleSubmit(e) {
        e.preventDefault()
        const field = e.target.dataset.field
        let value
        let isEdit
        if (field === 'id') {
            value = this.state.id
            isEdit = "isEditID"
        }
        if (field === 'name') {
            value = this.state.name
            isEdit = "isEditName"
        }
        if (field === 'mail') {
            value = this.state.mail
            isEdit = "isEditMail"
        }
        this.setState({
            [field]: '',
            [isEdit]: false
        })
        this.props.updateData(field, value)
    }
    
    render(){
        return (
            <div>
                <h1>Profile</h1>
                <div>
                    <AccountRow
                        user={this.props.user}
                        fieldLabel={"UserID"}
                        fieldName={"id"}
                        fieldValue={this.state.id}
                        isEdit={this.state.isEditID}
                        value={this.props.user.id}
                        handleClick={this.handleClickEdit}
                        handleChange={this.handleChange}
                        handleSubmit={this.handleSubmit}
                        isFetching={this.props.isFetching}
                    />
                    <AccountRow
                        user={this.props.user}
                        fieldLabel={"Name"}
                        fieldName={"name"}
                        fieldValue={this.state.name}
                        isEdit={this.state.isEditName}
                        value={this.props.user.name}
                        handleClick={this.handleClickEdit}
                        handleChange={this.handleChange}
                        handleSubmit={this.handleSubmit}
                    />
                    <AccountRow
                        user={this.props.user}
                        fieldLabel={"Mail"}
                        fieldName={"mail"}
                        fieldValue={this.state.mail}
                        isEdit={this.state.isEditMail}
                        value={this.props.user.mail}
                        handleClick={this.handleClickEdit}
                        handleChange={this.handleChange}
                        handleSubmit={this.handleSubmit}
                    />
                </div>
            </div>

        )
    }
}

function AccountRow(props) {
    let value
    if(props.isFetching) {
        value = "NowLoading"
    } else {
        value = props.value
    }

    if (props.fieldName === 'id') {
        return(
            <div>
                <div>
                    <h2 className="h2">{props.fieldLabel}</h2>
                </div>
                <div>
                    <span>{value}</span>
                </div>
            </div>
        )
    }
    if (props.isEdit) {
        return (
            <div>
                <div>
                    <h2 className="h2">{props.fieldLabel}</h2>
                    <button type="submit" data-field={props.fieldName} onClick={props.handleSubmit} className="btn btn-primary ml-10">save</button>
                </div>
                <div>
                    <input type="text" name={props.fieldName} value={props.fieldValue} onChange={props.handleChange} className="edit-text" size="40" />
                </div>
            </div>
        )
    }
    return (
        <div>
            <div>
                <h2 className="h2">{props.fieldLabel}</h2>
                <button onClick={props.handleClick} data-field={props.fieldName} className="btn btn-primary ml-10">edit</button>
            </div>
            <div>
                <span>{value}</span>
            </div>
        </div>
    )
}

export default Account