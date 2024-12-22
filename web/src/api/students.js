import axios from 'axios';

export const getStudents = (number) => {
    return axios.get(`/api/students`, { params: { number } });
};

export const addStudent = (student) => {
    return axios.post(`/api/students`, student);
};

export const deleteStudent = (studentId) => {
    return axios.delete(`api/students/${studentId}`);
};
