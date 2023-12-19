import Header from "../Header/Header";
import { Routes, Route } from "react-router-dom";
import Home from "../Home/Home";
import "./BaseLayout.scss";
import Movie from "../Movie/Movie";
import AdminPage from "../AdminPage/AdminPage";

const BaseLayout = () => {
  return (
    <div className="BaseLayout">
      <Header />
      <Routes>
        <Route exact path={"/"} element={<Home />} />
        <Route path="/movie/:id" element={<Movie />} />
        <Route exact path={"/admin"} element={<AdminPage />} />
        {/* <Route exact path={"/portfolio"} element={<Portfolio />} /> */}
      </Routes>
    </div>
  );
};

export default BaseLayout;
