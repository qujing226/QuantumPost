import axios from 'axios';

const API_URL = 'http://localhost:8080';

export const getLabs = (number) => {
    return axios.get(`/api/labs`,{ params: { number } });
};

export const updateLabStatus = (labId, status) => {
    return axios.put(`${API_URL}/labs/${labId}/status`, { status });
};
