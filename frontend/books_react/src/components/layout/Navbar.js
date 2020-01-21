import React from "react";

const Navbar = () => {
  return (
    <div>
      <nav className="main-nav">
        <ul>
          <li>
            <a href="#home">
              <h1>Nossos Livros</h1>
            </a>
          </li>
          <li style={{ float: "right" }}>
            <a class="active" href="#">
              <i class="fas fa-user"></i>
            </a>
          </li>
        </ul>
      </nav>
    </div>
  );
};

export default Navbar;
