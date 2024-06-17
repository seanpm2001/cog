<?php

namespace Grafana\Foundation\Types\StructOptionalFields;

final class SomeStructOperator implements \JsonSerializable, \Stringable {
    /**
     * @var string|int
     */
    private $value;

    /**
     * @var array<string, SomeStructOperator>
     */
    private static $instances = [];

    private function __construct(string|int $value)
    {
        $this->value = $value;
    }

    public static function greaterThan(): self
    {
        if (!isset(self::$instances["GreaterThan"])) {
            self::$instances["GreaterThan"] = new self(">");
        }

        return self::$instances["GreaterThan"];
    }

    public static function lessThan(): self
    {
        if (!isset(self::$instances["LessThan"])) {
            self::$instances["LessThan"] = new self("<");
        }

        return self::$instances["LessThan"];
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

