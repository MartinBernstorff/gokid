# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Gokid < Formula
  desc "A simple CLI to submit changes"
  homepage "https://github.com/martinbernstorff/gokid"
  version "0.3.2"
  license "MIT"

  on_macos do
    on_intel do
      url "https://github.com/MartinBernstorff/gokid/releases/download/v0.3.2/gokid_Darwin_x86_64.tar.gz"
      sha256 "4ffc09938369e9d403ad9ec20556d03d9fb109a94355db2b15afbd903f02213f"

      def install
        bin.install "gokid"
      end
    end
    on_arm do
      url "https://github.com/MartinBernstorff/gokid/releases/download/v0.3.2/gokid_Darwin_arm64.tar.gz"
      sha256 "eafd2045e79d5dd02bca0d41ea6c4c329696e91b1aaea91b11fe8700faa956c5"

      def install
        bin.install "gokid"
      end
    end
  end

  on_linux do
    on_intel do
      if Hardware::CPU.is_64_bit?
        url "https://github.com/MartinBernstorff/gokid/releases/download/v0.3.2/gokid_Linux_x86_64.tar.gz"
        sha256 "0138b8fe16cc1baa21376377332cbc3c4fc5f8ee34b288719928e6f3e121d0a4"

        def install
          bin.install "gokid"
        end
      end
    end
    on_arm do
      if Hardware::CPU.is_64_bit?
        url "https://github.com/MartinBernstorff/gokid/releases/download/v0.3.2/gokid_Linux_arm64.tar.gz"
        sha256 "2d5cd92a1873824eea51f592b4c5142525b6bd4fca09b07b8ab8a390cf4f4f05"

        def install
          bin.install "gokid"
        end
      end
    end
  end
end
