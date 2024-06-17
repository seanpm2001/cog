<?php

namespace Grafana\Foundation\Types\Enums;

/**
 * 0 for no shared crosshair or tooltip (default).
 * 1 for shared crosshair.
 * 2 for shared crosshair AND shared tooltip.
 */
final class DashboardCursorSync implements \JsonSerializable, \Stringable {
    /**
     * @var string|int
     */
    private $value;

    /**
     * @var array<string, DashboardCursorSync>
     */
    private static $instances = [];

    private function __construct(string|int $value)
    {
        $this->value = $value;
    }

    public static function off(): self
    {
        if (!isset(self::$instances["Off"])) {
            self::$instances["Off"] = new self(0);
        }

        return self::$instances["Off"];
    }

    public static function crosshair(): self
    {
        if (!isset(self::$instances["Crosshair"])) {
            self::$instances["Crosshair"] = new self(1);
        }

        return self::$instances["Crosshair"];
    }

    public static function tooltip(): self
    {
        if (!isset(self::$instances["Tooltip"])) {
            self::$instances["Tooltip"] = new self(2);
        }

        return self::$instances["Tooltip"];
    }

    public function jsonSerialize(): string|int
    {
        return $this->value;
    }

    public function __toString(): string
    {
        return (string) $this->value;
    }
}

