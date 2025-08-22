<template>
  <div>
    <div class="flex flex-col w-full border-opacity-50">
      <div class="grid card bg-base-300 rounded-box place-items-center">
        <WelcomeIcon />Hello daisyUI
      </div>
    </div>
    <div class="overflow-x-auto">
      <table class="table">
        <!-- head -->
        <thead>
          <tr>
            <th></th>
            <th>Name</th>
            <th>Job</th>
            <th>Favorite Color</th>
            <th>Create At</th>
          </tr>
        </thead>
        <tbody>
          <!-- row 1 -->
          <tr class="hover" v-for="(item, index) in list" :key="item.id">
            <td>{{ index + 1 }}</td>
            <td>{{ item.name }}</td>
            <td>{{ item.job }}</td>
            <td>{{ item.favorite_color }}</td>
            <td>{{ formatTime(item.create_at) }}</td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="flex flex-col w-full border-opacity-50">
      <div class="grid card bg-base-300 rounded-box place-items-center">
        <IconAirplaneLanding /><IconAppFilled
          style="font-size: 2em; color: red"
        />
      </div>
    </div>
  </div>
</template>
  


<script>
import dayjs from "dayjs";
import api from "@/api/api";
import IconAirplaneLanding from "~icons/fluent/airplane-landing-24-regular";
import IconAppFilled from "~icons/fluent/apps-24-filled";

export default {
  name: "Home",
  data() {
    return {
      list: [],
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      this.list = [
        {
          id: 1,
          name: "Cy Ganderton",
          job: "Quality Control Specialist",
          favorite_color: "Blue",
          create_at: 1687150538,
        },
        {
          id: 2,
          name: "Hart Hagerty",
          job: "Desktop Support Technician",
          favorite_color: "Purple",
          create_at: 1687150538,
        },
        {
          id: 3,
          name: "Brice Swyre",
          job: "Tax Accountant",
          favorite_color: "Red",
          create_at: 1687150538,
        },
      ];
    },
    formatTime(time) {
      return dayjs.unix(time).format("YYYY-MM-DD");
    },
  },
};
</script>

<script setup>
import WelcomeIcon from "@/assets/icons/welcome.svg";
</script>