<!-- Modified version of: https://github.com/serresebastien/Nuxt-Tailwind-Responsive-Navbar -->
<template>
  <header class="sticky top-0 bg-gray-950 shadow-lg">
    <nav class="w-full p-6 text-white">
      <div class="flex items-center justify-between">
        <!-- Header logo -->
        <div>
          <p class="text-xl font-semibold">Grader Connect</p>
        </div>

        <!-- Navbar -->
        <div
          class="transform md:block top-0 left-0 max-md:p-5 max-md:w-64 max-md:bg-gray-950 max-md:fixed max-md:h-full max-md:overflow-auto max-md:z-30"
          :class="
            (isOpen ? ' translate-x-0 ease-in-out duration-300' : '') +
            (isScreenSmall ? ' -translate-x-full' : '') +
            (shouldApplyTransition ? ' transition-transform' : ' transition-none')
          "
        >
          <ul
            :class="{
              'flex space-x-8 text-md font-sans': !isScreenSmall,
              'divide-y': isScreenSmall
            }"
          >
            <li>
              <RouterLink
                :to="{ name: 'home' }"
                :active-class="!isScreenSmall ? 'border-b-emerald-600' : 'text-emerald-400'"
                :class="{
                  'border-y-2 border-y-transparent pb-1 hover:border-b-emerald-600 hover:transition-colors':
                    !isScreenSmall,
                  'my-4 inline-block hover:text-emerald-400 transition-colors': isScreenSmall
                }"
                @click="isOpen = false"
              >
                Home
              </RouterLink>
            </li>
            <li>
              <RouterLink
                :to="{ name: 'assignments' }"
                :active-class="!isScreenSmall ? 'border-b-emerald-600' : 'text-emerald-400'"
                :class="{
                  'border-y-2 border-y-transparent pb-1 hover:border-b-emerald-600 hover:transition-colors':
                    !isScreenSmall,
                  'my-4 inline-block hover:text-emerald-400 transition': isScreenSmall
                }"
                @click="isOpen = false"
              >
                Assignments
              </RouterLink>
            </li>
            <li>
              <RouterLink
                :to="{ name: 'login' }"
                :active-class="'bg-emerald-500'"
                @click="isOpen = false"
                :class="{
                  'bg-emerald-600 hover:bg-emerald-500 px-3 py-2 rounded text-white font-semibold transition-colors':
                    !isScreenSmall,
                  'my-4 w-full text-center font-semibold cta inline-block bg-emerald-600 hover:bg-emerald-500 px-3 py-2 rounded text-white transition-colors':
                    isScreenSmall
                }"
              >
                Login
              </RouterLink>
            </li>
          </ul>
        </div>

        <!-- Mobile toggle -->
        <div class="md:hidden inline-flex">
          <button @click="isOpen = true">
            <svg
              class="h-6 w-6 fill-current"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              viewBox="2 2 20 20"
              stroke="currentColor"
            >
              <path d="M4 6h16M4 12h16M4 18h16"></path>
            </svg>
          </button>
        </div>

        <!-- Dark Background Transition -->
        <Transition
          enter-class="opacity-0"
          enter-active-class="ease-out transition-medium"
          enter-to-class="opacity-100"
          leave-class="opacity-100"
          leave-active-class="ease-out transition-medium"
          leave-to-class="opacity-0"
        >
          <div v-show="isOpen && isScreenSmall" class="z-10 fixed inset-0 transition-opacity">
            <div
              @click="isOpen = false"
              class="absolute inset-0 bg-black opacity-50"
              tabindex="0"
            ></div>
          </div>
        </Transition>
      </div>
    </nav>
  </header>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

const isOpen = ref(false)
const isScreenSmall = ref(window.innerWidth <= 768)
const shouldApplyTransition = ref(true)
const resizeTimeout = ref(null)

const handleResize = () => {
  isScreenSmall.value = window.innerWidth <= 768
  if (!isScreenSmall.value) isOpen.value = false
  // Prevent slider on resize
  shouldApplyTransition.value = false
  clearTimeout(resizeTimeout)
  resizeTimeout.value = setTimeout(() => {
    shouldApplyTransition.value = true
  }, 300)
}

onMounted(() => {
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
})
</script>
