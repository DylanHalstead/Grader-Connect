<!-- Modified version of: https://github.com/serresebastien/Nuxt-Tailwind-Responsive-Navbar -->
<template>
  <header class="sticky top-0 bg-gray-950 shadow-lg">
    <nav class="w-full p-6 text-white">
      <div class="flex items-center justify-between">
        <!-- Header logo -->
        <div>
          <p class="text-xl font-semibold">Grader Connect</p>
        </div>

        <!-- Mobile toggle -->
        <div class="md:hidden">
          <button @click="toggleMenu">
            <svg
              class="h-8 w-8 fill-current"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path d="M4 6h16M4 12h16M4 18h16"></path>
            </svg>
          </button>
        </div>

        <!-- Navbar -->
        <div class="hidden md:block">
          <ul class="flex space-x-8 text-md font-sans">
            <li>
              <RouterLink
                :to="{ name: 'home' }"
                :active-class="'border-emerald-600'"
                class="border-b-2 border-transparent pb-1 hover:lg:border-emerald-600 transition-all"
              >
                Home
              </RouterLink>
            </li>
            <li>
              <RouterLink
                :to="{ name: 'assignments' }"
                :active-class="'border-emerald-600'"
                class="border-b-2 border-transparent pb-1 hover:lg:border-emerald-600 transition-all"
              >
                Assignments
              </RouterLink>
            </li>
            <li>
              <RouterLink
                :to="{ name: 'login' }"
                :active-class="'bg-emerald-500'"
                class="cta bg-emerald-600 hover:bg-emerald-500 px-3 py-2 rounded text-white font-semibold transition-all"
              >
                Login
              </RouterLink>
            </li>
          </ul>
        </div>

        <!-- Dark Background Transition -->
        <transition
          enter-class="opacity-0"
          enter-active-class="ease-out transition-medium"
          enter-to-class="opacity-100"
          leave-class="opacity-100"
          leave-active-class="ease-out transition-medium"
          leave-to-class="opacity-0"
        >
          <div
            @keydown.esc="isOpen = false"
            v-show="isOpen"
            class="z-10 fixed inset-0 transition-opacity"
          >
            <div
              @click="isOpen = false"
              class="absolute inset-0 bg-black opacity-50"
              tabindex="0"
            ></div>
          </div>
        </transition>

        <!-- Drawer Menu -->
        <aside
          class="p-5 transform top-0 left-0 w-64 bg-gray-950 fixed h-full overflow-auto ease-in-out transition-all duration-300 z-30"
          :class="{ 'translate-x-0': isOpen, '-translate-x-full': !isOpen }"
        >
          <div class="close">
            <button class="absolute top-0 right-0 mt-4 mr-4" @click="isOpen = false">
              <svg
                class="w-6 h-6"
                fill="none"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>

          <span @click="isOpen = false" class="flex w-full items-center py-4 border-b">
            <p class="text-xl font-semibold">Grader Connect</p>
          </span>

          <ul class="divide-y font-sans">
            <li>
              <RouterLink
                :to="{ name: 'home' }"
                :active-class="'text-emerald-400'"
                @click="isOpen = false"
                class="my-4 inline-block hover:text-emerald-400 transition-all"
              >
                Home
              </RouterLink>
            </li>
            <li>
              <RouterLink
                :to="{ name: 'assignments' }"
                :active-class="'text-emerald-400'"
                @click="isOpen = false"
                class="my-4 inline-block hover:text-emerald-400 transition-all"
              >
                Assignments
              </RouterLink>
            </li>
            <li>
              <RouterLink
                :to="{ name: 'login' }"
                :active-class="'bg-emerald-500'"
                @click="isOpen = false"
                class="my-4 w-full text-center font-semibold cta inline-block bg-emerald-600 hover:bg-emerald-500 px-3 py-2 rounded text-white transition-all"
              >
                Login
              </RouterLink>
            </li>
          </ul>
        </aside>
      </div>
    </nav>
  </header>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'

const isOpen = ref(false)

function toggleMenu() {
  isOpen.value = !isOpen.value
}

watch(
  isOpen,
  (isOpen) => {
    if (isOpen) document.body.style.setProperty('overflow', 'hidden')
    else document.body.style.removeProperty('overflow')
  },
  { immediate: true }
)

onMounted(() => {
  document.addEventListener('keydown', (e) => {
    if (e.key === 'Escape' && isOpen.value) isOpen.value = false
  })
})
</script>
