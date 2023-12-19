/* ------------------------------------ Fetch Data Hooks --------------------------------------*/

// Fetching a movie by Name
const getMovieByID = async (id) => {
  try {
    //Setup request options
    const requestOptions = {
      method: "GET",
      mode: "cors",
      redirect: "follow",
      headers: {
        // "content-type": "application/json",
        "content-type": "application/x-www-form-urlencoded",
      },
    };

    // TODO: Finish fetchUrl after establishing auto user in API
    const baseURL = `${process.env.REACT_APP_BASE_API_URL}${process.env.REACT_APP_GET_ONE_MOVIE_BASE}/${id}`;

    const response = await fetch(baseURL, requestOptions);
    const data = await response.json();

    return data;
  } catch (error) {
    console.log("error", error);
  }
};

// Fetch all movies
const getAllMovies = async () => {
  try {
    //Setup request options
    const requestOptions = {
      method: "GET",
      mode: "cors",
      redirect: "follow",
      headers: {
        // "content-type": "application/json",
        "content-type": "application/x-www-form-urlencoded",
      },
    };

    // TODO: Finish fetchUrl after establishing auto user in API
    const baseURL = `${process.env.REACT_APP_BASE_API_URL}${process.env.REACT_APP_GET_ALL_MOVIES}`;

    const response = await fetch(baseURL, requestOptions);
    const data = await response.json();

    return data;
  } catch (error) {
    console.log("error", error);
  }
};

// Fetch all schedules
const getAllSchedules = async () => {
  try {
    //Setup request options
    const requestOptions = {
      method: "GET",
      mode: "cors",
      redirect: "follow",
      headers: {
        // "content-type": "application/json",
        "content-type": "application/x-www-form-urlencoded",
      },
    };

    // TODO: Finish fetchUrl after establishing auto user in API
    const baseURL = `${process.env.REACT_APP_BASE_API_URL}${process.env.REACT_APP_GET_SCHEDULES_RECORDS}`;

    const response = await fetch(baseURL, requestOptions);
    const data = await response.json();

    return data;
  } catch (error) {
    console.log("error", error);
  }
};

export { getMovieByID, getAllMovies, getAllSchedules };
