import axios from 'axios';

export const login = (number, password) => {
    return axios.post(`/api/login`, { number, password });
};

export const changePassword = (number, newPassword) => {
    return axios.put(`/api/password`, {number, newPassword });
};
