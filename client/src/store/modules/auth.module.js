import authApiService from "@/services/auth.api.service";

const authModule = {
    namespaced: true,
    state: {
        token:"",
        registerUser:{}
    },
    getters:{
        token: (state) => state.token,
        registerUser: (state) => state.registerUser,
    },
    mutations: {
        SET_TOKEN(state, token){
            state.token = token;
        },
        SET_REGISTER_USER(state, registerUser){
            state.registerUser = registerUser;
        }
    },
    actions: {
        async login({commit}, {email, password}){
            const response = await authApiService.login(email, password);
            if(response){
                commit('SET_TOKEN', response?.data?.data || "");
                console.log("token: ", response?.data?.data);
            }else {
                commit('SET_TOKEN',"");
            }
        },
        async register({commit}, registerForm){
            const response = await authApiService.register(registerForm);
            if(response){
                commit('SET_REGISTER_USER', response?.data?.data || {});
                console.log("Register User: ", response?.data?.data);
            }else {
                commit('SET_REGISTER_USER', {});
            }
        }
    },
};

export default authModule;