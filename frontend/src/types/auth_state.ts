import type { User } from './user'

export interface AuthState {
  user: User | null
  accessToken: string | null
  refreshToken: string | null
}
