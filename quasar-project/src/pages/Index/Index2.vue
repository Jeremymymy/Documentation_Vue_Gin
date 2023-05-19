<template>
  <div  class="text-h3 q-pa-sm " align="center"><b class="title"><br />TSMC最新公告</b></div>
  <div class="q-pa-md">
    <q-carousel
      v-model="slide"
      :autoplay="autoplay"
      infinite
      swipeable
      animated
      padding
      arrows
      navigation
      navigation-icon="radio_button_unchecked"
      height="500px"
      class="bg-dark glossy text-white rounded-borders"
    >
      <q-carousel-slide name="style" class="text-center">
        <q-scroll-area class="fit">
          <q-icon name="style" size="56px" />
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
        </q-scroll-area>
      </q-carousel-slide>

      <q-carousel-slide name="tv" class="text-center">
        <q-scroll-area class="fit">
          <q-icon name="live_tv" size="56px" />
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
        </q-scroll-area>
      </q-carousel-slide>

      <q-carousel-slide name="layers" class="text-center">
        <q-scroll-area class="fit">
          <q-icon name="layers" size="56px" />
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
        </q-scroll-area>
      </q-carousel-slide>

      <q-carousel-slide name="map" class="text-center">
        <q-scroll-area class="fit">
          <q-icon name="terrain" size="56px" />
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
        </q-scroll-area>
      </q-carousel-slide>
    </q-carousel>
  </div>
  <div  class="text-h3 q-pa-sm " align="center"><b class="title">公開文件</b></div>
  <div class=" q-pa-md row section-card items-start q-gutter-md"  align="center">
    <q-card class="my-card" flat bordered>
      <q-img
        src="https://cdn.quasar.dev/img/parallax2.jpg"
      />
      <q-card-section>
        <div class="text-overline text-orange-9">Overline</div>
        <div class="text-h5 q-mt-sm q-mb-xs">Title</div>
        <div class="text-caption text-grey">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
        </div>
      </q-card-section>
      <q-card-actions align="right">
        <q-btn flat round color="red" icon="favorite" />
        <q-btn flat round color="teal" icon="bookmark" />
        <q-btn flat round color="primary" icon="share" />
      </q-card-actions>
    </q-card>
     <q-card class="my-card" flat bordered>
      <q-img
        src="https://cdn.quasar.dev/img/parallax2.jpg"
      />

      <q-card-section>
        <div class="text-overline text-orange-9">Overline</div>
        <div class="text-h5 q-mt-sm q-mb-xs">Title</div>
        <div class="text-caption text-grey">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
        </div>
      </q-card-section>

      <q-card-actions align="right">
        <q-btn flat round color="red" icon="favorite" />
        <q-btn flat round color="teal" icon="bookmark" />
        <q-btn flat round color="primary" icon="share" />
      </q-card-actions>
    </q-card>
     <q-card class="my-card" flat bordered>
      <q-img
        src="https://cdn.quasar.dev/img/parallax2.jpg"
      />

      <q-card-section>
        <div class="text-overline text-orange-9">Overline</div>
        <div class="text-h5 q-mt-sm q-mb-xs">Title</div>
        <div class="text-caption text-grey">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
        </div>
      </q-card-section>

      <q-card-actions align="right">
        <q-btn flat round color="red" icon="favorite" />
        <q-btn flat round color="teal" icon="bookmark" />
        <q-btn flat round color="primary" icon="share" />
      </q-card-actions>
    </q-card>
     <q-card class="my-card" flat bordered>
      <q-img
        src="https://cdn.quasar.dev/img/parallax2.jpg"
      />

      <q-card-section>
        <div class="text-overline text-orange-9">Overline</div>
        <div class="text-h5 q-mt-sm q-mb-xs">Title</div>
        <div class="text-caption text-grey">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
        </div>
      </q-card-section>

      <q-card-actions align="right">
        <q-btn flat round color="red" icon="favorite" />
        <q-btn flat round color="teal" icon="bookmark" />
        <q-btn flat round color="primary" icon="share" />
      </q-card-actions>
    </q-card>
  </div>
</template>

<script>
import { ref } from 'vue'
export default {
  name: 'Index2',
  props: ['data'],
  setup () {
    return {
      expanded: ref(false),
      slide: ref('style'),
      autoplay: ref(true),
      lorem: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.'
    }
  }
}
</script>

<style lang="sass" scoped>

.section-card
  margin-right: 10%
  margin-left: 10%

.my-card
  width: 100%
  max-width: 350px
</style>
