import { reactive, readonly } from 'vue'
import type { AuthState } from './types/auth_state'
import type { User } from './types/user'

const state = reactive<AuthState>({
  user: null,
  accessToken: null,
  refreshToken: null
})

function setUser(user: User, accessToken: string, refreshToken: string) {
  state.user = user
  state.accessToken = accessToken
  state.refreshToken = refreshToken
}

function clearUser() {
  state.user = null
  state.accessToken = null
  state.refreshToken = null
}

export default {
  state: readonly(state),
  setUser,
  clearUser
}
