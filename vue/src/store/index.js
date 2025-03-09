import { createStore } from 'vuex';

export default createStore({
    state: {
        number: '',
        name:''
    },
    mutations: {
        setNumber(state, number) {
            state.number = number;
        },
        setName(state, name) {
            state.name = name;
        }
    },
    actions: {
        setNumber({ commit }, number) {
            commit('setNumber', number);
        },
        setName({ commit }, name){
            commit('setName',name);
        }
    },
    getters: {
        getNumber: state => state.number,
        getName: state => state.name
    }
});
