<script setup lang="ts">
import UserInfo from '@/components/UserInfo.vue'
import LoginUser from '../components/LoginUser.vue'
import store from '../store'
import type { User } from '@/types/user'
import Toast from 'primevue/toast'
import { useToast } from 'primevue/usetoast'

const toast = useToast()

const onLogout = (user: User) => {
  toast.add({
    severity: 'success',
    summary: `Goodbye, ${user.full_name}!`,
    detail: `You have successfully logged out.`,
    life: 3000
  })
  store.clearUser()
}
</script>

<template>
  <main>
    <Toast />
    <h1 class="green">Welcome to Simple Bank!</h1>
    <UserInfo v-if="store.state.user" :user="store.state.user" @logout="onLogout" />
    <LoginUser v-else />
  </main>
</template>

<style scoped>
h1 {
  font-weight: 500;
  font-size: 2.6rem;
}
</style>
