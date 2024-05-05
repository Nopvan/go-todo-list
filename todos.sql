-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 05 Bulan Mei 2024 pada 18.43
-- Versi server: 10.4.28-MariaDB
-- Versi PHP: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `todos`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `scopes`
--

CREATE TABLE `scopes` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `scopes`
--

INSERT INTO `scopes` (`id`, `name`) VALUES
(2, 'todos:delete'),
(3, 'todos:read'),
(4, 'todos:update'),
(5, 'todos:create');

-- --------------------------------------------------------

--
-- Struktur dari tabel `todos`
--

CREATE TABLE `todos` (
  `id` int(11) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `done` tinyint(4) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `todos`
--

INSERT INTO `todos` (`id`, `title`, `description`, `done`, `user_id`) VALUES
(3, 'test', 'testttt', 0, NULL),
(7, 'jalan jalan', 'kemana kek', 1, 2),
(8, 'test scopes ajah', 'test scopes', 0, 2);

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`) VALUES
(2, 'jamal', 'jamal1@gmail.com', '$2a$10$ID41LbCjBmYkzSc0sAXlg.llN1j12e9dfIormCEmkS5TXilWwux1y'),
(3, 'diluc', 'diluc@gmail.com', '$2a$10$2AFkOpPwz4cv37EpZXucPOUXY/tFUaveAQufXN5MZn6EtTh.V6c4q');

-- --------------------------------------------------------

--
-- Struktur dari tabel `user_scopes`
--

CREATE TABLE `user_scopes` (
  `id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `scope_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `user_scopes`
--

INSERT INTO `user_scopes` (`id`, `user_id`, `scope_id`) VALUES
(2, 2, 3),
(3, 2, 4),
(4, 2, 2),
(5, 2, 5);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `scopes`
--
ALTER TABLE `scopes`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `todos`
--
ALTER TABLE `todos`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `user_scopes`
--
ALTER TABLE `user_scopes`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_roles_ibfk_1` (`user_id`),
  ADD KEY `user_roles_ibfk_2` (`scope_id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `scopes`
--
ALTER TABLE `scopes`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT untuk tabel `todos`
--
ALTER TABLE `todos`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT untuk tabel `user_scopes`
--
ALTER TABLE `user_scopes`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `user_scopes`
--
ALTER TABLE `user_scopes`
  ADD CONSTRAINT `user_scopes_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `user_scopes_ibfk_2` FOREIGN KEY (`scope_id`) REFERENCES `scopes` (`id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
