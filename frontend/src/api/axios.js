import axios from "axios";

const instance = axios.create({
	baseURL: "/api/v1",
});

instance.interceptors.request.use((config) => {
	const token = localStorage.getItem("access_token");
	if (token) {
		config.headers.Authorization = `Bearer ${token}`;
	}
	return config;
});

export default instance;
