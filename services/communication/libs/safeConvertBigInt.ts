export function safeToBigInt(s : string) {
    try {
        // Loại bỏ khoảng trắng đầu cuối
        const trimmed = s.trim()

        // Kiểm tra chuỗi rỗng
        if (!trimmed) {
            return null
        }

        // Kiểm tra format số hợp lệ bằng regex
        if (!/^-?\d+$/.test(trimmed)) {
            return null
        }

        // Convert sang BigInt
        const result = BigInt(trimmed)
        return result

    } catch (error) {
        return null
    }
}
