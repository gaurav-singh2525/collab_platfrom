import api from "./api";

export const executeCode =
	async (
		language,
		code
	) => {

	const response =
		await api.post(
			"/execute",
			{
				language,
				code,
			}
		);

	return response.data;
};