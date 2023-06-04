<template>
  <Teleport to="body">
    <div class="flex justify-center items-center">
      <Transition
        enter-from-class="opacity-0 scale-75"
        leave-to-class="opacity-0 scale-75"
        enter-active-class="transition-all duration-200 ease-in-out"
        leave-active-class="transition-all duration-200 ease-in-out"
        enter-to-class="opacity-100 scale-100"
        leave-from-class="opacity-100 scale-100"
      >
        <div
          v-if="modalActive"
          class="fixed top-0 z-20 m-8 mt-28 p-8 pt-4 bg-slate-950 rounded-lg shadow-[0_0_12px_rgba(31,41,55,.3)] box-border"
        >
          <button
            class="w-full flex justify-end text-white hover:text-emerald-600 transition-colors ease-in-out"
            @click="$emit('close-modal')"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="2"
              stroke="currentColor"
              class="w-6 h-6"
            >
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
          <div class="mb-6 text-white">
            <div class="mb-8">
              <h1 class="mb-1 text-5xl font-semibold">Title Goes Here</h1>
              <h2 class="text-xl italic">Due: 12/12/2023</h2>
            </div>
            <div class="mb-4">
              <div class="flex items-center">
                <span
                  class="w-1.5 h-1.5 ml-[2px] mr-2 rounded-full bg-emerald-600 shadow-[0_0_1px_2px_rgba(22,163,74,0.3)]"
                ></span>
                <p>Grader's Name</p>
              </div>
              <hr class="h-px my-2 bg-gray-700 bg-opacity-50 border-0" />
              <div class="flex items-center">
                <span
                  class="w-1.5 h-1.5 mr-2 ml-[2px] rounded-full bg-gray-500 shadow-[0_0_1px_2px_rgba(107,114,128,0.3)]"
                ></span>
                <p>Grader's Name</p>
              </div>
              <hr class="h-px my-2 bg-gray-700 bg-opacity-50 border-0" />
              <div class="flex items-center">
                <span
                  class="w-1.5 h-1.5 mr-2 ml-[2px] rounded-full bg-gray-500 shadow-[0_0_1px_2px_rgba(107,114,128,0.3)]"
                ></span>
                <p>Grader's Name</p>
              </div>
            </div>
          </div>
          <button
            class="w-full py-1 text-white font-semibold bg-emerald-600 rounded-md cursor-pointer hover:bg-emerald-500 hover:shadow-[0_0_8px_1px_rgba(22,163,74,0.5)] transition-all duration-300 ease-in-out"
          >
            Submit
          </button>
        </div>
      </Transition>
    </div>
    <!-- Dark Background Transition -->
    <Transition
      enter-class="opacity-0"
      enter-active-class="ease-out transition-all"
      enter-to-class="opacity-100"
      leave-class="opacity-100"
      leave-active-class="ease-out transition-all"
      leave-to-class="opacity-0"
    >
      <div v-show="modalActive" class="fixed z-10 inset-0 transition-opacity">
        <div
          @click="$emit('close-modal')"
          class="absolute inset-0 bg-black opacity-50"
          tabindex="0"
        ></div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
defineEmits(['close-modal'])
defineProps({
  modalActive: {
    type: Boolean,
    default: false
  }
})
</script>
