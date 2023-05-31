<template>
    <q-page>
        <div class = 'text-h2 q-ma-lg'>文件版本快照</div>
        <div class = 'text-h2 q-mb-sm text-center'>{{articletitle}}</div>
        <div class = 'text-h4 q-lb-sm text-center'>{{articleauthor}}</div>
        <hr class = "h1">
        <div v-for = "(item,index) in historyinfo" :key ="index">
            <div class="q-ma-md show-row-equal-width">
                <div class = 'row'>
                    <div class = 'text-h5 col q-ma-md text-center'>{{item[0]}}</div>
                    <div class = 'text-h5 col q-ma-md text-center'>{{item[1]}}</div>
                    <div class = 'text-h5 col q-ma-md text-center'>{{item[2]}}</div>
                    <div class = 'text-h5 col q-ma-md text-center'>{{item[3]}}</div>
                    <div class = 'col q-ma-md text-center'>
                        <q-btn color="white" text-color="black" label="view" @click = 'fetchdata'></q-btn>
                    </div>
                </div>
            </div>
            <hr class = "h1">
        </div>
    </q-page>
</template>

<script>
import axios from 'axios';
import { defineComponent, ref } from 'vue';
import { useRouter } from 'vue-router';
export default defineComponent({
  setup () {
    const router = useRouter()
    const fetchdata = async () => {
      try {
        const response = await axios.get('')
        const data = response.data
        router.push({ name: '/detail', parem: { data } })
      } catch (error) {
        console.error(error)
      }
    }
    // const data = $route.params.data
    const historyinfo = [
      ['title1', 'author1', 'department1', 'date1'],
      ['title2', 'author2', 'department2', 'date2'],
      ['title3', 'author3', 'department3', 'date3']];
    const articletitle = ref('9');
    const articleauthor = ref('author');
    return {
      fetchdata,
      historyinfo,
      articletitle,
      articleauthor
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
