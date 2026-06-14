import api from "./api";

export const executeCode =
	async (
		roomID,
		language,
		code,
		input
	) => {

		const response =
			await api.post(
				"/execute",
				{
					roomID,
					language,
					code,
					input,
				}
			);

		return response.data;
	};