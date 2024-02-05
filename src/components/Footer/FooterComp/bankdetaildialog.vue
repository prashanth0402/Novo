<template>
    <div>
        <v-dialog v-model="dialog" max-width="900" transition="dialog-top-transition"
            :class="$vuetify.breakpoint.width > 900 ? 'card-position' : undefined">
            <v-card>
                <v-card-title>
                    <v-layout class="d-flex">
                        <v-flex v-if="indicator != 'bank'">
                            <h3 class="font-weight-medium black--text">Attention Investors!</h3>

                        </v-flex>
                        <v-flex class="d-flex justify-end">
                            <v-icon @click="hideBankDetails">mdi-close</v-icon>
                        </v-flex>
                    </v-layout>
                </v-card-title>
                <v-divider class="my-2" v-if="indicator != 'bank'"></v-divider>
                <v-card-text>
                    <v-simple-table v-if="indicator == 'bank'">
                        <template v-slot:default>
                            <thead>
                                <tr>
                                    <th class="text-left">
                                        Bank Account No.
                                    </th>
                                    <th class="text-left">
                                        IFSC
                                    </th>
                                    <th class="text-left">
                                        Bank Account Type
                                    </th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="item, idx in desserts" :key="idx">
                                    <td>{{ item.acNo }}</td>
                                    <td class="text-left">{{ item.ifsc }}</td>
                                    <td class="text-left">{{ item.accountType }}</td>

                                </tr>
                            </tbody>
                        </template>
                    </v-simple-table>
                    <div v-else>
                        <ul style="font-size:12px;">
                            <li>Beware of fixed/guaranteed/regular returns/ capital protection schemes. Brokers or their
                                authorized persons or any of their associates are not authorized to offer
                                fixed/guaranteed/regular returns/ capital protection on your investment or authorized to
                                enter
                                into any loan agreement with you to pay interest on the funds offered by you. <br>Please
                                note
                                that in case of default of a member claim for funds or securities given to the broker under
                                any
                                arrangement/ agreement of indicative return will not be accepted by the relevant Committee
                                of
                                the Exchange as per the approved norms.</li>
                            <li>Do not keep funds idle with the Stock Broker. Please note that your stock broker has to
                                return
                                the credit balance lying with them, within three working days in case you have not done any
                                transaction within last 30 calendar days. Please note that in case of default of a Member,
                                claim
                                for funds and securities, without any transaction on the exchange will not be accepted by
                                the
                                relevant Committee of the Exchange as per the approved norms.</li>
                            <li>Check the frequency of accounts settlement opted for. If you have opted for running account,
                                please ensure that your broker settles your account and, in any case, not later than once in
                                90
                                days (or 30 days if you have opted for 30 days settlement). In case of declaration of
                                trading
                                member as defaulter, the claims of clients against such defaulter member would be subject to
                                norms for eligibility of claims for compensation from IPF to the clients of the defaulter
                                member. These norms are available on Exchange website at following link: <a
                                    href="https://www.nseindia.com/invest/about-defaulter-section" target="_blank">NSC</a>,
                                <a href="https://www.mcxindia.com/docs/default-source/investor-grievance-arbitration/defaulters/norms-on-rejections-of-claims/sop-process-faqs-for-handling-of-claims-of-investors-of-defaulter-member.pdf?sfvrsn=b5737291_6"
                                    target="_blank">MCX</a>
                            </li>
                            <li>Brokers are not permitted to accept transfer of securities as margin. Securities offered as
                                margin/ collateral MUST remain in the account of the client and can be pledged to the broker
                                only by way of ‘margin pledge’, created in the Depository system. Clients are not permitted
                                to
                                place any securities with the broker or associate of the broker or authorized person of the
                                broker for any reason. Broker can take securities belonging to clients only for settlement
                                of
                                securities sold by the client.</li>
                            <li>Always keep your contact details viz. Mobile number/Email ID updated with the stock broker.
                                Email and mobile number is mandatory and you must provide the same to your broker for
                                updation
                                in Exchange records. You must immediately take up the matter with Stock Broker/Exchange if
                                you
                                are not receiving the messages from Exchange/Depositories regularly.</li>
                            <li>Don't ignore any emails/SMSs received from the Exchange for trades done by you. Verify the
                                same
                                with the Contract notes/Statement of accounts received from your broker and report
                                discrepancy,
                                if any, to your broker in writing immediately and if the Stock Broker does not respond,
                                please
                                take this up with the Exchange/Depositories forthwith.</li>
                            <li>Check messages sent by Exchanges on a weekly basis regarding funds and securities balances
                                reported by the trading member, compare it with the weekly statement of account sent by
                                broker
                                and immediately raise a concern to the exchange if you notice a discrepancy.</li>
                            <li>Please do not transfer funds, for the purposes of trading to anyone, including an authorized
                                person or an associate of the broker, other than a SEBI registered Stock broker.</li>
                            <li>Do not deal with unregistered intermediaries (who are not registered with SEBI/Exchanges).
                            </li>
                        </ul>
                    </div>
                </v-card-text>
            </v-card>
        </v-dialog>
    </div>
</template>

<script>
export default {
    data() {
        return {
            desserts: [
                {
                    acNo: '"059485700000129"',
                    ifsc: 'YESB0000594',
                    accountType: 'Up Streaming Client Nodal Bank Account'
                },
                {
                    acNo: '"01660340001487"',
                    ifsc: 'HDFC0000166',
                    accountType: 'Up Streaming Client Nodal Bank Account'
                },
                {
                    acNo: '"039505008960"',
                    ifsc: 'ICIC0000395',
                    accountType: 'Up Streaming Client Nodal Bank Account'
                },
                {
                    acNo: '"01660340001512"',
                    ifsc: 'HDFC0000166',
                    accountType: 'Up Streaming Client Nodal Bank Account'
                },
                {
                    acNo: '"01660340001470"',
                    ifsc: 'HDFC0000166',
                    accountType: 'Up Streaming Client Nodal Bank Account'
                },
                {
                    acNo: '"2221262044881111"',
                    ifsc: 'AUBL0002620',
                    accountType: 'Up Streaming Client Nodal Bank Account'
                },
                {
                    acNo: '"01660340001460"',
                    ifsc: 'HDFC0000166',
                    accountType: 'Up Streaming Client Nodal Bank Account'
                },

            ],
        }
    },
    props: {
        dialog: Boolean,
        indicator: String
    },
    methods: {
        hideBankDetails() {
            this.$emit('hideBankDetails')
        }
    }
}
</script>

<style scoped>
.v-dialog__content {
    top: -60px !important;
}
</style>