<template>
  <div class="min-h-screen">
    <SiteNavigation />
    <div class="md:mx-16 my-10">
      <RouterView />
    </div>
  </div>
</template>

<script setup>
import { RouterView } from 'vue-router'
import SiteNavigation from './components/SiteNavigation.vue'
</script>
