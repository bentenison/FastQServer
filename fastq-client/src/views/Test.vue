<!-- App.vue
<template>
    <div id="app">
        <v-btn @click="printToUSBPrinter">Print to USB Printer</v-btn>
        <div class="row d-flex align-items-center justify-content-center">
            <div class="col-md-3" id="ticket">
                <div class="heading text-center d-flex align-items-center justify-content-between">
                    <v-img max-height="50" max-width="50" src="../assets/sec-icon.svg"></v-img>
                    <h3>Welcome to FASTQ</h3>
                </div>
                <div class="number">
                    <p class="text-body font-bold">Your number is :</p>
                    <h1>12201</h1>
                </div>
            </div>
        </div>
    </div>
</template>
   
<script>
import html2canvas from 'html2canvas';
export default {
    data() {
        return {
            capturedImageUrl: null,
            imageUint8Array: null,
        }
    },
    methods: {
        async printToUSBPrinter() {
            try {
                const device = await navigator.usb.requestDevice({ filters: [] });

                await device.open();
                const interfaceNumber = device.configuration.interfaces[0].interfaceNumber;

                await device.selectConfiguration(1);
                await device.claimInterface(interfaceNumber)
                console.log("selected device", device);
                const endpoint = device.configuration.interfaces[0].alternate.endpoints[1];
                console.log("object,", endpoint);
                this.convertBase64(this.capturedImageUrl)
                console.log("binary",this.imageUint8Array);
                const hexCode = this.convertImageToHex(this.imageUint8Array)
                console.log("object hex",hexCode);
                // const outEndpoint = endpoint.direction === "out" ? endpoint : null;
                // this.captureContent().then(async res => {
                //     console.log("image captured", this.capturedImageUrl);
                //     const imageData = await this.capturedImageUrl.arrayBuffer();

                //     // Combine image data into a single Uint8Array
                //     const combinedData = new Uint8Array(imageData);
                //     // await device.transferOut(endpoint.endpointNumber, new TextEncoder().encode(this.capturedImageUrl));
                //     // await device.transferOut(endpoint.endpointNumber, combinedData);
                //     await this.sendImageToPrinter(combinedData, device, endpoint)

                //     await device.close();
                // })
                // Replace this with your actual content to be printed
                const contentToPrint = "Hello, USB Printer!\n";
                // await device.transferOut(endpoint.endpointNumber, new TextEncoder().encode(contentToPrint));
                // await device.transferOut(endpoint.endpointNumber,this.imageUint8Array);
                await device.close();
                // Write to the USB printer
            } catch (error) {
                // resolve(true)
                console.error("Error printing to USB printer:", error);
            }
        },
        async sendImageToPrinter(imageData, device, endpoint) {
            const CHUNK_SIZE = 64; // Adjust this based on your printer's specifications

            // Split the image data into chunks
            for (let offset = 0; offset < imageData.byteLength; offset += CHUNK_SIZE) {
                const chunk = imageData.slice(offset, offset + CHUNK_SIZE);

                // Send each chunk to the printer
                await device.transferOut(endpoint.endpointNumber, chunk);
            }
        },

        captureContent() {
            return new Promise((resolve, reject) => {
                try {
                    const element = document.getElementById('ticket'); // Replace with the ID of the HTML element you want to capture
                    html2canvas(element).then(async canvas => {
                        this.showCapturedImage(canvas.toDataURL())
                        resolve(true)
                    });


                    // console.log("element",element);
                    // console.log("object",canvas);
                    // Do something with the canvas, such as saving it as an image or displaying it
                } catch (error) {
                    console.error('Error capturing content:', error);
                    reject(error)
                }

            })
        },
        showCapturedImage(dataUrl) {
            // Do something with the captured image, such as displaying it in an <img> element
            this.capturedImageUrl = dataUrl;
            // console.log("image",dataUrl);
        },
        convertBase64(imagebase64) {
            // Convert base64 image data to Uint8Array
            this.imageUint8Array = Uint8Array.from(atob(imagebase64), (c) => c.charCodeAt(0));
        },
        // Create a function to convert a binary image to hex codes
        convertImageToHex(binaryData) {
            let hexCodes = [];
            for (let i = 0; i < binaryData.length; i++) {
                hexCodes.push(binaryData[i].toString(16).padStart(2, "0"));
            }
            return hexCodes.join("");
        }
    },
    mounted() {
        if ('usb' in navigator) {
            // WebUSB is supported
            window.alert("usb is supported on the device")
        } else {
            // WebUSB is not supported
            window.alert("usb is not supported on the device")
        }
    }
};
</script>
  
