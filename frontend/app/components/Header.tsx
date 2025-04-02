import {Link} from "@remix-run/react";
import '../styles/Header.scss';

const Header = () => {
  return (
    <header className="header">
      <h1 className="title">Quiz</h1>
      <nav className="nav-links">
        <ul>
          <li><Link to="/">Home</Link></li>
          <li><Link to="/table">Table</Link></li>
          <li><Link to="/list">List</Link></li>
          <li><Link to="/quiz">Quiz</Link></li>
        </ul>
      </nav>
    </header>
  );
};

export default Header;
