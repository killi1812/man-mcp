class ManMcp < Formula
  desc "Model Context Protocol (MCP) server for local system man pages"
  homepage "https://github.com/killi1812/man-mcp"
  url "https://github.com/killi1812/man-mcp/archive/refs/tags/v0.0.1.tar.gz"
  sha256 "0000000000000000000000000000000000000000000000000000000000000000" # Update on release
  license "GPL-3.0-or-later"

  depends_on "go" => :build

  def install
    cd "src" do
      system "go", "build", *std_go_args(ldflags: "-X github.com/killi1812/man-mcp/app.Build=prod -X github.com/killi1812/man-mcp/app.Version=#{version}")
    end
  end

  test do
    assert_match "man-mcp", shell_output("#{bin}/man-mcp --help")
  end
end
