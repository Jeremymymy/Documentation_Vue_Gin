<template>
  <router-view />
</template>

<script>
import { defineComponent } from 'vue'
import axios from 'axios';
axios.defaults.withCredentials = true
export default defineComponent({
  name: 'App'
})
</script>
