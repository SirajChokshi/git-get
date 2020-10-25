import React from 'react';
import './Nav.scss'
import { Link } from '@reach/router'
import { FaGlasses } from 'react-icons/fa'

const NavLink = (props: any) => (
    <Link
      {...props}
      getProps={({ isCurrent }) => {
        return {
          className: isCurrent ? "active" : ""
        };
      }}
    />
  );

const Nav = () => {
    return (
        <nav>
            <Link to="/" id="logo"><FaGlasses style={{marginRight: "0.4em", marginTop: "0.15em"}} /> GitGet</Link>
            <div>
                {/* <NavLink to="/page">Page</NavLink>
                <NavLink to="/page">Page 2</NavLink> */}
                <NavLink to="/about">About</NavLink>
            </div>
        </nav>
    )
}

export default Nav;