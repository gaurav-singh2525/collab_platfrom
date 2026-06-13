import api from "./api";

export const executeCode =
	async (
		roomID,
		language,
		code
	) => {

		const response =
			await api.post(
				"/execute",
				{
					roomID,
					language,
					code,
				}
			);

		return response.data;
	};