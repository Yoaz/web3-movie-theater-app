import { info } from "../../Info";
import { Link, useLocation } from "react-router-dom";
import style from "./Header.scss";
import { useState } from "react";
import ConnectButton from "../ConnectButton/ConnectButton";

const Header = () => {
  const location = useLocation();
  const [active, setActive] = useState(
    location.pathname === "/"
      ? "home"
      : location.pathname.slice(1, location.pathname.length)
  );
  return (
    <nav className="Header">
      <Link
        className={info.Company.active === active ? style.active : undefined}
        to={"/"}
      >
        <h1>{info.Company.name}</h1>
      </Link>
      <ul>
        {info["links"].map((link, key) => (
          <Link
            className={link.active === active ? style.active : undefined}
            to={link.to}
            key={key}
            onClick={() => setActive()}
          >
            {!link.type && (
              <li style={{ paddingBottom: "0.5rem" }}>{link.name}</li>
            )}
          </Link>
        ))}

        <li>
          <ConnectButton />
        </li>
      </ul>
    </nav>
  );
};

export default Header;
