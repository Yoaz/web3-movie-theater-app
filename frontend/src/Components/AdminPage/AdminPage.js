import { useEffect, useState } from "react";
import { getAllMovies } from "../Utils/Fetch Data/fetchData";
import { updateMovieByID, deleteMovieByID } from "../Utils/Post Data/postData";
import { formatDuration } from "../Utils/Helpers";
import "./AdminPage.scss";

const AdminPage = () => {
  const [movies, setMovies] = useState(null);
  const [movieUpdates, setMovieUpdates] = useState({});

  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await getAllMovies();
        setMovies(data.data.data);
      } catch (error) {
        console.error("Error fetching movie:", error);
      }
    };

    fetchData();
  }, [movies]);

  // Handling update movie event
  const handleUpdateMovieField = (id, field, value) => {
    setMovieUpdates((prevUpdates) => ({
      ...prevUpdates,
      [id]: { ...prevUpdates[id], [field]: value },
    }));
  };

  const handleUpdateMovie = async (id) => {
    const updatedFields = movieUpdates[id];
    if (!updatedFields) {
      // No updates for this movie
      return;
    }

    try {
      await updateMovieByID(id, updatedFields);

      // Refetch the movies after the update
      const updatedMovies = await getAllMovies();
      setMovies(updatedMovies.data.data);

      // Clear the updates for this movie
      setMovieUpdates((prevUpdates) => ({
        ...prevUpdates,
        [id]: undefined,
      }));
    } catch (error) {
      console.error("Error updating movie:", error);
    }
  };

  // Handling delete movie event
  const handleDeleteMovie = async (id) => {
    try {
      await deleteMovieByID(id);
      // Refetch the movies after the update
      const updatedMovies = await getAllMovies();
      setMovies(updatedMovies.data.data);
    } catch (error) {
      console.error("Error deleting movie:", error);
    }
  };

  if (!movies) {
    // Loading spinner while waiting for fetch
    return <div className="loading-spinner"></div>;
  }

  return (
    <div className="movies-container">
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Description</th>
            <th>Duration</th>
            <th>Date</th>
            <th>Active</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          {movies.map((movie) => (
            <tr key={movie._id}>
              <td>{movie._id}</td>
              <td>
                <input
                  type="text"
                  value={movieUpdates[movie._id]?.name || movie.name}
                  onChange={(e) =>
                    handleUpdateMovieField(movie._id, "name", e.target.value)
                  }
                />
              </td>
              <td>
                <textarea
                  value={
                    movieUpdates[movie._id]?.description || movie.description
                  }
                  onChange={(e) =>
                    handleUpdateMovieField(
                      movie._id,
                      "description",
                      e.target.value
                    )
                  }
                />
              </td>
              <td>{formatDuration(movie.duration)}</td>
              <td>{movie.date}</td>
              <td>
                <select
                  value={movieUpdates[movie._id]?.active || movie.active}
                  onChange={(e) =>
                    handleUpdateMovieField(movie._id, "active", e.target.value)
                  }
                >
                  <option value={true}>True</option>
                  <option value={false}>False</option>
                </select>
              </td>
              <td>
                <button onClick={() => handleUpdateMovie(movie._id)}>
                  Update
                </button>
                <button onClick={() => handleDeleteMovie(movie._id)}>
                  Delete
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default AdminPage;
