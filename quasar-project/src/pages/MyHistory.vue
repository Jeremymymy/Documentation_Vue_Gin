<template>
  <q-page>
    <div class="text-h3" align="center">
       <b class="title"><br /><q-icon name="update" /> History</b><br />
    </div>
    <q-card dark bordered class="bg-grey-9 my-card">
      <q-card-section>
        <div class="text-h4 q-mb-sm text-center">{{ documentTitle }}</div>
        <hr class="h1" />
        <div v-for="(item, index) in historyInfo" :key="index">
          <div class="q-ma-md show-row-equal-width">
            <div class="row">
              <div class="text-h5 col q-ma-md text-center">{{ item[0] }}</div>
              <div class="text-h5 col q-ma-md text-center">{{ item[1] }}</div>
              <div class="text-h5 col q-ma-md text-center">{{ item[2] }}</div>
              <div class="text-h5 col q-ma-md text-center">{{ item[3] }}</div>
              <div class="col q-ma-md text-center">
                <q-btn color="white" text-color="black" label="View" @click="fetchData" />
              </div>
            </div>
          </div>
          <hr class="h1" />
        </div>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script>
import axios from 'axios';
import { defineComponent, ref } from 'vue';
import { useRouter } from 'vue-router';

export default defineComponent({
  setup () {
    const router = useRouter();
    const fetchData = async () => {
      try {
        const response = await axios.get('');
        const data = response.data;
        router.push({ name: '/detail', params: { data } });
      } catch (error) {
        console.error(error);
      }
    };

    const historyInfo = [
      ['Version1', 'Author1', 'Date1'],
      ['Version2', 'Author2', 'Date2'],
      ['Version3', 'Author3', 'Date3']
    ];
    const documentTitle = ref('Title');

    return {
      fetchData,
      historyInfo,
      documentTitle
    };
  }
});
</script>

<style lang="sass">
.show-row-equal-width
  .row > div
    padding: 5px 10px
  .row + .row
    margin-top: 1rem

hr.h1
  border: 1px solid black
</style>
