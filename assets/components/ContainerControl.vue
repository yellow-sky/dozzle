<template>
  <div class="is-size-9 is-uppercase columns is-marginless is-mobile">
    <span @click.stop.prevent="pullContainerImage()" title="Pull">
      <mdi-light-download />
    </span>
    <span @click.stop.prevent="restartContainer()" title="Restart">
      <mdi-light-refresh />
    </span>
  </div>
</template>

<script lang="ts" setup>
import config from "@/stores/config";
import { useProgrammatic } from "@oruga-ui/oruga-next";
const { oruga } = useProgrammatic();

const props = defineProps({
  id: {
    type: String,
    required: true,
  },
});

const restartContainer = async function () {
  oruga.notification.open({
    message: "Try to restart container...",
    position: "top",
    variant: "info",
    duration: 5000,
  });
  const resp = await fetch(`${config.base}/api/container/restart?id=${props.id}`);
  let respText = await resp.text();
  if (respText.length == 0)
    if (resp.status < 400) {
      respText = "sucsessfull";
    } else {
      respText = "failed";
    }
  const msgText = "Restart container: " + respText;
  oruga.notification.open({
    message: msgText,
    position: "top",
    variant: resp.status < 400 ? "success" : "danger",
    duration: 5000,
  });
};

const pullContainerImage = async function () {
  oruga.notification.open({
    message: "Try to pull container image...",
    position: "top",
    variant: "info",
    duration: 5000,
  });
  const resp = await fetch(`${config.base}/api/container/image/pull?id=${props.id}`);
  let respText = await resp.text();
  if (respText.length == 0)
    if (resp.status < 400) {
      respText = "sucsessfull";
    } else {
      respText = "failed";
    }
  const msgText = "Image pull: " + respText;
  oruga.notification.open({
    message: msgText,
    position: "top",
    variant: resp.status < 400 ? "success" : "danger",
    duration: 5000,
  });
};
</script>

<style></style>
