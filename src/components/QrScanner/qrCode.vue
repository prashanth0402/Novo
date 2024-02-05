<template>
    <div>
        <div class="qrcode text-center" v-if="open && $vuetify.breakpoint.width >= 900">
            <div class="close-qr">
                <v-btn icon x-small class="error white--text" @click="closeQr"><v-icon>mdi-close</v-icon></v-btn>
            </div>
            <div class="qr-content">
                <h3>Download Novo Mobile App now</h3>
                <div class="image">
                    <picture>
                        <!-- <source srcset="../../assets/flattrade.png" type="image/webp"> -->
                        <source srcset="https://flattrade.s3.ap-south-1.amazonaws.com/promo/Ft_AllInOnelink.webp"
                            type="image/webp">
                        <img src="https://flattrade.s3.ap-south-1.amazonaws.com/promo/Ft_AllInOnelink.webp"
                            alt="upstox-qr-code" width="78" height="78">
                    </picture>

                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data: () => ({
        dialog: true,
    }),
    props: {
        open: Boolean
    },
    methods: {
        closeQr() {
            this.$emit('close');
        }
    }
}
</script>

<style  scoped>
.qrcode {
    width: 128px;
    background-color: #fff;
    padding: 12px 8px;
    border-radius: 8px;
    position: fixed;
    bottom: 30px;
    box-shadow: 0 4px 24px 0 rgba(0, 0, 0, .25);
    z-index: 1;
    right: 20px;
}

.close-qr {
    position: absolute;
    top: -12px;
    left: -6px;
}

.qr-content h3 {
    font-size: 12px;
    font-weight: 700;
    line-height: 16px;
    margin: 0 0 8px;
}

.qr-code .qr-content .image {
    width: 88px;
    padding: 4px;
    border: 1px solid #e5e5e5;
    border-radius: 4px;
}

.logo-img {
    padding-top: 10px;
}
</style>