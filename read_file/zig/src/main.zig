const std = @import("std");

fn r1() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();

    var file = try std.fs.cwd().openFile("../input.txt", .{});
    defer file.close();

    var buffered = std.io.bufferedReader(file.reader());
    var reader = buffered.reader();

    var arr = std.ArrayList(u8).init(allocator);
    defer arr.deinit();

    var lines: usize = 0;

    while (true) {
        reader.streamUntilDelimiter(arr.writer(), '\n', null) catch |err| switch (err) {
            error.EndOfStream => break,
            else => return err,
        };

        lines += 1;

        arr.clearRetainingCapacity();
    }

    std.debug.print("{d}\n", .{lines});
}

fn r2() !void {
    const file = try std.fs.cwd().openFile("../input.txt", .{});
    defer file.close();

    var lines: usize = 0;

    var buf_reader = std.io.bufferedReader(file.reader());
    var in_stream = buf_reader.reader();

    var buf: [1024]u8 = undefined;
    while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        _ = line;
        lines += 1;
    }

    std.debug.print("{d}\n", .{lines});
}

fn r3() anyerror!void {
    const allocator = std.heap.page_allocator;

    // Open a file for reading
    var file = try std.fs.cwd().openFile("../input.txt", .{});

    // Read the entire contents into a buffer
    var contents = try file.readToEndAlloc(allocator, std.math.maxInt(usize));

    // Split the contents into lines
    const content = std.mem.split(u8, &contents, "\n");

    var lines: usize = 0;

    // Iterate over lines
    for (content) |_| {
        lines += 1;
    }

    // Clean up
    allocator.free(contents);
}

pub fn main() !void {
    try r3();
}
