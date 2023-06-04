<!-- Modified version of: https://github.com/serresebastien/Nuxt-Tailwind-Responsive-Navbar -->
<template>
  <header class="w-full relative top-0 bg-gray-950 shadow-lg">
    <nav class="w-full p-6 text-white">
      <div class="flex items-center justify-between">
        <div>
          <p class="text-xl font-semibold">Grader Connect</p>
        </div>

        <div
          class="transform top-0 left-0 md:p-0 p-5 md:w-auto w-64 md:bg-transparent bg-gray-950 md:static fixed md:h-auto h-full md:overflow-visible overflow-auto z-30"
          :class="
            (isNavOpen ? 'translate-x-0 ease-in-out duration-300' : '') +
            (isScreenSmall ? ' -translate-x-full' : '') +
            (shouldApplySwipe ? ' transition-transform' : ' transition-none')
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
                class="transition-colors"
                :active-class="isScreenSmall ? 'text-emerald-400' : 'border-b-emerald-600'"
                :class="
                  isScreenSmall
                    ? 'my-4 inline-block hover:text-emerald-400'
                    : 'pb-1 border-y-2 border-y-transparent hover:border-b-emerald-600'
                "
                @click="isNavOpen = false"
              >
                Home
              </RouterLink>
            </li>
            <li>
              <RouterLink
                :to="{ name: 'assignments' }"
                class="transition-colors"
                :active-class="isScreenSmall ? 'text-emerald-400' : 'border-b-emerald-600'"
                :class="
                  isScreenSmall
                    ? 'my-4 inline-block hover:text-emerald-400'
                    : 'pb-1 border-y-2 border-y-transparent hover:border-b-emerald-600'
                "
                @click="isNavOpen = false"
              >
                Assignments
              </RouterLink>
            </li>
            <li>
              <RouterLink
                :to="{ name: 'login' }"
                class="px-3 py-2 font-semibold bg-emerald-600 rounded hover:bg-emerald-500 hover:shadow-[0_0_12px_0_rgba(22,163,74,0.5)] duration-300 transition-all ease-in-out"
                :active-class="'bg-emerald-500 shadow-[0_2px_12px_0_rgba(22,163,74,0.5)]'"
                :class="isScreenSmall ? 'w-full my-4 inline-block text-center' : ''"
                @click="isNavOpen = false"
              >
                Login
              </RouterLink>
            </li>
          </ul>
        </div>

        <!-- Mobile toggle -->
        <div class="md:hidden inline-flex">
          <button @click="isNavOpen = true">
            <svg
              class="h-6 w-6"
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
          <div v-show="isNavOpen && isScreenSmall" class="fixed z-10 inset-0 transition-opacity">
            <div
              @click="isNavOpen = false"
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

const isNavOpen = ref(false)
const isScreenSmall = ref(window.innerWidth <= 768)
const shouldApplySwipe = ref(true)
const resizeTimeout = ref(null)

const handleResize = () => {
  isScreenSmall.value = window.innerWidth <= 768
  if (!isScreenSmall.value) isNavOpen.value = false
  // Prevent slider effect on resize
  shouldApplySwipe.value = false
  clearTimeout(resizeTimeout)
  resizeTimeout.value = setTimeout(() => {
    shouldApplySwipe.value = true
  }, 300)
}

onMounted(() => {
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
})
</script>
