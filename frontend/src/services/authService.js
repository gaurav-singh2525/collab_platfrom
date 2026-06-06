import api from "./api";

export const signup = async (
  email,
  password
) => {

  const response =
    await api.post(
      "/signup",
      {
        email,
        password,
      }
    );

  return response.data;
};

export const login = async (
  email,
  password
) => {

  const response =
    await api.post(
      "/login",
      {
        email,
        password,
      }
    );

  return response.data;
};

export const getCurrentUser =
  async () => {

    const response =
      await api.get("/me");

    return response.data;
  };