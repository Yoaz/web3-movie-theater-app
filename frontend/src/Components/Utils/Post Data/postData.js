/* ------------------------------------ Post Data Hooks --------------------------------------*/

const deleteMovieByID = async (id) => {
  try {
    const baseURL = `${process.env.REACT_APP_BASE_API_URL}${process.env.REACT_APP_GET_ONE_MOVIE_BASE}/${id}`;

    // Request options
    const requestOptions = {
      method: "POST",
      mode: "cors",
      redirect: "follow",
      headers: {
        "content-type": "application/x-www-form-urlencoded",
      },
    };

    // Send DELETE request to API
    const response = await fetch(baseURL, requestOptions);

    if (!response.ok) {
      throw new Error(`Failed to delete movie. Status: ${response.status}`);
    }

    return { success: true };
  } catch (error) {
    console.error("Error deleting movie:", error);
    throw error; // Rethrow the error to handle it in the calling code
  }
};

// POST movie update to update
const updateMovieByID = async (id, requestBody) => {
  try {
    // Base API URL
    const baseURL = `${process.env.REACT_APP_BASE_API_URL}${process.env.REACT_APP_GET_ONE_MOVIE_BASE}/${id}`;

    // Request options
    const requestOptions = {
      method: "PATCH",
      mode: "cors",
      headers: {
        "content-type": "application/x-www-form-urlencoded",
      },
      body: JSON.stringify(requestBody),
    };

    // Send PUT request to API
    const response = await fetch(baseURL, requestOptions);

    if (!response.ok) {
      // Handle errors here
      throw new Error(`Failed to update movie. Status: ${response.status}`);
    }

    // Return the updated movie data
    return await response.json();
  } catch (error) {
    console.error("Error updating movie:", error);
    throw error; // Rethrow the error to handle it in the calling code
  }
};

export { updateMovieByID, deleteMovieByID };
