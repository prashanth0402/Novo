<template>
    <div>
        <div>
            <DIpoWindow @showalert="showCard2" @Recall="recall" />
            <DSuccessCard :Alert="isCardVisible2" @timeout2="hideCard2" />
        </div>
    </div>
</template>

<script>
// Desktop view
import DIpoWindow from "../components/Ipo/IpoWindow/IpoWindow.vue"
import DSuccessCard from "./Ipo/SuccessCard/SuccessCard.vue"
export default {
    components: {
        DSuccessCard,
        DIpoWindow,
    },
    data() {
        return {
            timeout: 3000,
            isCardVisible1: false,
            isCardVisible2: false,
            screenWidth: window.innerWidth,
            callHistory: 0,
        }
    },
    methods: {
        showCard1() {
            // console.log("showCard1")
            this.isCardVisible1 = true;
        },
        showCard2() {
            // console.log("showCard2")
            this.isCardVisible2 = true;
        },
        hideCard1() {
            // console.log("hideCard1")
            this.isCardVisible1 = false;
        },
        hideCard2() {
            // console.log("hideCard2")
            this.isCardVisible2 = false;
        },
        // to resize the screen
        handleResize() {
            this.screenWidth = window.innerWidth;
        },
        recall() {
            this.callHistory++
        },
    },
    computed: {
        isSmallScreen() {
            return this.screenWidth <= 600;
        },
    },
    mounted() {
        window.addEventListener('resize', this.handleResize);
    },
    beforeDestroy() {
        window.removeEventListener('resize', this.handleResize);
    },
};


</script>
<style>
.top-right-dialog {
    position: fixed;
    top: 20px;
    right: 20px;
    z-index: 9999;
    /* Adjust the z-index as needed */
}
</style>