<style>
#app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    text-align: center;
    color: #2c3e50;
    margin-top: 35px;
}

button {
    font-size: 16px;
    padding: 10px 20px;
    cursor: pointer;
}
</style>
   -->

<template>
    <div class="flex flex-center">
        <div style="width:50vw" class="d-flex align-items-center justify-content-between">
            <div class="text-center q-mt-md">
                <v-btn @click="printwithConnected">Print to USB Printer</v-btn>
            </div>
            <div class="text-center q-mt-md">
                <v-btn @click="connectBluetoothPrinter">connect printer</v-btn>
            </div>
        </div>
    </div>
</template>
  
  
<script>
import EscPosEncoder from '../../node_modules/esc-pos-encoder/dist/esc-pos-encoder.esm.js';
import axios from "axios"
export default {
    name: 'PageIndex',
    data() {
        return {
            msg: '',
            printCharacteristic: null,
            device: null,
            // isMobile: this.$q.platform.is.mobile
        }
    },
    methods: {
        async listBluetoothDevices() {
            try {
                // Check if Web Bluetooth API is available
                if (!navigator.bluetooth) {
                    console.error('Web Bluetooth API is not supported in this browser.');
                    return;
                }

                // Request permission to use Web Bluetooth API
                await navigator.bluetooth.requestDevice({
                    acceptAllDevices: true,
                    optionalServices: ['battery_service'], // Specify services if needed
                });

                // Search for Bluetooth devices
                const devices = await navigator.bluetooth.getDevices();

                // Handle the list of devices
                console.log('List of Bluetooth devices:', devices);
            } catch (error) {
                console.error('Error while listing Bluetooth devices:', error);
            }
        },
        printwithConnected() {
            const self = this
            this.device.gatt
                .connect()
                .then(server =>
                    server.getPrimaryService('000018f0-0000-1000-8000-00805f9b34fb')
                )
                .then(service =>
                    service.getCharacteristic('00002af1-0000-1000-8000-00805f9b34fb')
                )
                .then(characteristic => {
                    console.log('characteristic', characteristic)
                    self.printCharacteristic = characteristic
                    let encoder = new EscPosEncoder({
                        imageMode: 'raster'
                    });

                    // let result = encoder
                    //     .initialize()
                    //     .text('The quick brown fox jumps over the lazy dog')
                    //     .newline()
                    //     .qrcode('https://nielsleenheer.com')
                    //     .encode();

                    console.log("router", window.location);
                    const url = new URL(window.location.origin);

                    // Get the IP address from the URL
                    const ipAddress = url.hostname;
                    let img = new Image();
                    // if (ipAddress === "localhost") {
                    //     return
                    // }
                    // img.src = `https://${ipAddress}:443/uploaded/logo.png`;
                    img.src = `http://localhost:8080/server/uploaded/logo.png`;

                    img.onload = async function () {
                        let result = encoder.align('center')
                            .image(img, 160, 160, 'atkinson').newline()
                            .encode()
                        console.log("hello", result);
                        const dataSize = result.length;
                        let startIndex = 0;
                        const maxChunkSize = characteristic.properties.writeWithoutResponse ? 512 : 20;
                        // Send chunks of data until the entire image is sent

                        // this.loop(0, result, device)
                        // characteristic.writeValue(result)
                        // 
                        self.printlogo(startIndex, result, maxChunkSize).then(async () => {

                            let data = {
                                service: 'Parts Order',
                                ticket_status: 'CREATED',
                                ticket_number: '15',
                                started_serving_at: '2017-01-01 10:31:00',
                                end_serving_at: '2017-01-01 10:31:00',
                                ticket_name: 'A-15',
                                company_code: 'nDmmyD0',
                                company_name: 'fastq solutions ltd',
                                updated_by: '73ac3366-c04d-11ee-a5d5-ce47405e851e',
                                timeDate: '18 Feb 24 23:02 PM',
                                Waiting: '15'
                            }
                            let text = encoder.align('center')
                                .width(3)
                                .height(3)
                                .line(data.ticket_name).bold()
                                .width(2)
                                .height(2)
                                .line(data.service).bold()
                                .width(2)
                                .height(2)
                                .line(`Waiting Count : ${data.Waiting}`)
                                .width(1)
                                .height(1)
                                .line(`Created At : ${data.timeDate}`)
                                .line("")
                                .line("")
                                .encode()
                            self.printCharacteristic.writeValue(text);
                            // self.device.gatt
                            //     .disconnect()
                        })
                    }
                    // self.sendTextData(device)
                })
                .catch(error => {
                    this.handleError(error, device)
                })
        },
        printlogo(startIndex, result, maxChunkSize) {
            return new Promise(async (resolve, reject) => {
                let dataSize = result.length
                while (startIndex < dataSize) {
                    const endIndex = Math.min(startIndex + maxChunkSize, dataSize);
                    const chunk = result.slice(startIndex, endIndex);
                    // Send the current chunk to the printer
                    await this.printCharacteristic.writeValue(chunk);
                    startIndex = endIndex;
                }
                resolve()
            })
        },
        connectBluetoothPrinter() {
            navigator.bluetooth
                .requestDevice({
                    filters: [
                        {
                            name: 'MPT-III',
                            services: ['000018f0-0000-1000-8000-00805f9b34fb']
                        }
                    ]
                },
                    {
                        optionalServices: ['00002af1-0000-1000-8000-00805f9b34fb']
                    })
                .then(device => {
                    console.log('device', device)
                    // window.alert(device)
                    this.device = device
                })
                .catch(this.handleError)
        },
        // connect(device) {
        //     const self = this
        //     device.addEventListener('gattserverdisconnected', this.onDisconnected)
        //     return device.gatt
        //         .connect()
        //         .then(server =>
        //             server.getPrimaryService('000018f0-0000-1000-8000-00805f9b34fb')
        //         )
        //         .then(service =>
        //             service.getCharacteristic('00002af1-0000-1000-8000-00805f9b34fb')
        //         )
        //         .then(characteristic => {
        //             console.log('characteristic', characteristic)
        //             self.printCharacteristic = characteristic
        //             let encoder = new EscPosEncoder({
        //                 imageMode: 'raster'
        //             });

        //             // let result = encoder
        //             //     .initialize()
        //             //     .text('The quick brown fox jumps over the lazy dog')
        //             //     .newline()
        //             //     .qrcode('https://nielsleenheer.com')
        //             //     .encode();

        //             console.log("router", window.location);
        //             const url = new URL(window.location.origin);

        //             // Get the IP address from the URL
        //             const ipAddress = url.hostname;
        //             let img = new Image();
        //             if (ipAddress === "localhost") {
        //                 return
        //             }
        //             img.src = `https://${ipAddress}:443/uploaded/logo.png`;

        //             img.onload = async function () {
        //                 let result = encoder.align('center')
        //                     .image(img, 120, 120, 'atkinson').newline().newline()
        //                     .encode()
        //                 console.log("hello", result);
        //                 const dataSize = result.length;
        //                 let startIndex = 0;
        //                 const maxChunkSize = characteristic.properties.writeWithoutResponse ? 512 : 20;
        //                 // Send chunks of data until the entire image is sent
        //                 while (startIndex < dataSize) {
        //                     const endIndex = Math.min(startIndex + maxChunkSize, dataSize);
        //                     const chunk = result.slice(startIndex, endIndex);

        //                     // Send the current chunk to the printer

        //                     await characteristic.writeValue(chunk);

        //                     startIndex = endIndex;
        //                 }

        //                 // this.loop(0, result, device)
        //                 // characteristic.writeValue(result)
        //             }
        //             // self.sendTextData(device)
        //         })
        //         .catch(error => {
        //             this.handleError(error, device)
        //         })
        // },
        // handleError(error, device) {
        //     console.error('handleError => error', error)
        //     if (device != null) {
        //         device.gatt.disconnect()
        //     }
        //     let erro = JSON.stringify({
        //         code: error.code,
        //         message: error.message,
        //         name: error.name
        //     })

        //     console.log('handleError => erro', erro)
        //     if (error.code !== 8) {
        //         alert('Could not connect with the printer. Try it again')
        //     }
        // },
        // getBytes(text) {
        //     console.log('text', text)
        //     let br = '\u000A\u000D'
        //     text = text === undefined ? br : text
        //     let replaced = text
        //     console.log('replaced', replaced)
        //     let bytes = new TextEncoder('utf-8').encode(replaced + br)
        //     console.log('bytes', bytes)
        //     return bytes
        // },
        // addText(arrayText) {
        //     let text = this.msg
        //     arrayText.push(text)
        // },
        // sendTextData(device) {
        //     let arrayText = []
        //     this.addText(arrayText)
        //     console.log('sendTextData => arrayText', arrayText)
        //     this.loop(0, arrayText, device)
        // },
        // loop(i, arrayText, device) {
        //     let arrayBytes = this.getBytes(arrayText[i])
        //     this.write(device, arrayBytes, () => {
        //         i++
        //         if (i < arrayText.length) {
        //             this.loop(i, arrayText, device)
        //         } else {
        //             let arrayBytes = this.getBytes()
        //             this.write(device, arrayBytes, () => {
        //                 device.gatt.disconnect()
        //             })
        //         }
        //     })
        // },
        // write(device, array, callback) {
        //     this.printCharacteristic
        //         .writeValue(array)
        //         .then(() => {
        //             console.log('Printed Array: ' + array.length)
        //             setTimeout(() => {
        //                 if (callback) {
        //                     callback()
        //                 }
        //             }, 250)
        //         })
        //         .catch(error => {
        //             this.handleError(error, device)
        //         })
        // },
        // async fileToBytes(file) {
        //     return new Promise((resolve, reject) => {
        //         const reader = new FileReader();

        //         reader.onload = (event) => {
        //             const arrayBuffer = event.target.result;
        //             const byteArray = new Uint8Array(arrayBuffer);
        //             resolve(byteArray);
        //         };

        //         reader.onerror = (error) => {
        //             reject(error);
        //         };

        //         reader.readAsArrayBuffer(file);
        //     });
        // },
        // async printFile(event) {
        //     const selectedFile = event.target.files[0];

        //     if (selectedFile) {
        //         try {
        //             // Convert file content to byte array
        //             const byteArray = await this.fileToBytes(selectedFile);

        //             // Use Web Bluetooth API to connect to your thermal printer
        //             const device = await navigator.bluetooth.requestDevice({
        //                 filters: [{ services: ['000018f0-0000-1000-8000-00805f9b34fb'] }],
        //             });

        //             const server = await device.gatt.connect();
        //             const service = await server.getPrimaryService('000018f0-0000-1000-8000-00805f9b34fb');
        //             const characteristic = await service.getCharacteristic("00002af0-0000-1000-8000-00805f9b34fb");
        //             const maxChunkSize = characteristic.properties.writeWithoutResponse ? 512 : 20;
        //             const bitmapdata = this.convertByteArrayToBitmap(byteArray, 250, 250)
        //             // this.sendBitmapToPrinter(bitmapdata, maxChunkSize, characteristic)
        //             // Split the byte array into chunks and send them sequentially
        //             // for (let offset = 0; offset < byteArray.length; offset += maxChunkSize) {
        //             //     const chunk = byteArray.slice(offset, offset + maxChunkSize);
        //             //     
        //             //     break
        //             // }
        //             const dataSize = bitmapdata.length;
        //             let startIndex = 0;

        //             // Send chunks of data until the entire image is sent
        //             while (startIndex < dataSize) {
        //                 const endIndex = Math.min(startIndex + maxChunkSize, dataSize);
        //                 const chunk = bitmapdata.slice(startIndex, endIndex);

        //                 // Send the current chunk to the printer

        //                 await characteristic.writeValue(chunk);

        //                 startIndex = endIndex;
        //             }

        //             console.log("file printing characteristics", characteristic);
        //             // Send the byte array to the printer
        //             // await characteristic.writeValue(byteArray);

        //             console.log('File printed successfully!');
        //         } catch (error) {
        //             console.error('Error printing file:', error);
        //         }
        //     }
        // },
        // convertByteArrayToBitmap(byteArray, width, height) {
        //     const dataSize = width * height;
        //     const bitmapData = new Uint8Array(dataSize);

        //     // Iterate through each byte in the byte array
        //     for (let i = 0; i < dataSize; i++) {
        //         // Check if the byte is greater than a threshold (e.g., 128)
        //         // and set the corresponding pixel in the bitmap data
        //         bitmapData[i] = byteArray[i] > 128 ? 255 : 0;
        //     }

        //     // Return the bitmap data
        //     return bitmapData;
        // },
        // async sendBitmapToPrinter(bitmapData, chunkSize, characteristic) {

        // }
    },
    async created() {
        // await this.listBluetoothDevices()
    }
}
</script>