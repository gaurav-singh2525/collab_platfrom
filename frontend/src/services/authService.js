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