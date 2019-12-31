import React from 'react'
import { Link } from 'react-router-dom'


function Header() {
    return (
        <div className="header">
            <div className="back"><a href="javascript:history.back()">＜</a></div>
            <label className="logo"><Link className="link" to='/'>SimpleNote</Link></label>
            <div className="icon"><Link to='/account' className="link" /></div>
        </div>
    )
}

export default Header