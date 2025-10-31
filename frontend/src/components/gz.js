// Gzip 压缩函数
async function gzipCompress(str) {
    const encoder = new TextEncoder();
    const data = encoder.encode(str);

    // 使用 pipeThrough 自动处理流控制
    const compressedReadable = new ReadableStream({
        start(controller) {
            controller.enqueue(data);
            controller.close();
        }
    }).pipeThrough(new CompressionStream('gzip'));

    const compressedData = await new Response(compressedReadable).arrayBuffer();
    return new Uint8Array(compressedData);
}

// Gzip 解压缩函数
async function gzipDecompress(uint8Array) {
    // 使用 pipeThrough 自动处理流控制
    const compressedReadable = new ReadableStream({
        start(controller) {
            controller.enqueue(uint8Array);
            controller.close();
        }
    }).pipeThrough(new DecompressionStream('gzip'));

    const decompressedData = await new Response(compressedReadable).arrayBuffer();
    // 将 Uint8Array 转换为字符串
    const decoder = new TextDecoder();
    return decoder.decode(decompressedData);
}

// 将 Uint8Array 转换为 Base64 字符串（用于显示）
function uint8ArrayToBase64(uint8Array) {
    let binary = '';
    const len = uint8Array.byteLength;
    for (let i = 0; i < len; i++) {
        binary += String.fromCharCode(uint8Array[i]);
    }
    return btoa(binary);
}

// 将 Base64 字符串转换为 Uint8Array
function base64ToUint8Array(base64) {
    const binaryString = atob(base64);
    const len = binaryString.length;
    const bytes = new Uint8Array(len);
    for (let i = 0; i < len; i++) {
        bytes[i] = binaryString.charCodeAt(i);
    }
    return bytes;
}

export const strToGz = async (str) => {
    const gz = await gzipCompress(str);
    return uint8ArrayToBase64(gz);
};
export const gzToStr = async (b64) => {
    const bytes = base64ToUint8Array(b64);
    return await gzipDecompress(bytes);
};
