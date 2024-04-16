import {
	ActionTree,
	createStore,
	MutationTree,
	Store as VuexStore,
} from 'vuex';
import { state, IAppState } from './state';


const mutations: MutationTree<IAppState> = {
};

const actions: ActionTree<IAppState, IAppState> = {
};

const store = createStore<IAppState>({
	state: state,
	mutations: mutations,
	actions: actions,
});
export default store;

// Allow usage of this.$store in components
declare module '@vue/runtime-core' {
    interface ComponentCustomProperties  {
        $store: VuexStore
    }
}
