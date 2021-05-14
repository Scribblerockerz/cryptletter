import CryptoJS from 'crypto-js';
import encodingUTF8 from 'crypto-js/enc-utf8';

export async function encryptFileInput(fileInput, encryptionKey) {
    return new Promise((resolve) => {
        console.log(`Reading ${fileInput.name} with size of ${fileInput.size}`);
        const mimeType = fileInput.type;

        const reader = new FileReader();
        reader.onload = (e) => {
            const fileData = e.target.result;
            console.log('onload finished', fileData);

            // Create a WordArray out of ArrayBuffer
            let wordBuffer = CryptoJS.lib.WordArray.create(fileData);

            // Encrypt data
            const encryptedData = encrypt(wordBuffer, encryptionKey);

            // Transform data to plain string (used for transfer to server)
            const encryptedDataAsString = encryptedData.toString();
            const encryptedName = encryptString(fileInput.name, encryptionKey);
            const encryptedMimeType = encryptString(mimeType, encryptionKey);
            const encryptedSize = encryptString(
                `${fileInput.size}`,
                encryptionKey
            );

            resolve({
                name: fileInput.name,
                encryptedName,
                mimeType: mimeType,
                encryptedMimeType,
                size: fileInput.size,
                encryptedSize,
                data: encryptedDataAsString,
            });
        };
        reader.readAsArrayBuffer(fileInput);
    });
}

export function decryptFileData(encryptedFileData, encryptionKey) {
    // Decrypt data
    const decryptedData = decrypt(encryptedFileData, encryptionKey);

    // Convert WordArray to Uint8Array
    const uint8buffer = wordArrayToUint8Array(decryptedData);

    return uint8buffer;

    // // Create a new file from Uint8Array
    // const fakeFile = new File([uint8buffer], 'new-file.jpg', {
    //     type: mimeType,
    // });
}

// function handleFile(file) {
//     console.log(`Reading ${file.name} with size of ${file.size}`);
//     const mimeType = file.type;

//     const reader = new FileReader();
//     reader.onload = (e) => {
//         const fileData = e.target.result;
//         console.log('onload finished', fileData);

//         // Create a WordArray out of ArrayBuffer
//         let wordBuffer = CryptoJS.lib.WordArray.create(fileData);

//         // Encrypt data
//         const encryptedData = encrypt(wordBuffer, ENCRYPTION_KEY);

//         console.log(humanFileSize(encryptedData.ciphertext.sigBytes, true));

//         // Transform data to plain string (used for transfer to server)
//         const encryptedDataAsString = encryptedData.toString();

//         // Decrypt data
//         const decryptedData = decrypt(encryptedDataAsString, ENCRYPTION_KEY);

//         // Convert WordArray to Uint8Array
//         const uint8buffer = wordArrayToUint8Array(decryptedData);

//         // Create a new file from Uint8Array
//         const fakeFile = new File([uint8buffer], 'new-file.jpg', {
//             type: mimeType,
//         });

//         // Preview file
//         previewFile(fakeFile);
//     };
//     reader.readAsArrayBuffer(file);
// }

// https://stackoverflow.com/a/35536933/1704139
function encrypt(data, key) {
    return CryptoJS.AES.encrypt(data, key);
}

function encryptString(data, key) {
    return encrypt(data, key).toString();
}

function decrypt(encryptedData, key) {
    return CryptoJS.AES.decrypt(encryptedData, key);
}

export function decryptString(encryptedData, key) {
    return decrypt(encryptedData, key).toString(encodingUTF8);
}

// function previewFile(file) {
//     const isDataURI = typeof file == 'string';
//     const preview = document.querySelector('#preview');
//     const img = document.createElement('img');
//     if (!isDataURI) {
//         img.file = file;
//     }

//     preview.appendChild(img);

//     if (isDataURI) {
//         img.src = file;
//         return;
//     }

//     const reader = new FileReader();
//     reader.onload = (function (aImg) {
//         return function (e) {
//             aImg.src = e.target.result;
//         };
//     })(img);
//     reader.readAsDataURL(file);
// }

/*
 * Converts a cryptjs WordArray to native Uint8Array
 * https://github.com/brix/crypto-js/issues/274
 */
function wordArrayToUint8Array(wordArray) {
    const l = wordArray.sigBytes;
    const words = wordArray.words;
    const result = new Uint8Array(l);
    var i = 0 /*dst*/,
        j = 0; /*src*/

    // eslint-disable-next-line no-constant-condition
    while (true) {
        // here i is a multiple of 4
        if (i == l) break;
        var w = words[j++];
        result[i++] = (w & 0xff000000) >>> 24;
        if (i == l) break;
        result[i++] = (w & 0x00ff0000) >>> 16;
        if (i == l) break;
        result[i++] = (w & 0x0000ff00) >>> 8;
        if (i == l) break;
        result[i++] = w & 0x000000ff;
    }

    return result;
}

/*

Concept:

  POST https://example.com/
  {
      id: [uuid]
      message: "[encrypted]",
      ttl: [time to live in minutes],
      files: [
          {
              name: "[encrypted file name]",
              mimeType: "[encrypted file mimeType]",
              data: "[encrypted file binary]"
          },
          {
              name: "[encrypted file name]",
              mimeType: "[encrypted file mimeType]",
              data: "[encrypted file binary]"
          },
      ]
  }


  --> SAVE
  <-- "Here is the url https://example.com/[uuid]/#/[encryption_key]/


  GET https://example.com/[uuid]/#/[encryption_key]/
  {
      id: [uuid]
      message: "[encrypted]",
      ttl: [time to live in minutes],
      files: [
          {
              name: "[encrypted file name]",
              mimeType: "[encrypted file mimeType]",
              fileId: "[uuid] to request the binary"
          },
          {
              name: "[encrypted file name]",
              mimeType: "[encrypted file mimeType]",
              fileId: "[uuid] to request the binary"
          },
      ]
  }

  [download file] -> GET https://example.com/file/[fileId]/
  <-- FILE BINARY

  new File([FILE BINARY]) -> download this thing!


  https://stackoverflow.com/a/17696608/1704139
  var blob = new Blob([xhr.response], {type: "application/pdf"});
  var objectUrl = URL.createObjectURL(blob);
  window.open(objectUrl);













  */
