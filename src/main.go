package main

import (
    "container/list"
    "strings"
)

// Node merepresentasikan vertex/titik dalam graf
// Value adalah nilai dari node tersebut
// Neighbors adalah daftar node yang terhubung dengan node ini
type Node struct {
    Value     string
    Neighbors []string
}

// BreadthFirstSearch melakukan pencarian jalur dari root ke target menggunakan algoritma BFS
// Mengembalikan string berupa jalur yang ditemukan atau "not_found" jika tidak ada jalur
func BreadthFirstSearch(tree map[string]Node, root, target string) string {
    const notFound = "not_found"

    // Memeriksa apakah root dan target ada dalam graf
    rootNode, rootExists := tree[root]
    _, targetExists := tree[target]
    if !rootExists || !targetExists {
        return notFound
    }

    // Membuat antrian untuk BFS menggunakan double linked list
    q := list.New()
    // Masukkan node awal (root) ke dalam antrian
    q.PushBack(rootNode)

    // Map untuk menyimpan parent dari setiap node yang dikunjungi
    // Juga berfungsi sebagai penanda node yang sudah dikunjungi
    parents := make(map[string]string)
    parents[root] = ""

    // Selama antrian tidak kosong, lanjutkan pencarian
    for q.Len() > 0 {
        // Ambil node paling depan dari antrian
        currentNode := q.Front().Value.(Node)
        q.Remove(q.Front())

        // Jika node yang sedang diperiksa adalah target yang dicari
        if strings.EqualFold(currentNode.Value, target) {
            // Membuat jalur dari root ke target dengan menelusuri parents
            var route []string
            for len(currentNode.Value) > 0 {
                // Masukkan node ke awal array untuk mendapatkan urutan yang benar
                route = append([]string{currentNode.Value}, route...)
                currentNode.Value = parents[currentNode.Value]
            }

            // Gabungkan jalur dengan tanda panah
            return strings.Join(route, "->")
        }

        // Periksa semua tetangga dari node saat ini
        for _, neighbor := range currentNode.Neighbors {
            // Jika tetangga belum pernah dikunjungi
            if _, visited := parents[neighbor]; !visited {
                // Catat parent dari tetangga ini
                parents[neighbor] = currentNode.Value
                // Masukkan tetangga ke dalam antrian
                q.PushBack(tree[neighbor])
            }
        }
    }

    // Jika target tidak ditemukan, kembalikan "not_found"
    return notFound
}