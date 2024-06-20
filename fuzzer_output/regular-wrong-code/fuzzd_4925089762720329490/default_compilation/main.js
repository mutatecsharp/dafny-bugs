// Dafny program main.dfy compiled into JavaScript
// Copyright by the contributors to the Dafny Project
// SPDX-License-Identifier: MIT

const BigNumber = require('bignumber.js');
BigNumber.config({ MODULO_MODE: BigNumber.EUCLID })
let _dafny = (function() {
  let $module = {};
  $module.areEqual = function(a, b) {
    if (typeof a === 'string' && b instanceof _dafny.Seq) {
      // Seq.equals(string) works as expected,
      // and the catch-all else block handles that direction.
      // But the opposite direction doesn't work; handle it here.
      return b.equals(a);
    } else if (typeof a === 'number' && BigNumber.isBigNumber(b)) {
      // This conditional would be correct even without the `typeof a` part,
      // but in most cases it's probably faster to short-circuit on a `typeof`
      // than to call `isBigNumber`. (But it remains to properly test this.)
      return b.isEqualTo(a);
    } else if (typeof a !== 'object' || a === null || b === null) {
      return a === b;
    } else if (BigNumber.isBigNumber(a)) {
      return a.isEqualTo(b);
    } else if (a._tname !== undefined || (Array.isArray(a) && a.constructor.name == "Array")) {
      return a === b;  // pointer equality
    } else {
      return a.equals(b);  // value-type equality
    }
  }
  $module.toString = function(a) {
    if (a === null) {
      return "null";
    } else if (typeof a === "number") {
      return a.toFixed();
    } else if (BigNumber.isBigNumber(a)) {
      return a.toFixed();
    } else if (a._tname !== undefined) {
      return a._tname;
    } else {
      return a.toString();
    }
  }
  $module.escapeCharacter = function(cp) {
    let s = String.fromCodePoint(cp.value)
    switch (s) {
      case '\n': return "\\n";
      case '\r': return "\\r";
      case '\t': return "\\t";
      case '\0': return "\\0";
      case '\'': return "\\'";
      case '\"': return "\\\"";
      case '\\': return "\\\\";
      default: return s;
    };
  }
  $module.NewObject = function() {
    return { _tname: "object" };
  }
  $module.InstanceOfTrait = function(obj, trait) {
    return obj._parentTraits !== undefined && obj._parentTraits().includes(trait);
  }
  $module.Rtd_bool = class {
    static get Default() { return false; }
  }
  $module.Rtd_char = class {
    static get Default() { return 'D'; }  // See CharType.DefaultValue in Dafny source code
  }
  $module.Rtd_codepoint = class {
    static get Default() { return new _dafny.CodePoint('D'.codePointAt(0)); }
  }
  $module.Rtd_int = class {
    static get Default() { return BigNumber(0); }
  }
  $module.Rtd_number = class {
    static get Default() { return 0; }
  }
  $module.Rtd_ref = class {
    static get Default() { return null; }
  }
  $module.Rtd_array = class {
    static get Default() { return []; }
  }
  $module.ZERO = new BigNumber(0);
  $module.ONE = new BigNumber(1);
  $module.NUMBER_LIMIT = new BigNumber(0x20).multipliedBy(0x1000000000000);  // 2^53
  $module.Tuple = class Tuple extends Array {
    constructor(...elems) {
      super(...elems);
    }
    toString() {
      return "(" + arrayElementsToString(this) + ")";
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static Default(...values) {
      return Tuple.of(...values);
    }
    static Rtd(...rtdArgs) {
      return {
        Default: Tuple.from(rtdArgs, rtd => rtd.Default)
      };
    }
  }
  $module.Set = class Set extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Set.Empty;
    }
    toString() {
      return "{" + arrayElementsToString(this) + "}";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Set();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new Set();
      for (let k of elmts) {
        s.add(k);
      }
      return s;
    }
    contains(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i], k)) {
          return true;
        }
      }
      return false;
    }
    add(k) {  // mutates the Set; use only during construction
      if (!this.contains(k)) {
        this.push(k);
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        if (!other.contains(e)) {
          return false;
        }
      }
      return true;
    }
    get Elements() {
      return this;
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = Set.of(...this);
        for (let k of that) {
          s.add(k);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new Set();
        for (let k of this) {
          if (that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length == 0 || that.length == 0) {
        return this;
      } else {
        let s = new Set();
        for (let k of this) {
          if (!that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      for (let k of this) {
        if (that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsSubsetOf(that) {
      if (that.length < this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      if (that.length <= this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    get AllSubsets() {
      return this.AllSubsets_();
    }
    *AllSubsets_() {
      // Start by putting all set elements into a list, but don't include null
      let elmts = Array.of(...this);
      let n = elmts.length;
      let which = new Array(n);
      which.fill(false);
      let a = [];
      while (true) {
        yield Set.of(...a);
        // "add 1" to "which", as if doing a carry chain.  For every digit changed, change the membership of the corresponding element in "a".
        let i = 0;
        for (; i < n && which[i]; i++) {
          which[i] = false;
          // remove elmts[i] from a
          for (let j = 0; j < a.length; j++) {
            if (_dafny.areEqual(a[j], elmts[i])) {
              // move the last element of a into slot j
              a[j] = a[-1];
              a.pop();
              break;
            }
          }
        }
        if (i === n) {
          // we have cycled through all the subsets
          break;
        }
        which[i] = true;
        a.push(elmts[i]);
      }
    }
  }
  $module.MultiSet = class MultiSet extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return MultiSet.Empty;
    }
    toString() {
      let s = "multiset{";
      let sep = "";
      for (let e of this) {
        let [k, n] = e;
        let ks = _dafny.toString(k);
        while (!n.isZero()) {
          n = n.minus(1);
          s += sep + ks;
          sep = ", ";
        }
      }
      s += "}";
      return s;
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new MultiSet();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new MultiSet();
      for (let e of elmts) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    static FromArray(arr) {
      let s = new MultiSet();
      for (let e of arr) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    cardinality() {
      let c = _dafny.ZERO;
      for (let e of this) {
        let [k, n] = e;
        c = c.plus(n);
      }
      return c;
    }
    clone() {
      let s = new MultiSet();
      for (let e of this) {
        let [k, n] = e;
        s.push([k, n]);  // make sure to create a new array [k, n] here
      }
      return s;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return _dafny.ZERO;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return !this.get(k).isZero();
    }
    add(k, n) {
      let i = this.findIndex(k);
      if (i === this.length) {
        this.push([k, n]);
      } else {
        let m = this[i][1];
        this[i] = [k, m.plus(n)];
      }
    }
    update(k, n) {
      let i = this.findIndex(k);
      if (i < this.length && this[i][1].isEqualTo(n)) {
        return this;
      } else if (i === this.length && n.isZero()) {
        return this;
      } else if (i === this.length) {
        let m = this.slice();
        m.push([k, n]);
        return m;
      } else {
        let m = this.slice();
        m[i] = [k, n];
        return m;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let e of this) {
        let [k, n] = e;
        let m = other.get(k);
        if (!n.isEqualTo(m)) {
          return false;
        }
      }
      return this.cardinality().isEqualTo(other.cardinality());
    }
    get Elements() {
      return this.Elements_();
    }
    *Elements_() {
      for (let i = 0; i < this.length; i++) {
        let [k, n] = this[i];
        while (!n.isZero()) {
          yield k;
          n = n.minus(1);
        }
      }
    }
    get UniqueElements() {
      return this.UniqueElements_();
    }
    *UniqueElements_() {
      for (let e of this) {
        let [k, n] = e;
        if (!n.isZero()) {
          yield k;
        }
      }
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = this.clone();
        for (let e of that) {
          let [k, n] = e;
          s.add(k, n);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let m = that.get(k);
          if (!m.isZero()) {
            s.push([k, m.isLessThan(n) ? m : n]);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length === 0 || that.length === 0) {
        return this;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let d = n.minus(that.get(k));
          if (d.isGreaterThan(0)) {
            s.push([k, d]);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      let intersection = this.Intersect(that);
      return intersection.cardinality().isZero();
    }
    IsSubsetOf(that) {
      for (let e of this) {
        let [k, n] = e;
        let m = that.get(k);
        if (!n.isLessThanOrEqualTo(m)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      return this.IsSubsetOf(that) && this.cardinality().isLessThan(that.cardinality());
    }
  }
  $module.CodePoint = class CodePoint {
    constructor(value) {
      this.value = value
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      return this.value === other.value
    }
    isLessThan(other) {
      return this.value < other.value
    }
    isLessThanOrEqual(other) {
      return this.value <= other.value
    }
    toString() {
      return "'" + $module.escapeCharacter(this) + "'";
    }
    static isCodePoint(i) {
      return (
        (_dafny.ZERO.isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0xD800))) ||
        (new BigNumber(0xE000).isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0x11_0000))))
    }
  }
  $module.Seq = class Seq extends Array {
    constructor(...elems) {
      super(...elems);
    }
    static get Default() {
      return Seq.of();
    }
    static Create(n, init) {
      return Seq.from({length: n}, (_, i) => init(new BigNumber(i)));
    }
    static UnicodeFromString(s) {
      return new Seq(...([...s].map(c => new _dafny.CodePoint(c.codePointAt(0)))))
    }
    toString() {
      return "[" + arrayElementsToString(this) + "]";
    }
    toVerbatimString(asLiteral) {
      if (asLiteral) {
        return '"' + this.map(c => _dafny.escapeCharacter(c)).join("") + '"';
      } else {
        return this.map(c => String.fromCodePoint(c.value)).join("");
      }
    }
    static update(s, i, v) {
      if (typeof s === "string") {
        let p = s.slice(0, i);
        let q = s.slice(i.toNumber() + 1);
        return p.concat(v, q);
      } else {
        let t = s.slice();
        t[i] = v;
        return t;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static contains(s, k) {
      if (typeof s === "string") {
        return s.includes(k);
      } else {
        for (let x of s) {
          if (_dafny.areEqual(x, k)) {
            return true;
          }
        }
        return false;
      }
    }
    get Elements() {
      return this;
    }
    get UniqueElements() {
      return _dafny.Set.fromElements(...this);
    }
    static Concat(a, b) {
      if (typeof a === "string" || typeof b === "string") {
        // string concatenation, so make sure both operands are strings before concatenating
        if (typeof a !== "string") {
          // a must be a Seq
          a = a.join("");
        }
        if (typeof b !== "string") {
          // b must be a Seq
          b = b.join("");
        }
        return a + b;
      } else {
        // ordinary concatenation
        let r = Seq.of(...a);
        r.push(...b);
        return r;
      }
    }
    static JoinIfPossible(x) {
      try { return x.join(""); } catch(_error) { return x; }
    }
    static IsPrefixOf(a, b) {
      if (b.length < a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
    static IsProperPrefixOf(a, b) {
      if (b.length <= a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
  }
  $module.Map = class Map extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Map.of();
    }
    toString() {
      return "map[" + this.map(maplet => _dafny.toString(maplet[0]) + " := " + _dafny.toString(maplet[1])).join(", ") + "]";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Map();
      }
      return this._empty;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return undefined;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return this.findIndex(k) < this.length;
    }
    update(k, v) {
      let m = this.slice();
      m.updateUnsafe(k, v);
      return m;
    }
    // Similar to update, but make the modification in-place.
    // Meant to be used in the map constructor.
    updateUnsafe(k, v) {
      let m = this;
      let i = m.findIndex(k);
      m[i] = [k, v];
      return m;
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        let [k, v] = e;
        let w = other.get(k);
        if (w === undefined || !_dafny.areEqual(v, w)) {
          return false;
        }
      }
      return true;
    }
    get Keys() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(k);
      }
      return s;
    }
    get Values() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.add(v);
      }
      return s;
    }
    get Items() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(_dafny.Tuple.of(k, v));
      }
      return s;
    }
    Merge(that) {
      let m = that.slice();
      for (let e of this) {
        let [k, v] = e;
        let i = m.findIndex(k);
        if (i == m.length) {
          m[i] = [k, v];
        }
      }
      return m;
    }
    Subtract(keys) {
      if (this.length === 0 || keys.length === 0) {
        return this;
      }
      let m = new Map();
      for (let e of this) {
        let [k, v] = e;
        if (!keys.contains(k)) {
          m[m.length] = e;
        }
      }
      return m;
    }
  }
  $module.newArray = function(initValue, ...dims) {
    return { dims: dims, elmts: buildArray(initValue, ...dims) };
  }
  $module.BigOrdinal = class BigOrdinal {
    static get Default() {
      return _dafny.ZERO;
    }
    static IsLimit(ord) {
      return ord.isZero();
    }
    static IsSucc(ord) {
      return ord.isGreaterThan(0);
    }
    static Offset(ord) {
      return ord;
    }
    static IsNat(ord) {
      return true;  // at run time, every ORDINAL is a natural number
    }
  }
  $module.BigRational = class BigRational {
    static get ZERO() {
      if (this._zero === undefined) {
        this._zero = new BigRational(_dafny.ZERO);
      }
      return this._zero;
    }
    constructor (n, d) {
      // requires d === undefined || 1 <= d
      this.num = n;
      this.den = d === undefined ? _dafny.ONE : d;
      // invariant 1 <= den || (num == 0 && den == 0)
    }
    static get Default() {
      return _dafny.BigRational.ZERO;
    }
    // We need to deal with the special case `num == 0 && den == 0`, because
    // that's what C#'s default struct constructor will produce for BigRational. :(
    // To deal with it, we ignore `den` when `num` is 0.
    toString() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num.toFixed() + ".0";
      }
      let answer = this.dividesAPowerOf10(this.den);
      if (answer !== undefined) {
        let n = this.num.multipliedBy(answer[0]);
        let log10 = answer[1];
        let sign, digits;
        if (this.num.isLessThan(0)) {
          sign = "-"; digits = n.negated().toFixed();
        } else {
          sign = ""; digits = n.toFixed();
        }
        if (log10 < digits.length) {
          let digitCount = digits.length - log10;
          return sign + digits.slice(0, digitCount) + "." + digits.slice(digitCount);
        } else {
          return sign + "0." + "0".repeat(log10 - digits.length) + digits;
        }
      } else {
        return "(" + this.num.toFixed() + ".0 / " + this.den.toFixed() + ".0)";
      }
    }
    isPowerOf10(x) {
      if (x.isZero()) {
        return undefined;
      }
      let log10 = 0;
      while (true) {  // invariant: x != 0 && x * 10^log10 == old(x)
        if (x.isEqualTo(1)) {
          return log10;
        } else if (x.mod(10).isZero()) {
          log10++;
          x = x.dividedToIntegerBy(10);
        } else {
          return undefined;
        }
      }
    }
    dividesAPowerOf10(i) {
      let factor = _dafny.ONE;
      let log10 = 0;
      if (i.isLessThanOrEqualTo(_dafny.ZERO)) {
        return undefined;
      }

      // invariant: 1 <= i && i * 10^log10 == factor * old(i)
      while (i.mod(10).isZero()) {
        i = i.dividedToIntegerBy(10);
       log10++;
      }

      while (i.mod(5).isZero()) {
        i = i.dividedToIntegerBy(5);
        factor = factor.multipliedBy(2);
        log10++;
      }
      while (i.mod(2).isZero()) {
        i = i.dividedToIntegerBy(2);
        factor = factor.multipliedBy(5);
        log10++;
      }

      if (i.isEqualTo(_dafny.ONE)) {
        return [factor, log10];
      } else {
        return undefined;
      }
    }
    toBigNumber() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num;
      } else if (this.num.isGreaterThan(0)) {
        return this.num.dividedToIntegerBy(this.den);
      } else {
        return this.num.minus(this.den).plus(1).dividedToIntegerBy(this.den);
      }
    }
    isInteger() {
      return this.equals(new _dafny.BigRational(this.toBigNumber(), _dafny.ONE));
    }
    // Returns values such that aa/dd == a and bb/dd == b.
    normalize(b) {
      let a = this;
      let aa, bb, dd;
      if (a.num.isZero()) {
        aa = a.num;
        bb = b.num;
        dd = b.den;
      } else if (b.num.isZero()) {
        aa = a.num;
        dd = a.den;
        bb = b.num;
      } else {
        let gcd = BigNumberGcd(a.den, b.den);
        let xx = a.den.dividedToIntegerBy(gcd);
        let yy = b.den.dividedToIntegerBy(gcd);
        // We now have a == a.num / (xx * gcd) and b == b.num / (yy * gcd).
        aa = a.num.multipliedBy(yy);
        bb = b.num.multipliedBy(xx);
        dd = a.den.multipliedBy(yy);
      }
      return [aa, bb, dd];
    }
    compareTo(that) {
      // simple things first
      let asign = this.num.isZero() ? 0 : this.num.isLessThan(0) ? -1 : 1;
      let bsign = that.num.isZero() ? 0 : that.num.isLessThan(0) ? -1 : 1;
      if (asign < 0 && 0 <= bsign) {
        return -1;
      } else if (asign <= 0 && 0 < bsign) {
        return -1;
      } else if (bsign < 0 && 0 <= asign) {
        return 1;
      } else if (bsign <= 0 && 0 < asign) {
        return 1;
      }
      let [aa, bb, dd] = this.normalize(that);
      if (aa.isLessThan(bb)) {
        return -1;
      } else if (aa.isEqualTo(bb)){
        return 0;
      } else {
        return 1;
      }
    }
    equals(that) {
      return this.compareTo(that) === 0;
    }
    isLessThan(that) {
      return this.compareTo(that) < 0;
    }
    isAtMost(that) {
      return this.compareTo(that) <= 0;
    }
    plus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.plus(bb), dd);
    }
    minus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.minus(bb), dd);
    }
    negated() {
      return new BigRational(this.num.negated(), this.den);
    }
    multipliedBy(b) {
      return new BigRational(this.num.multipliedBy(b.num), this.den.multipliedBy(b.den));
    }
    dividedBy(b) {
      let a = this;
      // Compute the reciprocal of b
      let bReciprocal;
      if (b.num.isGreaterThan(0)) {
        bReciprocal = new BigRational(b.den, b.num);
      } else {
        // this is the case b.num < 0
        bReciprocal = new BigRational(b.den.negated(), b.num.negated());
      }
      return a.multipliedBy(bReciprocal);
    }
  }
  $module.EuclideanDivisionNumber = function(a, b) {
    if (0 <= a) {
      if (0 <= b) {
        // +a +b: a/b
        return Math.floor(a / b);
      } else {
        // +a -b: -(a/(-b))
        return -Math.floor(a / -b);
      }
    } else {
      if (0 <= b) {
        // -a +b: -((-a-1)/b) - 1
        return -Math.floor((-a-1) / b) - 1;
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return Math.floor((-a-1) / -b) + 1;
      }
    }
  }
  $module.EuclideanDivision = function(a, b) {
    if (a.isGreaterThanOrEqualTo(0)) {
      if (b.isGreaterThanOrEqualTo(0)) {
        // +a +b: a/b
        return a.dividedToIntegerBy(b);
      } else {
        // +a -b: -(a/(-b))
        return a.dividedToIntegerBy(b.negated()).negated();
      }
    } else {
      if (b.isGreaterThanOrEqualTo(0)) {
        // -a +b: -((-a-1)/b) - 1
        return a.negated().minus(1).dividedToIntegerBy(b).negated().minus(1);
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return a.negated().minus(1).dividedToIntegerBy(b.negated()).plus(1);
      }
    }
  }
  $module.EuclideanModuloNumber = function(a, b) {
    let bp = Math.abs(b);
    if (0 <= a) {
      // +a: a % bp
      return a % bp;
    } else {
      // c = ((-a) % bp)
      // -a: bp - c if c > 0
      // -a: 0 if c == 0
      let c = (-a) % bp;
      return c === 0 ? c : bp - c;
    }
  }
  $module.ShiftLeft = function(b, n) {
    return b.multipliedBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.ShiftRight = function(b, n) {
    return b.dividedToIntegerBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.RotateLeft = function(b, n, w) {  // truncate(b << n) | (b >> (w - n))
    let x = _dafny.ShiftLeft(b, n).mod(new BigNumber(2).exponentiatedBy(w));
    let y = _dafny.ShiftRight(b, w - n);
    return x.plus(y);
  }
  $module.RotateRight = function(b, n, w) {  // (b >> n) | truncate(b << (w - n))
    let x = _dafny.ShiftRight(b, n);
    let y = _dafny.ShiftLeft(b, w - n).mod(new BigNumber(2).exponentiatedBy(w));;
    return x.plus(y);
  }
  $module.BitwiseAnd = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 & b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    return r;
  }
  $module.BitwiseOr = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 | b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseXor = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 ^ b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseNot = function(a, bits) {
    let r = _dafny.ZERO;
    let h = _dafny.ONE;
    for (let i = 0; i < bits; i++) {
      let bit = a.mod(2);
      if (bit.isZero()) {
        r = r.plus(h);
      }
      a = a.dividedToIntegerBy(2);
      h = h.multipliedBy(2);
    }
    return r;
  }
  $module.Quantifier = function(vals, frall, pred) {
    for (let u of vals) {
      if (pred(u) !== frall) { return !frall; }
    }
    return frall;
  }
  $module.PlusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) + b.charCodeAt(0));
  }
  $module.UnicodePlusChar = function(a, b) {
    return new _dafny.CodePoint(a.value + b.value);
  }
  $module.MinusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) - b.charCodeAt(0));
  }
  $module.UnicodeMinusChar = function(a, b) {
    return new _dafny.CodePoint(a.value - b.value);
  }
  $module.AllBooleans = function*() {
    yield false;
    yield true;
  }
  $module.AllChars = function*() {
    for (let i = 0; i < 0x10000; i++) {
      yield String.fromCharCode(i);
    }
  }
  $module.AllUnicodeChars = function*() {
    for (let i = 0; i < 0xD800; i++) {
      yield new _dafny.CodePoint(i);
    }
    for (let i = 0xE0000; i < 0x110000; i++) {
      yield new _dafny.CodePoint(i);
    }
  }
  $module.AllIntegers = function*() {
    yield _dafny.ZERO;
    for (let j = _dafny.ONE;; j = j.plus(1)) {
      yield j;
      yield j.negated();
    }
  }
  $module.IntegerRange = function*(lo, hi) {
    if (lo === null) {
      while (true) {
        hi = hi.minus(1);
        yield hi;
      }
    } else if (hi === null) {
      while (true) {
        yield lo;
        lo = lo.plus(1);
      }
    } else {
      while (lo.isLessThan(hi)) {
        yield lo;
        lo = lo.plus(1);
      }
    }
  }
  $module.SingleValue = function*(v) {
    yield v;
  }
  $module.HaltException = class HaltException extends Error {
    constructor(message) {
      super(message)
    }
  }
  $module.HandleHaltExceptions = function(f) {
    try {
      f()
    } catch (e) {
      if (e instanceof _dafny.HaltException) {
        process.stdout.write("[Program halted] " + e.message + "\n")
        process.exitCode = 1
      } else {
        throw e
      }
    }
  }
  $module.FromMainArguments = function(args) {
    var a = [...args];
    a.splice(0, 2, args[0] + " " + args[1]);
    return a;
  }
  $module.UnicodeFromMainArguments = function(args) {
    return $module.FromMainArguments(args).map(_dafny.Seq.UnicodeFromString);
  }
  return $module;

  // What follows are routines private to the Dafny runtime
  function buildArray(initValue, ...dims) {
    if (dims.length === 0) {
      return initValue;
    } else {
      let a = Array(dims[0].toNumber());
      let b = Array.from(a, (x) => buildArray(initValue, ...dims.slice(1)));
      return b;
    }
  }
  function arrayElementsToString(a) {
    // like `a.join(", ")`, but calling _dafny.toString(x) on every element x instead of x.toString()
    let s = "";
    let sep = "";
    for (let x of a) {
      s += sep + _dafny.toString(x);
      sep = ", ";
    }
    return s;
  }
  function BigNumberGcd(a, b){  // gcd of two non-negative BigNumber's
    while (true) {
      if (a.isZero()) {
        return b;
      } else if (b.isZero()) {
        return a;
      }
      if (a.isLessThan(b)) {
        b = b.modulo(a);
      } else {
        a = a.modulo(b);
      }
    }
  }
})();
// Dafny program systemModulePopulator.dfy compiled into JavaScript
let _System = (function() {
  let $module = {};

  $module.nat = class nat {
    constructor () {
    }
    static get Default() {
      return _dafny.ZERO;
    }
    static _Is(__source) {
      let _0_x = (__source);
      return (_dafny.ZERO).isLessThanOrEqualTo(_0_x);
    }
  };

  return $module;
})(); // end of module _System
let _module = (function() {
  let $module = {};

  $module.__default = class __default {
    constructor () {
      this._tname = "_module._default";
    }
    _parentTraits() {
      return [];
    }
    static abs(x) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return (new BigNumber(-1)).multipliedBy(x);
      } else {
        return x;
      }
    };
    static safeIndex(x, length) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return _dafny.ZERO;
      } else if ((length).isLessThanOrEqualTo(x)) {
        return (x).mod(length);
      } else {
        return x;
      }
    };
    static safeDivisionInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return _dafny.EuclideanDivision(x1, x2);
      }
    };
    static safeModuloInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return (x1).mod(x2);
      }
    };
    static fm0(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(-229)),!(true))).contains((_dafny.MultiSet.fromElements(new BigNumber(49))).Union(_dafny.MultiSet.fromElements(new BigNumber(-112))));
    };
    static fm1(p0, p1, p2, globalState) {
      return (_module.__default.safeDivisionInt(new BigNumber((_dafny.Set.fromElements(true)).length), (_dafny.ZERO).minus(new BigNumber(-576)))).minus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(604),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(471),true))).length));
    };
    static fm2(p0, p1, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("byj")).length),new BigNumber(-962))).length), new BigNumber(-999)), _dafny.Seq.of(new BigNumber(227))), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(599))).length))));
    };
    static fm3(p0, p1, globalState) {
      return new _dafny.CodePoint('w'.codePointAt(0));
    };
    static fm9(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(false),true));
    };
    static fm10(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jyepokid"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gl"), _dafny.Seq.UnicodeFromString("djdolp")));
    };
    static fm11(p0, globalState) {
      let _source0 = ((true) ? (_module.D0.create_DC0(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(822), new BigNumber(444)), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true)).length))))) : (_module.D0.create_DC0(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(634))))));
      if (_source0.is_DC1) {
        let _0___mcc_h0 = (_source0).cf1;
        let _1_cf1 = _0___mcc_h0;
        if (false) {
          return _module.D3.create_DC11(new _dafny.CodePoint('q'.codePointAt(0)), !(true), !(true));
        } else {
          return _module.D3.create_DC11(new _dafny.CodePoint('m'.codePointAt(0)), !(true), true);
        }
      } else if (_source0.is_DC2) {
        let _2___mcc_h1 = (_source0).cf2;
        let _3___mcc_h2 = (_source0).cf3;
        let _4___mcc_h3 = (_source0).cf4;
        let _5___mcc_h4 = (_source0).cf5;
        let _6___mcc_h5 = (_source0).cf6;
        let _7_cf6 = _6___mcc_h5;
        let _8_cf5 = _5___mcc_h4;
        let _9_cf4 = _4___mcc_h3;
        let _10_cf3 = _3___mcc_h2;
        let _11_cf2 = _2___mcc_h1;
        return _module.D3.create_DC11(new _dafny.CodePoint('a'.codePointAt(0)), _9_cf4, _11_cf2);
      } else if (_source0.is_DC3) {
        let _12___mcc_h6 = (_source0).cf7;
        let _13_cf7 = _12___mcc_h6;
        return _module.D3.create_DC11(new _dafny.CodePoint('f'.codePointAt(0)), false, true);
      } else {
        let _14___mcc_h7 = (_source0).cf0;
        let _15_cf0 = _14___mcc_h7;
        return _module.D3.create_DC11(new _dafny.CodePoint('t'.codePointAt(0)), false, true);
      }
    };
    static fm12(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(312)), function (_16_i0) {
        return new BigNumber((function () {
          let _coll0 = new _dafny.Map();
          for (const _compr_0 of (function () {
            let _coll1 = new _dafny.Map();
            for (const _compr_1 of _dafny.IntegerRange(new BigNumber(-198), new BigNumber(-371))) {
              let _17_v1 = _compr_1;
              if (((new BigNumber(-198)).isLessThanOrEqualTo(_17_v1)) && ((_17_v1).isLessThan(new BigNumber(-371)))) {
                _coll1.push([(_17_v1).plus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-110))).cardinality())),!(false)]);
              }
            }
            return _coll1;
          }()).Keys.Elements) {
            let _18_v0 = _compr_0;
            if ((function () {
              let _coll2 = new _dafny.Map();
              for (const _compr_2 of _dafny.IntegerRange(new BigNumber(-198), new BigNumber(-371))) {
                let _17_v1 = _compr_2;
                if (((new BigNumber(-198)).isLessThanOrEqualTo(_17_v1)) && ((_17_v1).isLessThan(new BigNumber(-371)))) {
                  _coll2.push([(_17_v1).plus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-110))).cardinality())),!(false)]);
                }
              }
              return _coll2;
            }()).contains(_18_v0)) {
              _coll0.push([(_18_v0).minus(new BigNumber(873)),true]);
            }
          }
          return _coll0;
        }()).length);
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-962)), function (_19_i1) {
        return new BigNumber((_dafny.Seq.of(new BigNumber(987))).length);
      })), _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-915))));
    };
    static fm13(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(302), new BigNumber(656))) {
          let _20_v0 = _compr_3;
          if (((new BigNumber(302)).isLessThanOrEqualTo(_20_v0)) && ((_20_v0).isLessThan(new BigNumber(656)))) {
            _coll3.push([(_20_v0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("bkdjgk")).length)),new BigNumber((_dafny.MultiSet.fromElements((_module.D2.create_DC6(false)).dtor_cf10)).cardinality())]);
          }
        }
        return _coll3;
      }()), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(false),!(false))).length))));
    };
    static fm14(globalState) {
      return _module.D0.create_DC0(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(587)), _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-336)), function (_21_i0) {
  return new BigNumber(921);
}))), _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(531), new BigNumber(263)), _dafny.MultiSet.fromElements(new BigNumber(245), new BigNumber(950)))));
    };
    static fm15(globalState) {
      return _dafny.Set.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-122)))));
    };
    static fm16(p0, globalState) {
      return _dafny.Seq.of((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(178)), function (_22_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      })).length))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray((_module.D7.create_DC19(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(151))).length),new _dafny.CodePoint('v'.codePointAt(0)))).length),true)).length)))).dtor_cf34)).cardinality())))), (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("vqbthl")).length))).Union((_module.D4.create_DC14(_dafny.MultiSet.fromElements(new BigNumber(119), new BigNumber(-456)), new BigNumber((_dafny.Seq.UnicodeFromString("ty")).length), _module.D2.create_DC5(_dafny.Seq.UnicodeFromString("ehtttxx")))).dtor_cf28), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("lotvixcay")).length)), _dafny.MultiSet.fromElements(new BigNumber(508), new BigNumber(-985)));
    };
    static fm17(globalState) {
      let _source1 = _module.D7.create_DC21(true, true);
      if (_source1.is_DC20) {
        let _23___mcc_h0 = (_source1).cf35;
        let _24___mcc_h1 = (_source1).cf36;
        let _25___mcc_h2 = (_source1).cf37;
        let _26___mcc_h3 = (_source1).cf38;
        let _27___mcc_h4 = (_source1).cf39;
        let _28_cf39 = _27___mcc_h4;
        let _29_cf38 = _26___mcc_h3;
        let _30_cf37 = _25___mcc_h2;
        let _31_cf36 = _24___mcc_h1;
        let _32_cf35 = _23___mcc_h0;
        return _module.D2.create_DC6(_32_cf35);
      } else if (_source1.is_DC21) {
        let _33___mcc_h5 = (_source1).cf40;
        let _34___mcc_h6 = (_source1).cf41;
        let _35_cf41 = _34___mcc_h6;
        let _36_cf40 = _33___mcc_h5;
        if (_35_cf41) {
          return _module.D2.create_DC6(false);
        } else {
          return _module.D2.create_DC6(_36_cf40);
        }
      } else {
        let _37___mcc_h7 = (_source1).cf34;
        let _38_cf34 = _37___mcc_h7;
        return _module.D2.create_DC6(false);
      }
    };
    static fm18(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(897),false)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),new BigNumber(-664)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),new BigNumber(441)));
    };
    static fm19(globalState) {
      return _dafny.MultiSet.fromElements(_module.D0.create_DC0(_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(426)), function (_39_i0) {
  return new BigNumber(500);
})))), _module.D0.create_DC0(_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(410), new BigNumber(803))))), ((false) ? (_module.D0.create_DC0(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true)).length))))) : (_module.D0.create_DC0(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(723))).length)))))));
    };
    static fm20(p0, globalState) {
      return (((true) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.MultiSet.fromElements(false))).length),new BigNumber(-497))) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(130),new BigNumber(-588))))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(811)), function (_40_i0) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      })).length),(_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(false,true))).length), new BigNumber((_dafny.Seq.of(true)).length))).cardinality())))));
    };
    static fm21(p0, p1, globalState) {
      return _dafny.Set.fromElements(false, true, false);
    };
    static m1(p0, p1, globalState) {
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _41_v0;
      _41_v0 = _dafny.Map.Empty.slice().updateUnsafe((p1).isLessThan((_dafny.ZERO).minus(p1)),p0);
      _41_v0 = (_41_v0).update(p0, p0);
      let _42_v1;
      _42_v1 = _dafny.Seq.of(_dafny.Seq.of(p0));
      let _43_v2;
      _43_v2 = _module.D0.create_DC3((_42_v1)[_module.__default.safeIndex(new BigNumber(472), new BigNumber((_42_v1).length))]);
      let _rhs0 = ((p0) ? (_43_v2) : (((p0) ? (_43_v2) : (_43_v2))));
      let _rhs1 = (p1).multipliedBy(p1);
      let _lhs0 = globalState;
      _43_v2 = _rhs0;
      _lhs0.f2 = _rhs1;
      let _44_v3;
      let _nw0 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
      _44_v3 = _nw0;
      let _45_v4;
      _45_v4 = _dafny.Set.fromElements(_44_v3);
      let _46_v5;
      let _nw1 = new _module.C0();
      _nw1.__ctor();
      _46_v5 = _nw1;
      let _47_v6;
      _47_v6 = _dafny.Seq.of(_46_v5);
      let _48_v7;
      _48_v7 = new _dafny.CodePoint('o'.codePointAt(0));
      let _49_v8;
      _49_v8 = _dafny.MultiSet.fromElements(_48_v7, _48_v7, _48_v7);
      let _50_v9;
      _50_v9 = _dafny.Map.Empty.slice().updateUnsafe(_49_v8,p1);
      let _51_v10;
      _51_v10 = _module.D0.create_DC2(p0, _45_v4, !(!(_dafny.Seq.IsProperPrefixOf(_47_v6, _47_v6))), p0, new BigNumber((_50_v9).length));
      let _source2 = _51_v10;
      if (_source2.is_DC1) {
        let _52___mcc_h0 = (_source2).cf1;
        let _53_cf1 = _52___mcc_h0;
        let _54_v11;
        let _nw2 = Array((new BigNumber(15)).toNumber()).fill([]);
        _54_v11 = _nw2;
        let _55_v12;
        _55_v12 = _module.D6.create_DC16(_44_v3);
        let _index0 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_54_v11).length));
        (_54_v11)[_index0] = (_55_v12).dtor_cf32;
        let _56_v13;
        _56_v13 = _dafny.Seq.of(_44_v3);
        let _index1 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_54_v11).length));
        (_54_v11)[_index1] = ((p0) ? (_44_v3) : ((_56_v13)[_module.__default.safeIndex(new BigNumber(686), new BigNumber((_56_v13).length))]));
        let _57_v14;
        let _init0 = ((_58_p0) => function (_59_i0) {
          return !(_58_p0);
        })(p0);
        let _nw3 = Array((new BigNumber(25)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw3.length); _i0_0++) {
          _nw3[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _57_v14 = _nw3;
        let _index2 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_57_v14).length));
        (_57_v14)[_index2] = p0;
        let _index3 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_57_v14).length));
        (_57_v14)[_index3] = p0;
        if ((_57_v14)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_57_v14).length))]) {
          let _60_v15;
          let _nw4 = new _module.C0();
          _nw4.__ctor();
          _60_v15 = _nw4;
          (globalState).f8 = _53_cf1;
          let _61_v16;
          _61_v16 = _dafny.Set.fromElements((_57_v14)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_57_v14).length))]);
          let _62_v17;
          let _nw5 = new _module.C1();
          _nw5.__ctor(_61_v16, _53_cf1);
          _62_v17 = _nw5;
          let _63_v18;
          let _nw6 = Array((new BigNumber(15)).toNumber());
          _nw6[(_dafny.ZERO).toNumber()] = _62_v17;
          _nw6[(_dafny.ONE).toNumber()] = _62_v17;
          _nw6[(new BigNumber(2)).toNumber()] = _62_v17;
          _nw6[(new BigNumber(3)).toNumber()] = _62_v17;
          _nw6[(new BigNumber(4)).toNumber()] = _62_v17;
          _nw6[(new BigNumber(5)).toNumber()] = _62_v17;
          _nw6[(new BigNumber(6)).toNumber()] = _62_v17;
          _nw6[(new BigNumber(7)).toNumber()] = _62_v17;
          _nw6[(new BigNumber(8)).toNumber()] = _62_v17;
          _nw6[(new BigNumber(9)).toNumber()] = _62_v17;
          _nw6[(new BigNumber(10)).toNumber()] = _62_v17;
          _nw6[(new BigNumber(11)).toNumber()] = _62_v17;
          _nw6[(new BigNumber(12)).toNumber()] = _62_v17;
          _nw6[(new BigNumber(13)).toNumber()] = _62_v17;
          _nw6[(new BigNumber(14)).toNumber()] = _62_v17;
          _63_v18 = _nw6;
          let _64_v19;
          let _nw7 = Array((new BigNumber(8)).toNumber());
          _nw7[(_dafny.ZERO).toNumber()] = _63_v18;
          _nw7[(_dafny.ONE).toNumber()] = _63_v18;
          _nw7[(new BigNumber(2)).toNumber()] = _63_v18;
          _nw7[(new BigNumber(3)).toNumber()] = _63_v18;
          _nw7[(new BigNumber(4)).toNumber()] = _63_v18;
          _nw7[(new BigNumber(5)).toNumber()] = _63_v18;
          _nw7[(new BigNumber(6)).toNumber()] = _63_v18;
          _nw7[(new BigNumber(7)).toNumber()] = _63_v18;
          _64_v19 = _nw7;
          let _index4 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_64_v19).length));
          (_64_v19)[_index4] = _63_v18;
          let _65_v20;
          _65_v20 = _dafny.Seq.of(p0);
          let _66_v21;
          _66_v21 = _dafny.Set.fromElements(_65_v20);
          let _67_v22;
          _67_v22 = _dafny.Map.Empty.slice().updateUnsafe(p0,_63_v18);
          let _68_v23;
          _68_v23 = _dafny.Seq.of(_63_v18, (((_67_v22).contains((_57_v14)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_57_v14).length))])) ? ((_67_v22).get((_57_v14)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_57_v14).length))])) : (_63_v18)));
          let _69_v24;
          _69_v24 = _dafny.MultiSet.fromElements(_62_v17.f27);
          let _70_v25;
          _70_v25 = _dafny.Seq.of((_68_v23)[_module.__default.safeIndex((_module.D3.create_DC10(p0, new BigNumber((_69_v24).cardinality()), p0)).dtor_cf17, new BigNumber((_68_v23).length))], _63_v18, _63_v18);
          let _71_v26;
          let _nw8 = new _module.C0();
          _nw8.__ctor();
          _71_v26 = _nw8;
          let _72_v27;
          _72_v27 = _dafny.Seq.of(_71_v26);
          let _73_v28;
          _73_v28 = _dafny.MultiSet.fromElements((_72_v27)[_module.__default.safeIndex(p1, new BigNumber((_72_v27).length))], _71_v26);
          let _index5 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_57_v14).length));
          let _index6 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_64_v19).length));
          let _rhs2 = (_66_v21).IsProperSubsetOf(_dafny.Set.fromElements(_65_v20));
          let _rhs3 = (_57_v14)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_57_v14).length))];
          let _rhs4 = (_70_v25)[_module.__default.safeIndex(new BigNumber((_73_v28).cardinality()), new BigNumber((_70_v25).length))];
          let _lhs1 = globalState;
          let _lhs2 = _57_v14;
          let _lhs3 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_57_v14).length));
          let _lhs4 = _64_v19;
          let _lhs5 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_64_v19).length));
          _lhs1.f23 = _rhs2;
          _lhs2[_lhs3] = _rhs3;
          _lhs4[_lhs5] = _rhs4;
          (globalState).f2 = p1;
          (globalState).f23 = !((_57_v14)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_57_v14).length))]);
        } else {
          let _74_v29;
          _74_v29 = _dafny.Seq.of((_57_v14)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_57_v14).length))], (_57_v14)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_57_v14).length))]);
          let _75_v30;
          _75_v30 = _dafny.Seq.UnicodeFromString("n");
          (globalState).f19 = _module.__default.fm1((_74_v29)[_module.__default.safeIndex(new BigNumber((_75_v30).length), new BigNumber((_74_v29).length))], _48_v7, new BigNumber((_75_v30).length), globalState);
          let _76_v31;
          let _nw9 = new _module.C0();
          _nw9.__ctor();
          _76_v31 = _nw9;
          (globalState).f23 = ((p1).multipliedBy(new BigNumber(572))).isLessThanOrEqualTo(new BigNumber(474));
          let _77_v32;
          _77_v32 = _dafny.MultiSet.fromElements(p1);
          let _78_v33;
          _78_v33 = _dafny.Seq.of(_77_v32, _77_v32);
          let _79_v34;
          _79_v34 = _module.D0.create_DC0(_78_v33);
          let _80_v35;
          _80_v35 = _dafny.MultiSet.fromElements(_79_v34, _79_v34);
          let _81_v36;
          _81_v36 = _80_v35;
          let _82_v37;
          _82_v37 = _dafny.Set.fromElements((_57_v14)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_57_v14).length))]);
          let _83_v38;
          let _nw10 = new _module.C1();
          _nw10.__ctor(_82_v37, p1);
          _83_v38 = _nw10;
          let _84_v39;
          _84_v39 = _dafny.Map.Empty.slice().updateUnsafe((_54_v11)[_module.__default.safeIndex(new BigNumber(549), new BigNumber((_54_v11).length))],new BigNumber((_74_v29).length));
          let _85_v40;
          _85_v40 = _dafny.Map.Empty.slice().updateUnsafe(_83_v38,(((_84_v39).contains((_54_v11)[_module.__default.safeIndex(new BigNumber(549), new BigNumber((_54_v11).length))])) ? ((_84_v39).get((_54_v11)[_module.__default.safeIndex(new BigNumber(549), new BigNumber((_54_v11).length))])) : (_53_cf1)));
          let _86_v41;
          _86_v41 = _dafny.Seq.of(_85_v40);
          let _index7 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_57_v14).length));
          let _rhs5 = (_80_v35).Intersect(_80_v35);
          let _rhs6 = (_57_v14)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_57_v14).length))];
          let _rhs7 = new BigNumber((_86_v41).length);
          let _rhs8 = _48_v7;
          let _lhs6 = _57_v14;
          let _lhs7 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_57_v14).length));
          let _lhs8 = globalState;
          _81_v36 = _rhs5;
          _lhs6[_lhs7] = _rhs6;
          _lhs8.f10 = _rhs7;
          _48_v7 = _rhs8;
          let _87_v42;
          _87_v42 = _dafny.Map.Empty.slice().updateUnsafe(p1,(((_41_v0).contains(p0)) ? ((_41_v0).get(p0)) : (p0)));
          let _88_v43;
          _88_v43 = _dafny.Seq.of(_module.__default.fm20((((_87_v42).contains(p1)) ? ((_87_v42).get(p1)) : ((_57_v14)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_57_v14).length))])), globalState));
          let _89_v44;
          _89_v44 = _dafny.Map.Empty.slice().updateUnsafe(_44_v3,p0);
          let _90_v45;
          _90_v45 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_88_v43)[_module.__default.safeIndex(new BigNumber((_89_v44).length), new BigNumber((_88_v43).length))]).length),(_54_v11)[_module.__default.safeIndex(new BigNumber(549), new BigNumber((_54_v11).length))]);
          (globalState).f11 = (((_90_v45).contains(_module.__default.safeDivisionInt((_76_v31).fm8(_82_v37, globalState), p1))) ? ((_90_v45).get(_module.__default.safeDivisionInt((_76_v31).fm8(_82_v37, globalState), p1))) : (_44_v3));
        }
        let _91_v46;
        _91_v46 = _dafny.MultiSet.fromElements((_54_v11)[_module.__default.safeIndex(new BigNumber(549), new BigNumber((_54_v11).length))], _44_v3);
        let _index8 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_57_v14).length));
        (_57_v14)[_index8] = (_91_v46).IsSubsetOf(_91_v46);
      } else if (_source2.is_DC2) {
        let _92___mcc_h1 = (_source2).cf2;
        let _93___mcc_h2 = (_source2).cf3;
        let _94___mcc_h3 = (_source2).cf4;
        let _95___mcc_h4 = (_source2).cf5;
        let _96___mcc_h5 = (_source2).cf6;
        let _97_cf6 = _96___mcc_h5;
        let _98_cf5 = _95___mcc_h4;
        let _99_cf4 = _94___mcc_h3;
        let _100_cf3 = _93___mcc_h2;
        let _101_cf2 = _92___mcc_h1;
        r1 = _99_cf4;
        let _102_v47;
        let _nw11 = new _module.C1();
        _nw11.__ctor(_module.__default.fm21(_98_cf5, true, globalState), _97_cf6);
        _102_v47 = _nw11;
        let _103_v48;
        _103_v48 = _dafny.Map.Empty.slice().updateUnsafe(_48_v7,_98_cf5);
        let _104_v49;
        _104_v49 = _dafny.Seq.UnicodeFromString("lpfwbpp");
        let _105_v50;
        _105_v50 = _dafny.Map.Empty.slice().updateUnsafe(_103_v48,_104_v49);
        _102_v47 = ((!_dafny.Seq.contains((((_105_v50).contains(_103_v48)) ? ((_105_v50).get(_103_v48)) : (_104_v49)), _48_v7)) ? (_102_v47) : (_102_v47));
        _49_v8 = _dafny.MultiSet.FromArray(_104_v49);
        if (!(_101_cf2)) {
          (globalState).f19 = p1;
          let _106_v51;
          _106_v51 = _dafny.Set.fromElements(_99_cf4, !(_98_cf5), _99_cf4, p0);
          let _107_v52;
          _107_v52 = _dafny.Seq.of(_dafny.Set.fromElements(_101_cf2), _106_v51, _106_v51);
          let _108_v53;
          _108_v53 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_107_v52).length));
          let _109_v54;
          _109_v54 = _dafny.Map.Empty.slice().updateUnsafe(_97_cf6,new BigNumber((_108_v53).length));
          let _110_v55;
          _110_v55 = _dafny.Seq.of(new BigNumber(206), (((_109_v54).contains(p1)) ? ((_109_v54).get(p1)) : (p1)));
          _110_v55 = _110_v55;
          _48_v7 = _48_v7;
          let _index9 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_44_v3).length));
          (_44_v3)[_index9] = _97_cf6;
          let _111_v56;
          _111_v56 = _dafny.MultiSet.fromElements(_101_cf2);
          let _index10 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_44_v3).length));
          let _rhs9 = _97_cf6;
          let _rhs10 = (new BigNumber((_dafny.Seq.UnicodeFromString("mjpyri")).length)).multipliedBy(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(816)), function (_112_i1) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          }), _104_v49), _module.__default.safeIndex((_dafny.ZERO).minus((((_111_v56).contains(_98_cf5)) ? ((_111_v56).get(_98_cf5)) : (_97_cf6))), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(816)), function (_113_i1) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          }), _104_v49)).length)), _48_v7)).length));
          let _rhs11 = true;
          let _rhs12 = (_dafny.ZERO).minus((_97_cf6).plus((_dafny.ZERO).minus(_97_cf6)));
          let _lhs9 = _44_v3;
          let _lhs10 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_44_v3).length));
          let _lhs11 = globalState;
          let _lhs12 = globalState;
          _lhs9[_lhs10] = _rhs9;
          _lhs11.f5 = _rhs10;
          r1 = _rhs11;
          _lhs12.f10 = _rhs12;
          let _114_v57;
          let _nw12 = new _module.C0();
          _nw12.__ctor();
          _114_v57 = _nw12;
        } else {
          let _index11 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_44_v3).length));
          (_44_v3)[_index11] = p1;
          let _index12 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_44_v3).length));
          (_44_v3)[_index12] = p1;
          let _115_v58;
          _115_v58 = _dafny.MultiSet.fromElements(_104_v49);
          let _index13 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_44_v3).length));
          (_44_v3)[_index13] = (((_115_v58).contains(_dafny.Seq.update(_104_v49, _module.__default.safeIndex(_97_cf6, new BigNumber((_104_v49).length)), _48_v7))) ? ((_115_v58).get(_dafny.Seq.update(_104_v49, _module.__default.safeIndex(_97_cf6, new BigNumber((_104_v49).length)), _48_v7))) : (p1));
          let _116_v59;
          let _nw13 = Array((new BigNumber(21)).toNumber()).fill(false);
          _116_v59 = _nw13;
          let _117_v60;
          _117_v60 = _dafny.Map.Empty.slice().updateUnsafe(_116_v59,p1);
          let _118_v61;
          _118_v61 = _dafny.Map.Empty.slice().updateUnsafe(_99_cf4,_117_v60);
          (globalState).f23 = (((_41_v0).contains((p1).isLessThan(_97_cf6))) ? ((_41_v0).get((p1).isLessThan(_97_cf6))) : ((_117_v60).equals((((_118_v61).contains(_98_cf5)) ? ((_118_v61).get(_98_cf5)) : (_117_v60)))));
          let _119_v62;
          _119_v62 = _dafny.Set.fromElements(_98_cf5, _101_cf2, !(_98_cf5));
          let _120_v63;
          _120_v63 = _dafny.MultiSet.fromElements(new BigNumber((_104_v49).length));
          let _121_v64;
          _121_v64 = _dafny.Seq.of(_120_v63);
          let _122_v65;
          _122_v65 = _module.D0.create_DC0(_121_v64);
          let _123_v66;
          _123_v66 = _dafny.MultiSet.fromElements(_122_v65);
          let _124_v67;
          _124_v67 = _123_v66;
          let _125_v68;
          _125_v68 = _dafny.Map.Empty.slice().updateUnsafe(((_46_v5).fm8(_119_v62, globalState)).multipliedBy((_dafny.ZERO).minus((_44_v3)[_module.__default.safeIndex(new BigNumber(250), new BigNumber((_44_v3).length))])),_124_v67);
          _125_v68 = _125_v68;
          (globalState).f22 = _99_cf4;
        }
      } else if (_source2.is_DC3) {
        let _126___mcc_h6 = (_source2).cf7;
        let _127_cf7 = _126___mcc_h6;
        let _128_v69;
        _128_v69 = _dafny.MultiSet.fromElements(_44_v3);
        let _129_v70;
        _129_v70 = _dafny.Seq.of(p1);
        let _130_v71;
        _130_v71 = _dafny.Seq.of(_module.__default.safeModuloInt(new BigNumber((_128_v69).cardinality()), (_129_v70)[_module.__default.safeIndex(p1, new BigNumber((_129_v70).length))]), p1);
        _129_v70 = _dafny.Seq.Concat(_dafny.Seq.of(p1), _130_v71);
        let _index14 = _module.__default.safeIndex(new BigNumber(955), new BigNumber((_44_v3).length));
        (_44_v3)[_index14] = p1;
        let _index15 = _module.__default.safeIndex(new BigNumber(955), new BigNumber((_44_v3).length));
        (_44_v3)[_index15] = p1;
        let _131_v72;
        _131_v72 = _dafny.MultiSet.fromElements(_module.__default.fm0(p0, _41_v0, globalState));
        let _132_v73;
        _132_v73 = _module.D2.create_DC8(p1, p0, (((_131_v72).contains(p0)) ? ((_131_v72).get(p0)) : ((_44_v3)[_module.__default.safeIndex(new BigNumber(955), new BigNumber((_44_v3).length))])));
        _132_v73 = _132_v73;
        _46_v5 = _46_v5;
      } else {
        let _133___mcc_h7 = (_source2).cf0;
        let _134_cf0 = _133___mcc_h7;
        (globalState).f25 = p0;
        let _rhs13 = !(p0);
        let _rhs14 = p1;
        let _lhs13 = globalState;
        r1 = _rhs13;
        _lhs13.f2 = _rhs14;
        let _135_v74;
        _135_v74 = _module.D0.create_DC1(p1);
        let _136_v75;
        let _nw14 = new _module.C0();
        _nw14.__ctor();
        _136_v75 = _nw14;
        let _137_v76;
        _137_v76 = _module.D3.create_DC12(p0, _135_v74, p0, p0, _136_v75);
        let _138_v77;
        _138_v77 = _dafny.Map.Empty.slice().updateUnsafe(_137_v76,p1);
        _138_v77 = (_138_v77).update(_137_v76, new BigNumber((_dafny.Seq.UnicodeFromString("k")).length));
        let _139_v78;
        _139_v78 = _dafny.Seq.UnicodeFromString("akpsj");
        _139_v78 = _dafny.Seq.Concat(_139_v78, _139_v78);
      }
      let _hi0 = p1;
      for (let _140_i2 = new BigNumber(((_41_v0).update(p0, p0)).length); _140_i2.isLessThan(_hi0); _140_i2 = _140_i2.plus(_dafny.ONE)) {
        let _141_v79;
        _141_v79 = _module.D3.create_DC10(p0, _140_i2, p0);
        r1 = _module.__default.fm0((_141_v79).dtor_cf16, _41_v0, globalState);
        if (!(new BigNumber(106)).isEqualTo(_140_i2)) {
          let _142_v80;
          _142_v80 = _dafny.MultiSet.fromElements(p0, p0);
          let _143_v81;
          _143_v81 = _dafny.Seq.of(p0);
          let _rhs15 = !(!(_module.__default.fm0(!(p0), _dafny.Map.Empty.slice().updateUnsafe(false,_module.__default.fm0(true, _41_v0, globalState)), globalState)));
          let _rhs16 = _140_i2;
          let _rhs17 = _142_v80;
          let _rhs18 = (_143_v81)[_module.__default.safeIndex(p1, new BigNumber((_143_v81).length))];
          let _lhs14 = globalState;
          let _lhs15 = globalState;
          let _lhs16 = globalState;
          let _lhs17 = globalState;
          _lhs14.f25 = _rhs15;
          _lhs15.f19 = _rhs16;
          _lhs16.f16 = _rhs17;
          _lhs17.f23 = _rhs18;
          let _144_v82;
          _144_v82 = _dafny.Set.fromElements(p0, p0, p0, p0, p0);
          let _145_v84;
          _145_v84 = _dafny.Set.fromElements(_140_i2, p1, _140_i2);
          let _146_v85;
          _146_v85 = _dafny.Map.Empty.slice().updateUnsafe(_144_v82,(function () {
            let _coll4 = new _dafny.Set();
            for (const _compr_4 of _dafny.IntegerRange(new BigNumber(20), new BigNumber(792))) {
              let _147_v83 = _compr_4;
              if (((new BigNumber(20)).isLessThanOrEqualTo(_147_v83)) && ((_147_v83).isLessThan(new BigNumber(792)))) {
                _coll4.add((_147_v83).plus(p1));
              }
            }
            return _coll4;
          }()).Union(_145_v84));
          let _148_v86;
          let _nw15 = new _module.C1();
          _nw15.__ctor(_dafny.Set.fromElements(p0), _140_i2);
          _148_v86 = _nw15;
          let _149_v87;
          _149_v87 = _dafny.MultiSet.fromElements(_148_v86, _148_v86);
          let _rhs19 = (_dafny.MultiSet.fromElements(_148_v86, _148_v86)).IsSubsetOf((_149_v87).Intersect(_149_v87));
          let _rhs20 = _dafny.Map.Empty.slice().updateUnsafe(_144_v82,(_145_v84).Difference(_145_v84));
          let _rhs21 = _module.__default.fm0(_module.__default.fm0(false, _41_v0, globalState), _41_v0, globalState);
          let _rhs22 = p1;
          let _lhs18 = globalState;
          let _lhs19 = globalState;
          let _lhs20 = globalState;
          _lhs18.f22 = _rhs19;
          _146_v85 = _rhs20;
          _lhs19.f22 = _rhs21;
          _lhs20.f15 = _rhs22;
          let _150_v88;
          _150_v88 = _dafny.Map.Empty.slice().updateUnsafe(_148_v86,((p0) ? (_148_v86) : (_148_v86)));
          _150_v88 = (_150_v88).update(_148_v86, _148_v86);
          let _151_v89;
          _151_v89 = _module.D0.create_DC1(p1);
          let _152_v90;
          _152_v90 = _dafny.Seq.of(new BigNumber(800), _140_i2);
          let _153_v91;
          _153_v91 = _dafny.Map.Empty.slice().updateUnsafe((p0) && (!(_module.__default.fm0(true, _dafny.Map.Empty.slice().updateUnsafe(p0,p0), globalState))),new BigNumber((_dafny.Seq.Concat(_module.__default.fm12(new BigNumber((_145_v84).length), _151_v89, globalState), _152_v90)).length));
          _153_v91 = (_153_v91).update(((!(p0)) ? (true) : (_module.__default.fm0(p0, _41_v0, globalState))), _module.__default.safeDivisionInt(new BigNumber((_module.__default.fm21(p0, p0, globalState)).length), p1));
          let _154_v92;
          _154_v92 = _dafny.Seq.UnicodeFromString("ldritmxh");
          _154_v92 = _dafny.Seq.UnicodeFromString("oocr");
        } else {
          (globalState).f5 = ((p0) ? (_140_i2) : (_140_i2));
          let _155_v93;
          _155_v93 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
          let _156_v95;
          _156_v95 = _dafny.Seq.UnicodeFromString("b");
          let _157_v96;
          _157_v96 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_156_v95).length),p1);
          let _rhs23 = (function () {
            let _coll5 = new _dafny.Map();
            for (const _compr_5 of (_157_v96).Keys.Elements) {
              let _158_v94 = _compr_5;
              if ((_157_v96).contains(_158_v94)) {
                _coll5.push([(_158_v94).minus(p1),p0]);
              }
            }
            return _coll5;
          }()).Merge(_155_v93);
          let _rhs24 = _48_v7;
          _155_v93 = _rhs23;
          _48_v7 = _rhs24;
          (globalState).f22 = p0;
          let _index16 = _module.__default.safeIndex(new BigNumber(45), new BigNumber((_44_v3).length));
          (_44_v3)[_index16] = (p1).multipliedBy(new BigNumber(695));
          let _index17 = _module.__default.safeIndex(new BigNumber(45), new BigNumber((_44_v3).length));
          (_44_v3)[_index17] = _140_i2;
          let _159_v97;
          _159_v97 = _dafny.Set.fromElements(p0);
          let _160_v98;
          let _nw16 = new _module.C1();
          _nw16.__ctor(_159_v97, p1);
          _160_v98 = _nw16;
          let _161_v99;
          _161_v99 = _dafny.Map.Empty.slice().updateUnsafe(_160_v98,new BigNumber((_156_v95).length));
          let _162_v100;
          _162_v100 = _dafny.Map.Empty.slice().updateUnsafe(_160_v98,false);
          let _163_v101;
          _163_v101 = _dafny.Seq.of(_140_i2, new BigNumber((_162_v100).length), (_dafny.ZERO).minus(p1));
          let _164_v102;
          _164_v102 = _dafny.Seq.of(false, p0);
          let _165_v103;
          _165_v103 = _module.D0.create_DC1(new BigNumber((_dafny.MultiSet.FromArray(_164_v102)).cardinality()));
          let _166_v104;
          _166_v104 = _dafny.Map.Empty.slice().updateUnsafe(_156_v95,p0);
          let _pat_let_tv0 = p0;
          let _pat_let_tv1 = _48_v7;
          let _pat_let_tv2 = _44_v3;
          let _pat_let_tv3 = _44_v3;
          let _pat_let_tv4 = globalState;
          let _167_v105;
          let _nw17 = Array((new BigNumber(20)).toNumber());
          _nw17[(_dafny.ZERO).toNumber()] = (_44_v3)[_module.__default.safeIndex(new BigNumber(45), new BigNumber((_44_v3).length))];
          _nw17[(_dafny.ONE).toNumber()] = _140_i2;
          _nw17[(new BigNumber(2)).toNumber()] = ((true) ? (new BigNumber(952)) : (_140_i2));
          _nw17[(new BigNumber(3)).toNumber()] = (_44_v3)[_module.__default.safeIndex(new BigNumber(45), new BigNumber((_44_v3).length))];
          _nw17[(new BigNumber(4)).toNumber()] = (new BigNumber((_161_v99).length)).plus((_dafny.ZERO).minus((_44_v3)[_module.__default.safeIndex(new BigNumber(45), new BigNumber((_44_v3).length))]));
          _nw17[(new BigNumber(5)).toNumber()] = p1;
          _nw17[(new BigNumber(6)).toNumber()] = (_163_v101)[_module.__default.safeIndex(p1, new BigNumber((_163_v101).length))];
          _nw17[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(_160_v98.f27);
          _nw17[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(_160_v98.f27);
          _nw17[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(_160_v98.f27);
          _nw17[(new BigNumber(10)).toNumber()] = (_44_v3)[_module.__default.safeIndex(new BigNumber(45), new BigNumber((_44_v3).length))];
          _nw17[(new BigNumber(11)).toNumber()] = ((!(_module.__default.fm0(p0, _dafny.Map.Empty.slice().updateUnsafe(p0,p0), globalState))) ? ((_44_v3)[_module.__default.safeIndex(new BigNumber(45), new BigNumber((_44_v3).length))]) : (new BigNumber((_dafny.MultiSet.fromElements((_164_v102)[_module.__default.safeIndex((_dafny.ZERO).minus(_140_i2), new BigNumber((_164_v102).length))], p0, p0, p0)).cardinality())));
          _nw17[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("xrsul")).length);
          _nw17[(new BigNumber(13)).toNumber()] = (_44_v3)[_module.__default.safeIndex(new BigNumber(45), new BigNumber((_44_v3).length))];
          _nw17[(new BigNumber(14)).toNumber()] = new BigNumber((_156_v95).length);
          _nw17[(new BigNumber(15)).toNumber()] = ((p0) ? (_160_v98.f27) : (_module.__default.fm1(p0, _48_v7, _160_v98.f27, globalState)));
          _nw17[(new BigNumber(16)).toNumber()] = p1;
          _nw17[(new BigNumber(17)).toNumber()] = ((function (_pat_let0_0) {
            return function (_168_dt__update__tmp_h2) {
              return function (_pat_let1_0) {
                return function (_169_dt__update_hcf1_h0) {
                  return _module.D0.create_DC1(_169_dt__update_hcf1_h0);
                }(_pat_let1_0);
              }(_module.__default.fm1(_pat_let_tv0, _pat_let_tv1, (_pat_let_tv3)[_module.__default.safeIndex(new BigNumber(45), new BigNumber((_pat_let_tv2).length))], _pat_let_tv4));
            }(_pat_let0_0);
          }(_165_v103)).dtor_cf1).plus(new BigNumber(((_dafny.MultiSet.fromElements(p0)).update((((_166_v104).contains(_156_v95)) ? ((_166_v104).get(_156_v95)) : (p0)), _module.__default.abs(_140_i2))).cardinality()));
          _nw17[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(_140_i2);
          _nw17[(new BigNumber(19)).toNumber()] = (_44_v3)[_module.__default.safeIndex(new BigNumber(45), new BigNumber((_44_v3).length))];
          _167_v105 = _nw17;
          (globalState).f11 = _167_v105;
        }
        r0 = p1;
        let _170_v106;
        let _nw18 = new _module.C0();
        _nw18.__ctor();
        _170_v106 = _nw18;
      }
      let _171_v107;
      let _nw19 = new _module.C1();
      _nw19.__ctor(_module.__default.fm21(p0, !(p0), globalState), p1);
      _171_v107 = _nw19;
      let _172_v108;
      _172_v108 = _dafny.Set.fromElements(_171_v107, _171_v107);
      _172_v108 = (_dafny.Set.fromElements(_171_v107)).Intersect(_172_v108);
      let _173_v109;
      let _174_v110;
      let _out0;
      let _out1;
      let _outcollector0 = (_171_v107).m0(p0, globalState);
      _out0 = _outcollector0[0];
      _out1 = _outcollector0[1];
      _173_v109 = _out0;
      _174_v110 = _out1;
      let _175_v111;
      _175_v111 = _dafny.MultiSet.fromElements(true);
      r0 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(((p0) ? (_module.__default.fm1(p0, _48_v7, _173_v109, globalState)) : ((((_175_v111).contains(p0)) ? ((_175_v111).get(p0)) : (p1))))), _173_v109);
      r1 = p0;
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _176_v0;
      _176_v0 = true;
      let _177_v1;
      _177_v1 = new BigNumber(17);
      let _178_v2;
      _178_v2 = _dafny.Seq.of(_177_v1);
      let _179_v3;
      _179_v3 = _dafny.Map.Empty.slice().updateUnsafe(_176_v0,_178_v2);
      let _180_v4;
      let _init1 = ((_181_v1) => function (_182_i1) {
        return (_182_i1).plus(_181_v1);
      })(_177_v1);
      let _nw20 = Array((new BigNumber(8)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw20.length); _i0_1++) {
        _nw20[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _180_v4 = _nw20;
      let _183_v5;
      _183_v5 = _dafny.Seq.UnicodeFromString("cb");
      let _184_v6;
      _184_v6 = _dafny.MultiSet.fromElements(_183_v5);
      let _185_v8;
      _185_v8 = _dafny.MultiSet.fromElements(_176_v0, _176_v0, _176_v0);
      let _186_v9;
      let _nw21 = Array((new BigNumber(11)).toNumber());
      _nw21[(_dafny.ZERO).toNumber()] = _176_v0;
      _nw21[(_dafny.ONE).toNumber()] = false;
      _nw21[(new BigNumber(2)).toNumber()] = _176_v0;
      _nw21[(new BigNumber(3)).toNumber()] = _176_v0;
      _nw21[(new BigNumber(4)).toNumber()] = _176_v0;
      _nw21[(new BigNumber(5)).toNumber()] = _176_v0;
      _nw21[(new BigNumber(6)).toNumber()] = _176_v0;
      _nw21[(new BigNumber(7)).toNumber()] = _176_v0;
      _nw21[(new BigNumber(8)).toNumber()] = _176_v0;
      _nw21[(new BigNumber(9)).toNumber()] = _176_v0;
      _nw21[(new BigNumber(10)).toNumber()] = _176_v0;
      _186_v9 = _nw21;
      let _187_v10;
      _187_v10 = new _dafny.CodePoint('u'.codePointAt(0));
      let _188_v11;
      _188_v11 = _dafny.MultiSet.fromElements(_187_v10, _187_v10, new _dafny.CodePoint('f'.codePointAt(0)), _187_v10);
      let _189_v12;
      _189_v12 = _dafny.Map.Empty.slice().updateUnsafe(_177_v1,_dafny.MultiSet.fromElements(_187_v10));
      let _190_globalState;
      let _nw22 = new _module.GlobalState();
      _nw22.__ctor(_179_v3, new BigNumber(201), new BigNumber(-912), new BigNumber(-807), false, new BigNumber(990), _dafny.Seq.Create(_module.__default.abs(new BigNumber(144)), function (_191_i0) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }), new BigNumber(150), new BigNumber(944), true, new BigNumber(334), _180_v4, function () {
        let _coll6 = new _dafny.Set();
        for (const _compr_6 of (_184_v6).Elements) {
          let _192_v7 = _compr_6;
          if ((_184_v6).contains(_192_v7)) {
            _coll6.add(_192_v7);
          }
        }
        return _coll6;
      }(), _dafny.Seq.Concat(_183_v5, _183_v5), true, new BigNumber(850), _185_v8, _dafny.Seq.Create(_module.__default.abs(new BigNumber(20)), function (_193_i2) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      }), _186_v9, new BigNumber(-870), _183_v5, (_dafny.Map.Empty.slice().updateUnsafe(_177_v1,_188_v11)).Merge(_189_v12), false, false, _183_v5, true);
      _190_globalState = _nw22;
      let _194_v13;
      _194_v13 = _dafny.Seq.of(_180_v4, _180_v4, _180_v4, _180_v4);
      _194_v13 = _194_v13;
      let _195_v14;
      _195_v14 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_177_v1),_176_v0);
      let _196_v15;
      _196_v15 = _dafny.Map.Empty.slice().updateUnsafe(_176_v0,_176_v0);
      if (!(_module.__default.fm0((((_195_v14).contains(_177_v1)) ? ((_195_v14).get(_177_v1)) : (_176_v0)), (_196_v15).Merge(_196_v15), _190_globalState))) {
        (_190_globalState).f10 = (_dafny.ZERO).minus(_177_v1);
        (_190_globalState).f22 = !(_176_v0) || (_176_v0);
        let _index18 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_180_v4).length));
        (_180_v4)[_index18] = new BigNumber(945);
        let _index19 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_180_v4).length));
        (_180_v4)[_index19] = _module.__default.safeModuloInt(new BigNumber(669), _177_v1);
        if (_176_v0) {
          (_190_globalState).f2 = ((_180_v4)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_180_v4).length))]).multipliedBy((_180_v4)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_180_v4).length))]);
          let _197_v16;
          _197_v16 = _dafny.Seq.of(_176_v0);
          let _198_v17;
          _198_v17 = _dafny.MultiSet.fromElements(new BigNumber((_178_v2).length), new BigNumber((_197_v16).length), _177_v1);
          let _199_v18;
          _199_v18 = _dafny.MultiSet.fromElements(_198_v17);
          let _200_v19;
          _200_v19 = _dafny.MultiSet.fromElements((((_199_v18).contains(_198_v17)) ? ((_199_v18).get(_198_v17)) : (_177_v1)));
          let _201_v20;
          _201_v20 = _dafny.Map.Empty.slice().updateUnsafe(_176_v0,(_180_v4)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_180_v4).length))]);
          _176_v0 = ((((_185_v8).contains(_176_v0)) ? ((_185_v8).get(_176_v0)) : ((_180_v4)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_180_v4).length))]))).isLessThanOrEqualTo(_module.__default.safeDivisionInt((((_198_v17).contains(_module.__default.fm1(_176_v0, _187_v10, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_178_v2)).cardinality()),(_180_v4)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_180_v4).length))])).length), _190_globalState))) ? ((_198_v17).get(_module.__default.fm1(_176_v0, _187_v10, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_178_v2)).cardinality()),(_180_v4)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_180_v4).length))])).length), _190_globalState))) : ((_180_v4)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_180_v4).length))])), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_201_v20).length),false)).length))));
          (_190_globalState).f2 = _module.__default.safeModuloInt(_177_v1, (_180_v4)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_180_v4).length))]);
          let _202_v21;
          _202_v21 = _dafny.Seq.of((_200_v19).Union(_198_v17), (_dafny.MultiSet.FromArray(_178_v2)).Intersect(_module.__default.fm2(_module.__default.fm1(_176_v0, _module.__default.fm3(_177_v1, _176_v0, _190_globalState), _177_v1, _190_globalState), _176_v0, _190_globalState)));
          let _203_v22;
          _203_v22 = _module.D0.create_DC0(_202_v21);
          _202_v21 = (_203_v22).dtor_cf0;
          _183_v5 = _183_v5;
        } else {
          _183_v5 = _183_v5;
          let _204_v23;
          let _nw23 = new _module.C0();
          _nw23.__ctor();
          _204_v23 = _nw23;
          let _205_v24;
          let _init2 = ((_206_v10, _207_v0) => function (_208_i3) {
            return _module.D3.create_DC11(_206_v10, _207_v0, true);
          })(_187_v10, _176_v0);
          let _nw24 = Array((new BigNumber(25)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw24.length); _i0_2++) {
            _nw24[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _205_v24 = _nw24;
          let _209_v25;
          _209_v25 = _module.D3.create_DC11(_187_v10, _176_v0, _176_v0);
          let _index20 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_205_v24).length));
          (_205_v24)[_index20] = _209_v25;
          let _pat_let_tv5 = _180_v4;
          let _pat_let_tv6 = _180_v4;
          let _pat_let_tv7 = _176_v0;
          let _pat_let_tv8 = _190_globalState;
          let _index21 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_205_v24).length));
          (_205_v24)[_index21] = function (_pat_let2_0) {
            return function (_210_dt__update__tmp_h0) {
              return function (_pat_let3_0) {
                return function (_211_dt__update_hcf19_h0) {
                  return _module.D3.create_DC11(_211_dt__update_hcf19_h0, (_210_dt__update__tmp_h0).dtor_cf20, (_210_dt__update__tmp_h0).dtor_cf21);
                }(_pat_let3_0);
              }(_module.__default.fm3((_pat_let_tv6)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_pat_let_tv5).length))], _pat_let_tv7, _pat_let_tv8));
            }(_pat_let2_0);
          }(_209_v25);
          let _212_v26;
          _212_v26 = _187_v10;
          _212_v26 = _187_v10;
          (_190_globalState).f19 = (((_185_v8).contains((false) && (!(_176_v0)))) ? ((_185_v8).get((false) && (!(_176_v0)))) : (_177_v1));
        }
        let _213_v27;
        _213_v27 = _dafny.Seq.of(_dafny.Set.fromElements(_176_v0, _176_v0), _dafny.Set.fromElements(_176_v0, _176_v0));
        let _214_v28;
        let _nw25 = new _module.C1();
        _nw25.__ctor((_213_v27)[_module.__default.safeIndex(_177_v1, new BigNumber((_213_v27).length))], _177_v1);
        _214_v28 = _nw25;
        let _215_v29;
        _215_v29 = _dafny.Seq.of(_176_v0, !(_176_v0), (_module.D3.create_DC11(_187_v10, _176_v0, _176_v0)).dtor_cf20);
        let _216_v30;
        _216_v30 = _dafny.Map.Empty.slice().updateUnsafe(_215_v29,_214_v28);
        let _index22 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_180_v4).length));
        let _rhs25 = _module.__default.safeModuloInt(new BigNumber((_178_v2).length), _177_v1);
        let _rhs26 = (((_216_v30).contains(_dafny.Seq.update(((false) ? (_215_v29) : (_dafny.Seq.of(_176_v0, _176_v0, _176_v0))), _module.__default.safeIndex((_dafny.ZERO).minus(_177_v1), new BigNumber((((false) ? (_215_v29) : (_dafny.Seq.of(_176_v0, _176_v0, _176_v0)))).length)), _176_v0))) ? ((_216_v30).get(_dafny.Seq.update(((false) ? (_215_v29) : (_dafny.Seq.of(_176_v0, _176_v0, _176_v0))), _module.__default.safeIndex((_dafny.ZERO).minus(_177_v1), new BigNumber((((false) ? (_215_v29) : (_dafny.Seq.of(_176_v0, _176_v0, _176_v0)))).length)), _176_v0))) : (_214_v28));
        let _lhs21 = _180_v4;
        let _lhs22 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_180_v4).length));
        _lhs21[_lhs22] = _rhs25;
        _214_v28 = _rhs26;
      } else {
        if ((new BigNumber((_185_v8).cardinality())).isLessThan(_177_v1)) {
          let _217_v31;
          let _nw26 = new _module.C0();
          _nw26.__ctor();
          _217_v31 = _nw26;
          let _218_v32;
          _218_v32 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_177_v1, _177_v1));
          let _219_v33;
          _219_v33 = _module.D0.create_DC0(_dafny.Seq.Concat(_218_v32, _218_v32));
          let _220_v34;
          _220_v34 = _dafny.Seq.of(true);
          let _rhs27 = ((false) ? (_177_v1) : ((((_220_v34)[_module.__default.safeIndex(_177_v1, new BigNumber((_220_v34).length))]) ? (_177_v1) : (_177_v1))));
          let _rhs28 = (_177_v1).plus(_177_v1);
          let _rhs29 = _217_v31;
          let _rhs30 = _176_v0;
          let _rhs31 = _219_v33;
          let _lhs23 = _190_globalState;
          let _lhs24 = _190_globalState;
          let _lhs25 = _190_globalState;
          _lhs23.f10 = _rhs27;
          _lhs24.f19 = _rhs28;
          _217_v31 = _rhs29;
          _lhs25.f23 = _rhs30;
          _219_v33 = _rhs31;
          _217_v31 = _217_v31;
          _195_v14 = (_195_v14).update(_177_v1, _176_v0);
          let _221_v35;
          _221_v35 = _dafny.Set.fromElements(_176_v0, _176_v0);
          let _222_v36;
          let _nw27 = new _module.C1();
          _nw27.__ctor(_221_v35, (_177_v1).multipliedBy(new BigNumber(107)));
          _222_v36 = _nw27;
          let _223_v37;
          _223_v37 = _dafny.Map.Empty.slice().updateUnsafe(_222_v36,_187_v10);
          _223_v37 = (_223_v37).update(_222_v36, new _dafny.CodePoint('x'.codePointAt(0)));
        } else {
          let _224_v38;
          let _init3 = ((_225_v2) => function (_226_i4) {
            return _225_v2;
          })(_178_v2);
          let _nw28 = Array((new BigNumber(4)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw28.length); _i0_3++) {
            _nw28[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _224_v38 = _nw28;
          let _227_v39;
          _227_v39 = _dafny.Set.fromElements(_module.__default.safeDivisionInt(_177_v1, _177_v1));
          let _228_v40;
          _228_v40 = _dafny.Set.fromElements(_176_v0, _176_v0, _176_v0, _176_v0);
          let _229_v41;
          let _nw29 = new _module.C0();
          _nw29.__ctor();
          _229_v41 = _nw29;
          let _230_v42;
          _230_v42 = _dafny.Map.Empty.slice().updateUnsafe(_229_v41,_176_v0);
          let _231_v43;
          let _nw30 = new _module.C1();
          _nw30.__ctor((_228_v40).Union(_dafny.Set.fromElements((((_230_v42).contains(_229_v41)) ? ((_230_v42).get(_229_v41)) : (_176_v0)), _176_v0, !(_176_v0))), (_229_v41).fm4(_177_v1, _176_v0, _190_globalState));
          _231_v43 = _nw30;
          let _232_v44;
          _232_v44 = _dafny.Seq.of(_231_v43);
          let _rhs32 = _224_v38;
          let _rhs33 = _dafny.Set.fromElements(_231_v43.f27, _module.__default.safeModuloInt(new BigNumber((_232_v44).length), (_178_v2)[_module.__default.safeIndex(new BigNumber((_183_v5).length), new BigNumber((_178_v2).length))]), _177_v1);
          let _rhs34 = _231_v43;
          let _rhs35 = _231_v43.f27;
          let _lhs26 = _190_globalState;
          _224_v38 = _rhs32;
          _227_v39 = _rhs33;
          _231_v43 = _rhs34;
          _lhs26.f19 = _rhs35;
          let _233_v45;
          let _init4 = ((_234_v8) => function (_235_i5) {
            return _module.D4.create_DC13(_234_v8);
          })(_185_v8);
          let _nw31 = Array((new BigNumber(25)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw31.length); _i0_4++) {
            _nw31[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _233_v45 = _nw31;
          let _236_v46;
          let _nw32 = new _module.C1();
          _nw32.__ctor((_231_v43).f26, _177_v1);
          _236_v46 = _nw32;
          let _237_v47;
          _237_v47 = _dafny.Seq.of(_236_v46, _236_v46);
          let _238_v48;
          _238_v48 = _dafny.Map.Empty.slice().updateUnsafe((_237_v47)[_module.__default.safeIndex((_178_v2)[_module.__default.safeIndex(_177_v1, new BigNumber((_178_v2).length))], new BigNumber((_237_v47).length))],_177_v1);
          let _239_v49;
          _239_v49 = _dafny.Map.Empty.slice().updateUnsafe(_176_v0,_233_v45);
          let _rhs36 = (((_239_v49).contains(((_176_v0) ? (false) : (_176_v0)))) ? ((_239_v49).get(((_176_v0) ? (false) : (_176_v0)))) : (_233_v45));
          let _rhs37 = _238_v48;
          let _rhs38 = _231_v43.f27;
          let _lhs27 = _190_globalState;
          _233_v45 = _rhs36;
          _238_v48 = _rhs37;
          _lhs27.f15 = _rhs38;
          let _index23 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_180_v4).length));
          (_180_v4)[_index23] = _231_v43.f27;
          let _index24 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_180_v4).length));
          (_180_v4)[_index24] = _231_v43.f27;
          let _240_v50;
          let _nw33 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Map.Empty);
          _240_v50 = _nw33;
          let _241_v51;
          _241_v51 = _dafny.Map.Empty.slice().updateUnsafe(_176_v0,_177_v1);
          let _index25 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_240_v50).length));
          (_240_v50)[_index25] = _241_v51;
          let _index26 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_240_v50).length));
          (_240_v50)[_index26] = _241_v51;
          let _242_v52;
          _242_v52 = _module.D0.create_DC1((_180_v4)[_module.__default.safeIndex(new BigNumber(323), new BigNumber((_180_v4).length))]);
          let _243_v53;
          _243_v53 = _module.D3.create_DC12(false, _242_v52, _176_v0, _176_v0, _236_v46);
          let _244_v54;
          let _nw34 = Array((new BigNumber(19)).toNumber());
          _nw34[(_dafny.ZERO).toNumber()] = _236_v46;
          _nw34[(_dafny.ONE).toNumber()] = _236_v46;
          _nw34[(new BigNumber(2)).toNumber()] = _236_v46;
          _nw34[(new BigNumber(3)).toNumber()] = _236_v46;
          _nw34[(new BigNumber(4)).toNumber()] = _236_v46;
          _nw34[(new BigNumber(5)).toNumber()] = _236_v46;
          _nw34[(new BigNumber(6)).toNumber()] = _236_v46;
          _nw34[(new BigNumber(7)).toNumber()] = (_243_v53).dtor_cf26;
          _nw34[(new BigNumber(8)).toNumber()] = _236_v46;
          _nw34[(new BigNumber(9)).toNumber()] = _236_v46;
          _nw34[(new BigNumber(10)).toNumber()] = _236_v46;
          _nw34[(new BigNumber(11)).toNumber()] = _236_v46;
          _nw34[(new BigNumber(12)).toNumber()] = ((_176_v0) ? ((_module.D3.create_DC12(_176_v0, _242_v52, _176_v0, _176_v0, _236_v46)).dtor_cf26) : (_236_v46));
          _nw34[(new BigNumber(13)).toNumber()] = _236_v46;
          _nw34[(new BigNumber(14)).toNumber()] = _236_v46;
          _nw34[(new BigNumber(15)).toNumber()] = _236_v46;
          _nw34[(new BigNumber(16)).toNumber()] = _236_v46;
          _nw34[(new BigNumber(17)).toNumber()] = _236_v46;
          _nw34[(new BigNumber(18)).toNumber()] = _236_v46;
          _244_v54 = _nw34;
          let _index27 = _module.__default.safeIndex(new BigNumber(39), new BigNumber((_244_v54).length));
          (_244_v54)[_index27] = _236_v46;
          let _index28 = _module.__default.safeIndex(new BigNumber(39), new BigNumber((_244_v54).length));
          (_244_v54)[_index28] = _236_v46;
        }
        let _245_v55;
        _245_v55 = _module.D2.create_DC8(new BigNumber((_183_v5).length), _176_v0, _177_v1);
        if ((_245_v55).dtor_cf13) {
          let _246_v56;
          _246_v56 = _module.D3.create_DC11(_187_v10, _176_v0, _176_v0);
          let _pat_let_tv9 = _176_v0;
          let _247_v57;
          _247_v57 = _dafny.Seq.of(function (_pat_let4_0) {
            return function (_248_dt__update__tmp_h2) {
              return function (_pat_let5_0) {
                return function (_249_dt__update_hcf21_h0) {
                  return _module.D3.create_DC11((_248_dt__update__tmp_h2).dtor_cf19, (_248_dt__update__tmp_h2).dtor_cf20, _249_dt__update_hcf21_h0);
                }(_pat_let5_0);
              }(_pat_let_tv9);
            }(_pat_let4_0);
          }(_246_v56), _module.__default.fm11(_177_v1, _190_globalState), _module.D3.create_DC11(_187_v10, _176_v0, _176_v0));
          (_190_globalState).f5 = new BigNumber((_dafny.Seq.Concat(_247_v57, _dafny.Seq.Create(_module.__default.abs(new BigNumber(876)), ((_250_v10, _251_v0) => function (_252_i6) {
            return _module.D3.create_DC11(_250_v10, _251_v0, _251_v0);
          })(_187_v10, _176_v0)))).length);
          (_190_globalState).f25 = _176_v0;
          (_190_globalState).f2 = (_177_v1).plus(_module.__default.safeDivisionInt(_177_v1, _177_v1));
          let _253_v58;
          let _init5 = ((_254_v10, _255_v0) => function (_256_i7) {
            return _module.D3.create_DC11(_254_v10, !(_255_v0), !(!(!(_255_v0))));
          })(_187_v10, _176_v0);
          let _nw35 = Array((new BigNumber(3)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw35.length); _i0_5++) {
            _nw35[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _253_v58 = _nw35;
          let _index29 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_253_v58).length));
          (_253_v58)[_index29] = _246_v56;
          let _index30 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_253_v58).length));
          (_253_v58)[_index30] = _module.__default.fm11(_177_v1, _190_globalState);
          let _257_v59;
          _257_v59 = _module.D0.create_DC1(_177_v1);
          let _258_v60;
          _258_v60 = _dafny.Set.fromElements(_176_v0, _176_v0);
          let _259_v61;
          let _nw36 = new _module.C1();
          _nw36.__ctor(_258_v60, _177_v1);
          _259_v61 = _nw36;
          let _260_v62;
          _260_v62 = _module.D3.create_DC12(_176_v0, _257_v59, _176_v0, _176_v0, _259_v61);
          let _261_v63;
          _261_v63 = _dafny.Map.Empty.slice().updateUnsafe(_260_v62,_177_v1);
          let _262_v64;
          _262_v64 = _dafny.Map.Empty.slice().updateUnsafe(_177_v1,_261_v63);
          (_190_globalState).f22 = !((_dafny.Map.Empty.slice().updateUnsafe(_260_v62,_177_v1)).equals((((_262_v64).contains((_dafny.ZERO).minus(new BigNumber((_178_v2).length)))) ? ((_262_v64).get((_dafny.ZERO).minus(new BigNumber((_178_v2).length)))) : (_261_v63))));
        } else {
          let _263_v65;
          _263_v65 = _dafny.Set.fromElements(_176_v0);
          let _264_v66;
          let _nw37 = new _module.C1();
          _nw37.__ctor(_263_v65, _177_v1);
          _264_v66 = _nw37;
          let _265_v67;
          _265_v67 = _module.D0.create_DC1(new BigNumber(-492));
          let _266_v68;
          let _nw38 = Array((new BigNumber(13)).toNumber());
          _nw38[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(28)), function (_267_i8) {
            return new BigNumber((_dafny.Seq.UnicodeFromString("cwxonbmhb")).length);
          });
          _nw38[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_178_v2, _module.__default.fm12(_177_v1, _265_v67, _190_globalState));
          _nw38[(new BigNumber(2)).toNumber()] = _178_v2;
          _nw38[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(_264_v66.f27);
          _nw38[(new BigNumber(4)).toNumber()] = _178_v2;
          _nw38[(new BigNumber(5)).toNumber()] = _178_v2;
          _nw38[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_178_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(762)), ((_268_v8, _269_v0, _270_v66) => function (_271_i9) {
            return new BigNumber(((_268_v8).update(_269_v0, _module.__default.abs((_dafny.ZERO).minus(_270_v66.f27)))).cardinality());
          })(_185_v8, _176_v0, _264_v66)));
          _nw38[(new BigNumber(7)).toNumber()] = _178_v2;
          _nw38[(new BigNumber(8)).toNumber()] = _178_v2;
          _nw38[(new BigNumber(9)).toNumber()] = ((_176_v0) ? (_178_v2) : (_178_v2));
          _nw38[(new BigNumber(10)).toNumber()] = _178_v2;
          _nw38[(new BigNumber(11)).toNumber()] = _178_v2;
          _nw38[(new BigNumber(12)).toNumber()] = _178_v2;
          _266_v68 = _nw38;
          let _index31 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_266_v68).length));
          (_266_v68)[_index31] = _dafny.Seq.Concat(_178_v2, _178_v2);
          let _272_v69;
          _272_v69 = _dafny.Set.fromElements(_264_v66.f27, _177_v1);
          let _273_v70;
          _273_v70 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.update(_178_v2, _module.__default.safeIndex(new BigNumber((_272_v69).length), new BigNumber((_178_v2).length)), _264_v66.f27), _module.__default.safeIndex(_264_v66.f27, new BigNumber((_dafny.Seq.update(_178_v2, _module.__default.safeIndex(new BigNumber((_272_v69).length), new BigNumber((_178_v2).length)), _264_v66.f27)).length)), _177_v1), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-788)), ((_274_v1) => function (_275_i10) {
            return _274_v1;
          })(_177_v1)), _178_v2, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(377)), ((_276_v66) => function (_277_i11) {
            return _276_v66.f27;
          })(_264_v66)), _178_v2));
          let _index32 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_266_v68).length));
          (_266_v68)[_index32] = (_273_v70)[_module.__default.safeIndex(_177_v1, new BigNumber((_273_v70).length))];
          (_190_globalState).f15 = _264_v66.f27;
          let _278_v71;
          _278_v71 = _module.D2.create_DC7(_176_v0);
          let _rhs39 = _278_v71;
          let _rhs40 = (_264_v66.f27).multipliedBy(_177_v1);
          let _lhs28 = _190_globalState;
          _278_v71 = _rhs39;
          _lhs28.f8 = _rhs40;
          (_190_globalState).f22 = _dafny.Seq.IsPrefixOf(_183_v5, _183_v5);
        }
        (_190_globalState).f10 = _module.__default.fm1(_176_v0, _187_v10, _177_v1, _190_globalState);
        (_190_globalState).f19 = _177_v1;
        if ((_185_v8).IsDisjointFrom(_dafny.MultiSet.fromElements(_176_v0, _176_v0, true))) {
          let _279_v72;
          _279_v72 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_183_v5).length),_177_v1);
          (_190_globalState).f23 = !_dafny.Seq.contains(_module.__default.fm13(_190_globalState), _279_v72);
          let _280_v73;
          _280_v73 = _dafny.Map.Empty.slice().updateUnsafe(_185_v8,new BigNumber(363));
          _280_v73 = (_280_v73).update(_dafny.MultiSet.fromElements(_176_v0), _177_v1);
          let _index33 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_180_v4).length));
          (_180_v4)[_index33] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_183_v5, _183_v5), _module.__default.safeIndex(_177_v1, new BigNumber((_dafny.Seq.Concat(_183_v5, _183_v5)).length)), _187_v10)).length);
          let _index34 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_180_v4).length));
          (_180_v4)[_index34] = _177_v1;
          let _281_v74;
          _281_v74 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_180_v4)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_180_v4).length))],(_dafny.ZERO).minus(_177_v1)),_176_v0);
          let _rhs41 = (_281_v74).Merge(_281_v74);
          let _rhs42 = new BigNumber(465);
          let _rhs43 = _dafny.Seq.Concat(_dafny.Seq.Concat(_178_v2, _178_v2), _dafny.Seq.Create(_module.__default.abs(new BigNumber(23)), ((_282_v1) => function (_283_i12) {
            return _282_v1;
          })(_177_v1)));
          let _rhs44 = _module.__default.fm0((_176_v0) === (_176_v0), _dafny.Map.Empty.slice().updateUnsafe(_176_v0,_module.__default.fm0(_176_v0, _196_v15, _190_globalState)), _190_globalState);
          let _rhs45 = (_180_v4)[_module.__default.safeIndex(new BigNumber(451), new BigNumber((_180_v4).length))];
          let _lhs29 = _190_globalState;
          let _lhs30 = _190_globalState;
          let _lhs31 = _190_globalState;
          _281_v74 = _rhs41;
          _lhs29.f15 = _rhs42;
          _178_v2 = _rhs43;
          _lhs30.f25 = _rhs44;
          _lhs31.f2 = _rhs45;
          let _284_v75;
          _284_v75 = _dafny.Set.fromElements(_176_v0);
          let _285_v76;
          let _nw39 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _285_v76 = _nw39;
          let _286_v77;
          _286_v77 = _dafny.Set.fromElements(_285_v76);
          let _287_v78;
          _287_v78 = _dafny.Seq.of(_176_v0);
          let _288_v79;
          _288_v79 = _module.D0.create_DC2(_176_v0, _286_v77, !(_176_v0), _176_v0, new BigNumber((_287_v78).length));
          let _289_v80;
          let _nw40 = new _module.C1();
          _nw40.__ctor(_284_v75, (_288_v79).dtor_cf6);
          _289_v80 = _nw40;
          let _290_v81;
          _290_v81 = _dafny.Set.fromElements(_289_v80);
          (_190_globalState).f25 = (new BigNumber((_290_v81).length)).isLessThanOrEqualTo(new BigNumber(549));
        } else {
          _178_v2 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(762)), ((_291_v1) => function (_292_i13) {
            return _291_v1;
          })(_177_v1));
          let _293_v82;
          _293_v82 = _dafny.Set.fromElements(_176_v0, _176_v0, _176_v0);
          let _294_v83;
          _294_v83 = _187_v10;
          let _295_v84;
          _295_v84 = _dafny.Seq.of(_176_v0, _176_v0);
          let _296_v85;
          _296_v85 = _dafny.Set.fromElements(_177_v1);
          let _rhs46 = _180_v4;
          let _rhs47 = _module.__default.fm0((true) === (_module.__default.fm0(_176_v0, _196_v15, _190_globalState)), (_196_v15).Merge(_196_v15), _190_globalState);
          let _rhs48 = (_185_v8).IsSubsetOf((_dafny.MultiSet.FromArray(_295_v84)).Union(_185_v8));
          let _rhs49 = _dafny.Set.fromElements(_176_v0, _176_v0, !(_176_v0), (_295_v84)[_module.__default.safeIndex(new BigNumber((_296_v85).length), new BigNumber((_295_v84).length))]);
          let _rhs50 = _187_v10;
          let _lhs32 = _190_globalState;
          let _lhs33 = _190_globalState;
          _lhs32.f11 = _rhs46;
          _lhs33.f23 = _rhs47;
          _176_v0 = _rhs48;
          _293_v82 = _rhs49;
          _294_v83 = _rhs50;
          let _297_v86;
          let _nw41 = Array((new BigNumber(9)).toNumber());
          _nw41[(_dafny.ZERO).toNumber()] = _187_v10;
          _nw41[(_dafny.ONE).toNumber()] = _187_v10;
          _nw41[(new BigNumber(2)).toNumber()] = _187_v10;
          _nw41[(new BigNumber(3)).toNumber()] = _187_v10;
          _nw41[(new BigNumber(4)).toNumber()] = _187_v10;
          _nw41[(new BigNumber(5)).toNumber()] = _187_v10;
          _nw41[(new BigNumber(6)).toNumber()] = new _dafny.CodePoint('k'.codePointAt(0));
          _nw41[(new BigNumber(7)).toNumber()] = _187_v10;
          _nw41[(new BigNumber(8)).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
          _297_v86 = _nw41;
          let _index35 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_297_v86).length));
          (_297_v86)[_index35] = new _dafny.CodePoint('h'.codePointAt(0));
          let _index36 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_297_v86).length));
          (_297_v86)[_index36] = _187_v10;
          (_190_globalState).f22 = _176_v0;
          let _298_v87;
          _298_v87 = _dafny.MultiSet.fromElements(_177_v1);
          let _299_v88;
          _299_v88 = _dafny.Seq.of(_298_v87, _298_v87, _298_v87);
          let _300_v89;
          _300_v89 = _module.D0.create_DC0(_299_v88);
          _300_v89 = _module.__default.fm14(_190_globalState);
        }
      }
      (_190_globalState).f22 = !(((false) ? (_196_v15) : (_196_v15))).equals(_196_v15);
      let _301_v90;
      _301_v90 = _dafny.Set.fromElements(_177_v1, _177_v1);
      if ((_301_v90).IsSubsetOf((_301_v90).Difference(_301_v90))) {
        if (!(_176_v0) || (_176_v0)) {
          let _302_v91;
          _302_v91 = _dafny.MultiSet.fromElements(_177_v1);
          let _303_v92;
          _303_v92 = _module.D2.create_DC5(_dafny.Seq.UnicodeFromString("tvlnsoaai"));
          let _304_v93;
          _304_v93 = _module.D4.create_DC14(_302_v91, _177_v1, _303_v92);
          (_190_globalState).f2 = ((_176_v0) ? (_177_v1) : ((_304_v93).dtor_cf29));
          (_190_globalState).f22 = _module.__default.fm0(_176_v0, _196_v15, _190_globalState);
          let _305_v94;
          let _nw42 = new _module.C0();
          _nw42.__ctor();
          _305_v94 = _nw42;
          let _306_v95;
          _306_v95 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-221)), ((_307_v10) => function (_308_i14) {
            return _307_v10;
          })(_187_v10)),_dafny.Seq.Concat(_183_v5, _dafny.Seq.UnicodeFromString("jkvceykyp")));
          _306_v95 = (_306_v95).update(_183_v5, _183_v5);
          let _309_v96;
          _309_v96 = _dafny.Map.Empty.slice().updateUnsafe(_177_v1,_dafny.Seq.UnicodeFromString("tgejkifdl"));
          let _310_v97;
          _310_v97 = _dafny.Map.Empty.slice().updateUnsafe(_309_v96,_177_v1);
          _310_v97 = (_310_v97).update(_309_v96, _177_v1);
        } else {
          _187_v10 = new _dafny.CodePoint('w'.codePointAt(0));
          (_190_globalState).f8 = _module.__default.safeDivisionInt(_177_v1, _177_v1);
          let _index37 = _module.__default.safeIndex(new BigNumber(876), new BigNumber((_186_v9).length));
          (_186_v9)[_index37] = _176_v0;
          let _index38 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_186_v9).length));
          (_186_v9)[_index38] = !(_176_v0) || (_176_v0);
          let _index39 = _module.__default.safeIndex(new BigNumber(876), new BigNumber((_186_v9).length));
          let _index40 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_186_v9).length));
          let _rhs51 = _176_v0;
          let _rhs52 = _187_v10;
          let _rhs53 = _176_v0;
          let _rhs54 = _176_v0;
          let _lhs34 = _190_globalState;
          let _lhs35 = _186_v9;
          let _lhs36 = _module.__default.safeIndex(new BigNumber(876), new BigNumber((_186_v9).length));
          let _lhs37 = _186_v9;
          let _lhs38 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_186_v9).length));
          _lhs34.f23 = _rhs51;
          _187_v10 = _rhs52;
          _lhs35[_lhs36] = _rhs53;
          _lhs37[_lhs38] = _rhs54;
          let _311_v98;
          let _312_v99;
          let _out2;
          let _out3;
          let _outcollector1 = _module.__default.m1((((_196_v15).contains((_186_v9)[_module.__default.safeIndex(new BigNumber(876), new BigNumber((_186_v9).length))])) ? ((_196_v15).get((_186_v9)[_module.__default.safeIndex(new BigNumber(876), new BigNumber((_186_v9).length))])) : (_176_v0)), _module.__default.safeDivisionInt(_177_v1, new BigNumber(140)), _190_globalState);
          _out2 = _outcollector1[0];
          _out3 = _outcollector1[1];
          _311_v98 = _out2;
          _312_v99 = _out3;
          let _index41 = _module.__default.safeIndex(new BigNumber(876), new BigNumber((_186_v9).length));
          (_186_v9)[_index41] = (_module.__default.fm15(_190_globalState)).IsDisjointFrom(_dafny.Set.fromElements(_311_v98, _177_v1, _177_v1, (_dafny.ZERO).minus(_177_v1)));
        }
        let _313_v100;
        let _nw43 = new _module.C0();
        _nw43.__ctor();
        _313_v100 = _nw43;
        let _314_v101;
        _314_v101 = _dafny.Seq.of(((_176_v0) ? (_176_v0) : (_176_v0)), _176_v0, !(_176_v0));
        let _315_v102;
        _315_v102 = _dafny.Map.Empty.slice().updateUnsafe(_177_v1,_177_v1);
        let _rhs55 = (_313_v100).fm4((((_315_v102).contains(_177_v1)) ? ((_315_v102).get(_177_v1)) : (_177_v1)), !(_177_v1).isEqualTo(new BigNumber((_314_v101).length)), _190_globalState);
        let _rhs56 = (_dafny.ZERO).minus(_177_v1);
        let _rhs57 = _dafny.Seq.update(_314_v101, _module.__default.safeIndex(_177_v1, new BigNumber((_314_v101).length)), !(_195_v14).contains(_177_v1));
        let _rhs58 = _177_v1;
        let _lhs39 = _190_globalState;
        let _lhs40 = _190_globalState;
        let _lhs41 = _190_globalState;
        _lhs39.f10 = _rhs55;
        _lhs40.f8 = _rhs56;
        _314_v101 = _rhs57;
        _lhs41.f19 = _rhs58;
        let _316_v103;
        _316_v103 = _dafny.Set.fromElements(_313_v100, _313_v100, _313_v100, _313_v100, _313_v100);
        let _317_v104;
        _317_v104 = _dafny.MultiSet.fromElements(_316_v103);
        let _318_v105;
        _318_v105 = _dafny.Seq.of(_316_v103);
        _317_v104 = _dafny.MultiSet.FromArray(_dafny.Seq.update(_318_v105, _module.__default.safeIndex(_177_v1, new BigNumber((_318_v105).length)), _316_v103));
        (_190_globalState).f8 = new BigNumber((_183_v5).length);
      } else {
        _183_v5 = _183_v5;
        let _319_v106;
        _319_v106 = _module.D3.create_DC10((new BigNumber((_dafny.Seq.of(_177_v1)).length)).isLessThanOrEqualTo(_177_v1), (_dafny.ZERO).minus(_177_v1), _176_v0);
        _319_v106 = _319_v106;
        let _index42 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_186_v9).length));
        (_186_v9)[_index42] = _176_v0;
        let _320_v107;
        _320_v107 = _module.D2.create_DC8(_177_v1, _176_v0, _177_v1);
        let _index43 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_186_v9).length));
        let _rhs59 = (_176_v0) === (_176_v0);
        let _rhs60 = _module.__default.fm1(((_176_v0) ? (_176_v0) : (_176_v0)), _187_v10, _module.__default.safeModuloInt(new BigNumber((_183_v5).length), _177_v1), _190_globalState);
        let _rhs61 = (_320_v107).dtor_cf12;
        let _lhs42 = _186_v9;
        let _lhs43 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_186_v9).length));
        let _lhs44 = _190_globalState;
        _lhs42[_lhs43] = _rhs59;
        _lhs44.f2 = _rhs60;
        _177_v1 = _rhs61;
        let _321_v108;
        _321_v108 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(685),_180_v4);
        _321_v108 = (_321_v108).update((new BigNumber((_183_v5).length)).minus(_177_v1), _180_v4);
        let _322_v109;
        _322_v109 = _dafny.Set.fromElements((_186_v9)[_module.__default.safeIndex(new BigNumber(924), new BigNumber((_186_v9).length))]);
        let _323_v110;
        let _nw44 = new _module.C1();
        _nw44.__ctor(_322_v109, _177_v1);
        _323_v110 = _nw44;
      }
      let _324_v111;
      let _325_v112;
      let _out4;
      let _out5;
      let _outcollector2 = _module.__default.m1(false, new BigNumber(292), _190_globalState);
      _out4 = _outcollector2[0];
      _out5 = _outcollector2[1];
      _324_v111 = _out4;
      _325_v112 = _out5;
      if (_176_v0) {
        (_190_globalState).f23 = _176_v0;
        let _326_v113;
        _326_v113 = _dafny.MultiSet.fromElements(_177_v1, _177_v1);
        _326_v113 = (_326_v113).Union(_dafny.MultiSet.fromElements(_324_v111, _177_v1));
        if ((_301_v90).IsProperSubsetOf(_301_v90)) {
          let _index44 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_180_v4).length));
          (_180_v4)[_index44] = _324_v111;
          let _index45 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_180_v4).length));
          (_180_v4)[_index45] = (_177_v1).minus(((_176_v0) ? (new BigNumber(714)) : ((_dafny.ZERO).minus(_324_v111))));
          let _index46 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_186_v9).length));
          (_186_v9)[_index46] = (_module.__default.fm15(_190_globalState)).IsProperSubsetOf(function () {
            let _coll7 = new _dafny.Set();
            for (const _compr_7 of _dafny.IntegerRange(new BigNumber(722), new BigNumber(354))) {
              let _327_v114 = _compr_7;
              if (((new BigNumber(722)).isLessThanOrEqualTo(_327_v114)) && ((_327_v114).isLessThan(new BigNumber(354)))) {
                _coll7.add(_module.__default.safeModuloInt(_327_v114, new BigNumber(-455)));
              }
            }
            return _coll7;
          }());
          let _index47 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_186_v9).length));
          (_186_v9)[_index47] = _325_v112;
          _196_v15 = (_196_v15).update(_module.__default.fm0(_325_v112, _196_v15, _190_globalState), _176_v0);
          let _328_v115;
          _328_v115 = _dafny.Set.fromElements(_325_v112);
          let _329_v116;
          let _nw45 = new _module.C1();
          _nw45.__ctor(_328_v115, new BigNumber((_328_v115).length));
          _329_v116 = _nw45;
          let _330_v117;
          _330_v117 = _dafny.MultiSet.fromElements(_329_v116);
          let _331_v118;
          let _nw46 = new _module.C1();
          _nw46.__ctor(_dafny.Set.fromElements(_176_v0), (((_330_v117).contains(_329_v116)) ? ((_330_v117).get(_329_v116)) : (_177_v1)));
          _331_v118 = _nw46;
          (_190_globalState).f2 = _329_v116.f27;
        } else {
          (_190_globalState).f23 = _176_v0;
          _186_v9 = _186_v9;
          let _332_v119;
          let _nw47 = new _module.C0();
          _nw47.__ctor();
          _332_v119 = _nw47;
          _186_v9 = _186_v9;
          let _333_v120;
          let _334_v121;
          let _out6;
          let _out7;
          let _outcollector3 = _module.__default.m1(_176_v0, (new BigNumber(-429)).plus(_324_v111), _190_globalState);
          _out6 = _outcollector3[0];
          _out7 = _outcollector3[1];
          _333_v120 = _out6;
          _334_v121 = _out7;
        }
        (_190_globalState).f23 = _325_v112;
        let _335_v122;
        let _336_v123;
        let _out8;
        let _out9;
        let _outcollector4 = _module.__default.m1(_176_v0, new BigNumber((_dafny.MultiSet.fromElements(_324_v111)).cardinality()), _190_globalState);
        _out8 = _outcollector4[0];
        _out9 = _outcollector4[1];
        _335_v122 = _out8;
        _336_v123 = _out9;
      } else {
        let _337_v124;
        _337_v124 = _module.D0.create_DC1(_177_v1);
        (_190_globalState).f5 = _module.__default.fm1(_325_v112, _187_v10, (_337_v124).dtor_cf1, _190_globalState);
        _195_v14 = (_195_v14).update(new BigNumber(438), _176_v0);
        let _338_v125;
        _338_v125 = _dafny.Map.Empty.slice().updateUnsafe(_180_v4,(_dafny.ZERO).minus(new BigNumber(((_185_v8).Intersect(_185_v8)).cardinality())));
        _338_v125 = (_338_v125).update(_180_v4, _324_v111);
        _186_v9 = _186_v9;
        let _339_v126;
        let _340_v127;
        let _out10;
        let _out11;
        let _outcollector5 = _module.__default.m1(_325_v112, new BigNumber((((_176_v0) ? (_module.__default.fm10(_177_v1, new BigNumber(288), _176_v0, _176_v0, _190_globalState)) : (_183_v5))).length), _190_globalState);
        _out10 = _outcollector5[0];
        _out11 = _outcollector5[1];
        _339_v126 = _out10;
        _340_v127 = _out11;
      }
      let _341_v128;
      _341_v128 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_176_v0, _325_v112, _176_v0, _325_v112, _325_v112));
      let _342_v129;
      _342_v129 = _dafny.Set.fromElements(_185_v8, _185_v8, _185_v8, (_341_v128)[_module.__default.safeIndex(_177_v1, new BigNumber((_341_v128).length))], _185_v8);
      let _343_v130;
      _343_v130 = _dafny.Map.Empty.slice().updateUnsafe(_185_v8,(_324_v111).isEqualTo(new BigNumber((_342_v129).length)));
      _343_v130 = (_343_v130).update(_185_v8, _176_v0);
      _183_v5 = _183_v5;
      let _344_v131;
      _344_v131 = _module.D3.create_DC10(_176_v0, _177_v1, false);
      let _345_v132;
      _345_v132 = _dafny.Seq.of(true);
      let _346_v133;
      let _347_v134;
      let _out12;
      let _out13;
      let _outcollector6 = _module.__default.m1(!((_344_v131).dtor_cf16), new BigNumber(((_dafny.MultiSet.FromArray(_345_v132)).Intersect(_185_v8)).cardinality()), _190_globalState);
      _out12 = _outcollector6[0];
      _out13 = _outcollector6[1];
      _346_v133 = _out12;
      _347_v134 = _out13;
      _347_v134 = _347_v134;
      let _348_v135;
      _348_v135 = _dafny.MultiSet.fromElements(new BigNumber(60));
      let _349_v136;
      _349_v136 = _dafny.Map.Empty.slice().updateUnsafe(_348_v135,_346_v133);
      let _rhs62 = _349_v136;
      let _rhs63 = _module.__default.safeModuloInt((_177_v1).multipliedBy(_324_v111), (_324_v111).plus(_324_v111));
      let _rhs64 = _346_v133;
      let _rhs65 = (_dafny.ZERO).minus(_177_v1);
      let _lhs45 = _190_globalState;
      let _lhs46 = _190_globalState;
      _349_v136 = _rhs62;
      _324_v111 = _rhs63;
      _lhs45.f19 = _rhs64;
      _lhs46.f5 = _rhs65;
      if (_347_v134) {
        let _350_v137;
        _350_v137 = _dafny.Set.fromElements(_176_v0);
        let _351_v138;
        let _nw48 = new _module.C1();
        _nw48.__ctor((_350_v137).Union(_350_v137), _346_v133);
        _351_v138 = _nw48;
        (_190_globalState).f25 = _module.__default.fm0(_325_v112, (_196_v15).Merge(_196_v15), _190_globalState);
        let _352_v139;
        let _nw49 = Array((new BigNumber(8)).toNumber()).fill([]);
        _352_v139 = _nw49;
        let _353_v140;
        let _nw50 = Array((new BigNumber(10)).toNumber()).fill([]);
        _353_v140 = _nw50;
        let _index48 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_352_v139).length));
        (_352_v139)[_index48] = _353_v140;
        let _index49 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_352_v139).length));
        (_352_v139)[_index49] = _353_v140;
        let _354_v141;
        let _355_v142;
        let _out14;
        let _out15;
        let _outcollector7 = (_351_v138).m0(((!(_347_v134)) ? (_176_v0) : (_176_v0)), _190_globalState);
        _out14 = _outcollector7[0];
        _out15 = _outcollector7[1];
        _354_v141 = _out14;
        _355_v142 = _out15;
        (_190_globalState).f23 = _347_v134;
      } else {
        (_190_globalState).f5 = _346_v133;
        let _356_v143;
        _356_v143 = _module.D2.create_DC8(_346_v133, _325_v112, new BigNumber((_183_v5).length));
        let _357_v144;
        _357_v144 = _dafny.Map.Empty.slice().updateUnsafe(_325_v112,_177_v1);
        let _rhs66 = _346_v133;
        let _rhs67 = _dafny.Seq.of(_347_v134, _325_v112, ((_module.__default.fm0(_347_v134, _196_v15, _190_globalState)) ? (_347_v134) : (!(_347_v134))), (((_356_v143).dtor_cf13) ? (_325_v112) : (_325_v112)), (_346_v133).isLessThan(_module.__default.fm1(_325_v112, new _dafny.CodePoint('w'.codePointAt(0)), new BigNumber((_357_v144).length), _190_globalState)));
        let _rhs68 = _176_v0;
        let _lhs47 = _190_globalState;
        _lhs47.f5 = _rhs66;
        _345_v132 = _rhs67;
        _176_v0 = _rhs68;
        let _358_v145;
        _358_v145 = _dafny.Set.fromElements(false, true);
        let _359_v146;
        let _nw51 = new _module.C1();
        _nw51.__ctor(_358_v145, _177_v1);
        _359_v146 = _nw51;
        let _pat_let_tv10 = _183_v5;
        let _360_v147;
        _360_v147 = _module.D3.create_DC12(_325_v112, function (_pat_let6_0) {
  return function (_361_dt__update__tmp_h4) {
    return function (_pat_let7_0) {
      return function (_362_dt__update_hcf1_h0) {
        return _module.D0.create_DC1(_362_dt__update_hcf1_h0);
      }(_pat_let7_0);
    }(new BigNumber((_pat_let_tv10).length));
  }(_pat_let6_0);
}(_module.D0.create_DC1(_346_v133)), (_344_v131).dtor_cf18, false, _359_v146);
        _325_v112 = (_360_v147).dtor_cf22;
        let _363_v149;
        _363_v149 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_177_v1, _346_v133)),new BigNumber((function () {
          let _coll8 = new _dafny.Map();
          for (const _compr_8 of _dafny.IntegerRange(new BigNumber(967), new BigNumber(310))) {
            let _364_v148 = _compr_8;
            if (((new BigNumber(967)).isLessThanOrEqualTo(_364_v148)) && ((_364_v148).isLessThan(new BigNumber(310)))) {
              _coll8.push([(_364_v148).plus(_346_v133),_347_v134]);
            }
          }
          return _coll8;
        }()).length));
        let _365_v150;
        _365_v150 = _dafny.Seq.of(_183_v5);
        _363_v149 = (_363_v149).update(_177_v1, new BigNumber((_365_v150).length));
        let _366_v151;
        let _nw52 = new _module.C1();
        _nw52.__ctor(((true) ? (_358_v145) : (_dafny.Set.fromElements(_347_v134, _176_v0))), _177_v1);
        _366_v151 = _nw52;
      }
      let _367_v152;
      _367_v152 = _module.D0.create_DC0(_module.__default.fm16(_177_v1, _190_globalState));
      let _pat_let_tv11 = _187_v10;
      let _pat_let_tv12 = _187_v10;
      let _pat_let_tv13 = _187_v10;
      let _pat_let_tv14 = _325_v112;
      let _pat_let_tv15 = _187_v10;
      let _pat_let_tv16 = _187_v10;
      let _source3 = function (_source4) {
        if (_source4.is_DC1) {
          let _368___mcc_h1 = (_source4).cf1;
          let _369_cf1 = _368___mcc_h1;
          return _pat_let_tv11;
        } else if (_source4.is_DC2) {
          let _370___mcc_h2 = (_source4).cf2;
          let _371___mcc_h3 = (_source4).cf3;
          let _372___mcc_h4 = (_source4).cf4;
          let _373___mcc_h5 = (_source4).cf5;
          let _374___mcc_h6 = (_source4).cf6;
          let _375_cf6 = _374___mcc_h6;
          let _376_cf5 = _373___mcc_h5;
          let _377_cf4 = _372___mcc_h4;
          let _378_cf3 = _371___mcc_h3;
          let _379_cf2 = _370___mcc_h2;
          return _pat_let_tv12;
        } else if (_source4.is_DC3) {
          let _380___mcc_h7 = (_source4).cf7;
          let _381_cf7 = _380___mcc_h7;
          return _pat_let_tv13;
        } else {
          let _382___mcc_h8 = (_source4).cf0;
          let _383_cf0 = _382___mcc_h8;
          if (_pat_let_tv14) {
            return _pat_let_tv15;
          } else {
            return _pat_let_tv16;
          }
        }
      }(_367_v152);
      let _384___mcc_h0 = _source3;
      let _385_cf8 = _384___mcc_h0;
      let _386_v153;
      _386_v153 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),_325_v112);
      _386_v153 = (_386_v153).update(_385_cf8, !(true) || (_176_v0));
      let _387_v154;
      _387_v154 = _dafny.Set.fromElements((_module.D2.create_DC7(_347_v134)).dtor_cf11);
      let _388_v155;
      let _nw53 = new _module.C1();
      _nw53.__ctor(_387_v154, _177_v1);
      _388_v155 = _nw53;
      let _389_v156;
      let _390_v157;
      let _out16;
      let _out17;
      let _outcollector8 = _module.__default.m1(_325_v112, _324_v111, _190_globalState);
      _out16 = _outcollector8[0];
      _out17 = _outcollector8[1];
      _389_v156 = _out16;
      _390_v157 = _out17;
      (_190_globalState).f5 = _346_v133;
      if (true) {
        let _391_v158;
        _391_v158 = _module.D2.create_DC5(_183_v5);
        let _392_v159;
        _392_v159 = _dafny.Seq.of(_183_v5, _183_v5, _183_v5);
        let _393_v160;
        let _nw54 = Array((new BigNumber(12)).toNumber());
        _nw54[(_dafny.ZERO).toNumber()] = _183_v5;
        _nw54[(_dafny.ONE).toNumber()] = _183_v5;
        _nw54[(new BigNumber(2)).toNumber()] = _183_v5;
        _nw54[(new BigNumber(3)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(929)), ((_394_v10) => function (_395_i15) {
          return _394_v10;
        })(_187_v10));
        _nw54[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_183_v5, _183_v5);
        _nw54[(new BigNumber(5)).toNumber()] = _183_v5;
        _nw54[(new BigNumber(6)).toNumber()] = _183_v5;
        _nw54[(new BigNumber(7)).toNumber()] = _183_v5;
        _nw54[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat((_391_v158).dtor_cf9, _dafny.Seq.Create(_module.__default.abs(new BigNumber(879)), ((_396_v10) => function (_397_i16) {
          return _396_v10;
        })(_187_v10))), _module.__default.safeIndex(_346_v133, new BigNumber((_dafny.Seq.Concat((_391_v158).dtor_cf9, _dafny.Seq.Create(_module.__default.abs(new BigNumber(879)), ((_398_v10) => function (_399_i16) {
          return _398_v10;
        })(_187_v10)))).length)), new _dafny.CodePoint('w'.codePointAt(0)));
        _nw54[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_183_v5, _183_v5);
        _nw54[(new BigNumber(10)).toNumber()] = _dafny.Seq.update(_dafny.Seq.update((_392_v159)[_module.__default.safeIndex(_324_v111, new BigNumber((_392_v159).length))], _module.__default.safeIndex(new BigNumber((_178_v2).length), new BigNumber(((_392_v159)[_module.__default.safeIndex(_324_v111, new BigNumber((_392_v159).length))]).length)), _187_v10), _module.__default.safeIndex(_177_v1, new BigNumber((_dafny.Seq.update((_392_v159)[_module.__default.safeIndex(_324_v111, new BigNumber((_392_v159).length))], _module.__default.safeIndex(new BigNumber((_178_v2).length), new BigNumber(((_392_v159)[_module.__default.safeIndex(_324_v111, new BigNumber((_392_v159).length))]).length)), _187_v10)).length)), _187_v10);
        _nw54[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_183_v5, _183_v5);
        _393_v160 = _nw54;
        _393_v160 = _393_v160;
        let _400_v161;
        let _nw55 = Array((new BigNumber(23)).toNumber());
        _400_v161 = _nw55;
        let _index50 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_186_v9).length));
        (_186_v9)[_index50] = _347_v134;
        let _index51 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_186_v9).length));
        let _rhs69 = _400_v161;
        let _rhs70 = true;
        let _rhs71 = _dafny.Seq.update(_183_v5, _module.__default.safeIndex(_346_v133, new BigNumber((_183_v5).length)), _187_v10);
        let _rhs72 = _dafny.Seq.Concat(_dafny.Seq.Concat(_183_v5, _183_v5), _183_v5);
        let _lhs48 = _186_v9;
        let _lhs49 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_186_v9).length));
        _400_v161 = _rhs69;
        _lhs48[_lhs49] = _rhs70;
        _183_v5 = _rhs71;
        _183_v5 = _rhs72;
        let _401_v162;
        let _init6 = ((_402_v2) => function (_403_i17) {
          return _dafny.Seq.Concat(_402_v2, _402_v2);
        })(_178_v2);
        let _nw56 = Array((new BigNumber(20)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw56.length); _i0_6++) {
          _nw56[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _401_v162 = _nw56;
        let _index52 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_401_v162).length));
        (_401_v162)[_index52] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(70)), ((_404_v1) => function (_405_i18) {
          return (_dafny.ZERO).minus(_404_v1);
        })(_177_v1));
        let _406_v163;
        _406_v163 = _dafny.Set.fromElements(_347_v134, _347_v134);
        let _407_v164;
        let _nw57 = new _module.C1();
        _nw57.__ctor(_406_v163, _177_v1);
        _407_v164 = _nw57;
        let _408_v165;
        _408_v165 = _dafny.Map.Empty.slice().updateUnsafe(_324_v111,_407_v164);
        let _409_v166;
        _409_v166 = _dafny.Seq.of(_408_v165);
        let _index53 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_401_v162).length));
        (_401_v162)[_index53] = _dafny.Seq.of(_324_v111, new BigNumber((_409_v166).length));
        _407_v164 = _407_v164;
        if (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_183_v5, _183_v5), _183_v5)) {
          let _410_v167;
          let _411_v168;
          let _out18;
          let _out19;
          let _outcollector9 = _module.__default.m1((_347_v134) || (_176_v0), _177_v1, _190_globalState);
          _out18 = _outcollector9[0];
          _out19 = _outcollector9[1];
          _410_v167 = _out18;
          _411_v168 = _out19;
          _347_v134 = _dafny.areEqual(_dafny.Seq.Concat((_401_v162)[_module.__default.safeIndex(new BigNumber(630), new BigNumber((_401_v162).length))], _178_v2), _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(_410_v167, new BigNumber(270)), _178_v2), _module.__default.safeIndex(new BigNumber((_345_v132).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_410_v167, new BigNumber(270)), _178_v2)).length)), _324_v111));
          let _412_v169;
          let _nw58 = new _module.C0();
          _nw58.__ctor();
          _412_v169 = _nw58;
          let _413_v170;
          _413_v170 = _dafny.Seq.of(_412_v169);
          _413_v170 = _413_v170;
          (_190_globalState).f19 = (_dafny.ZERO).minus((_177_v1).minus(new BigNumber(254)));
          (_190_globalState).f25 = _176_v0;
        } else {
          let _414_v171;
          _414_v171 = _dafny.Map.Empty.slice().updateUnsafe(_176_v0,_dafny.Seq.update(_345_v132, _module.__default.safeIndex(new BigNumber(115), new BigNumber((_345_v132).length)), _325_v112));
          let _415_v172;
          _415_v172 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_178_v2).length),(_414_v171).Merge(_414_v171));
          _415_v172 = (_415_v172).update(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_324_v111,!((_186_v9)[_module.__default.safeIndex(new BigNumber(184), new BigNumber((_186_v9).length))]))).update(_177_v1, _347_v134)).length), _dafny.Map.Empty.slice().updateUnsafe(_347_v134,_345_v132));
          let _index54 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_186_v9).length));
          (_186_v9)[_index54] = !((_407_v164.f27).isLessThanOrEqualTo(_346_v133));
          let _416_v173;
          let _nw59 = Array((new BigNumber(17)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _416_v173 = _nw59;
          _416_v173 = _416_v173;
          let _417_v174;
          let _nw60 = new _module.C0();
          _nw60.__ctor();
          _417_v174 = _nw60;
          let _418_v175;
          _418_v175 = _dafny.Seq.of(_417_v174);
          let _419_v176;
          let _nw61 = Array((new BigNumber(25)).toNumber());
          _nw61[(_dafny.ZERO).toNumber()] = _417_v174;
          _nw61[(_dafny.ONE).toNumber()] = _417_v174;
          _nw61[(new BigNumber(2)).toNumber()] = _417_v174;
          _nw61[(new BigNumber(3)).toNumber()] = _417_v174;
          _nw61[(new BigNumber(4)).toNumber()] = _417_v174;
          _nw61[(new BigNumber(5)).toNumber()] = _417_v174;
          _nw61[(new BigNumber(6)).toNumber()] = _417_v174;
          _nw61[(new BigNumber(7)).toNumber()] = _417_v174;
          _nw61[(new BigNumber(8)).toNumber()] = _417_v174;
          _nw61[(new BigNumber(9)).toNumber()] = _417_v174;
          _nw61[(new BigNumber(10)).toNumber()] = _417_v174;
          _nw61[(new BigNumber(11)).toNumber()] = _417_v174;
          _nw61[(new BigNumber(12)).toNumber()] = _417_v174;
          _nw61[(new BigNumber(13)).toNumber()] = _417_v174;
          _nw61[(new BigNumber(14)).toNumber()] = _417_v174;
          _nw61[(new BigNumber(15)).toNumber()] = _417_v174;
          _nw61[(new BigNumber(16)).toNumber()] = ((_347_v134) ? (_417_v174) : (_417_v174));
          _nw61[(new BigNumber(17)).toNumber()] = _417_v174;
          _nw61[(new BigNumber(18)).toNumber()] = (_418_v175)[_module.__default.safeIndex(new BigNumber((_345_v132).length), new BigNumber((_418_v175).length))];
          _nw61[(new BigNumber(19)).toNumber()] = _417_v174;
          _nw61[(new BigNumber(20)).toNumber()] = _417_v174;
          _nw61[(new BigNumber(21)).toNumber()] = _417_v174;
          _nw61[(new BigNumber(22)).toNumber()] = _417_v174;
          _nw61[(new BigNumber(23)).toNumber()] = _417_v174;
          _nw61[(new BigNumber(24)).toNumber()] = _417_v174;
          _419_v176 = _nw61;
          let _index55 = _module.__default.safeIndex(new BigNumber(123), new BigNumber((_419_v176).length));
          (_419_v176)[_index55] = _417_v174;
          let _index56 = _module.__default.safeIndex(new BigNumber(123), new BigNumber((_419_v176).length));
          let _rhs73 = _324_v111;
          let _rhs74 = (_346_v133).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber((_418_v175).length), new BigNumber((_178_v2).length), new BigNumber(-511), _407_v164.f27, _324_v111)).length));
          let _rhs75 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(563)), ((_420_v10) => function (_421_i19) {
            return _420_v10;
          })(_187_v10));
          let _rhs76 = _417_v174;
          let _rhs77 = _347_v134;
          let _lhs50 = _190_globalState;
          let _lhs51 = _407_v164;
          let _lhs52 = _419_v176;
          let _lhs53 = _module.__default.safeIndex(new BigNumber(123), new BigNumber((_419_v176).length));
          let _lhs54 = _190_globalState;
          _lhs50.f8 = _rhs73;
          _lhs51.f27 = _rhs74;
          _183_v5 = _rhs75;
          _lhs52[_lhs53] = _rhs76;
          _lhs54.f25 = _rhs77;
          let _422_v177;
          let _423_v178;
          let _out20;
          let _out21;
          let _outcollector10 = _module.__default.m1(!((_185_v8).IsProperSubsetOf(_185_v8)), (_177_v1).plus(_407_v164.f27), _190_globalState);
          _out20 = _outcollector10[0];
          _out21 = _outcollector10[1];
          _422_v177 = _out20;
          _423_v178 = _out21;
        }
      } else {
        let _424_v179;
        _424_v179 = _dafny.Set.fromElements(!(true), _176_v0, _176_v0);
        let _425_v180;
        let _nw62 = new _module.C1();
        _nw62.__ctor(_424_v179, _324_v111);
        _425_v180 = _nw62;
        let _426_v181;
        _426_v181 = _dafny.MultiSet.fromElements(_425_v180, _425_v180);
        if (((_426_v181).Difference(_426_v181)).IsDisjointFrom(_426_v181)) {
          let _427_v182;
          _427_v182 = _dafny.Seq.of((_348_v135).Intersect(_348_v135));
          _427_v182 = _427_v182;
          let _428_v183;
          let _nw63 = new _module.C0();
          _nw63.__ctor();
          _428_v183 = _nw63;
          _186_v9 = _186_v9;
          let _429_v184;
          _429_v184 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_183_v5, _module.__default.safeIndex(_425_v180.f27, new BigNumber((_183_v5).length)), _187_v10),_176_v0);
          (_190_globalState).f22 = (((_429_v184).contains(_183_v5)) ? ((_429_v184).get(_183_v5)) : (_325_v112));
          let _430_v185;
          let _nw64 = new _module.C0();
          _nw64.__ctor();
          _430_v185 = _nw64;
        } else {
          let _431_v186;
          _431_v186 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber(625), (((_348_v135).contains(new BigNumber(((_425_v180).f26).length))) ? ((_348_v135).get(new BigNumber(((_425_v180).f26).length))) : (_324_v111))),_dafny.Seq.update(_183_v5, _module.__default.safeIndex(_324_v111, new BigNumber((_183_v5).length)), _187_v10));
          _431_v186 = _431_v186;
          let _432_v187;
          let _init7 = ((_433_v8) => function (_434_i20) {
            return _433_v8;
          })(_185_v8);
          let _nw65 = Array((new BigNumber(8)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw65.length); _i0_7++) {
            _nw65[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _432_v187 = _nw65;
          let _435_v188;
          _435_v188 = _dafny.Map.Empty.slice().updateUnsafe(_432_v187,_425_v180.f27);
          _435_v188 = (_435_v188).update(_432_v187, (_324_v111).plus(_324_v111));
          _183_v5 = _dafny.Seq.Concat(_dafny.Seq.Concat(_183_v5, _183_v5), _183_v5);
          let _436_v189;
          _436_v189 = _module.D2.create_DC6((_425_v180.f27).isEqualTo(_177_v1));
          let _rhs78 = _176_v0;
          let _rhs79 = _module.__default.fm17(_190_globalState);
          let _lhs55 = _190_globalState;
          _lhs55.f23 = _rhs78;
          _436_v189 = _rhs79;
          let _index57 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_180_v4).length));
          (_180_v4)[_index57] = _346_v133;
          let _index58 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_180_v4).length));
          (_180_v4)[_index58] = (_dafny.ZERO).minus(_module.__default.fm1(true, _187_v10, _346_v133, _190_globalState));
        }
        let _rhs80 = _177_v1;
        let _rhs81 = (_324_v111).multipliedBy(_324_v111);
        let _lhs56 = _190_globalState;
        let _lhs57 = _425_v180;
        _lhs56.f8 = _rhs80;
        _lhs57.f27 = _rhs81;
        let _437_v190;
        _437_v190 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),(_425_v180.f27).plus(_346_v133));
        let _438_v191;
        _438_v191 = _dafny.Map.Empty.slice().updateUnsafe(!(_176_v0),_324_v111);
        let _439_v192;
        _439_v192 = _dafny.Seq.of(_425_v180, _425_v180);
        _437_v190 = _module.__default.fm18((((_438_v191).contains(_176_v0)) ? ((_438_v191).get(_176_v0)) : (_425_v180.f27)), _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_425_v180), _439_v192), _190_globalState);
        let _440_v193;
        _440_v193 = _module.D0.create_DC1(new BigNumber((_195_v14).length));
        if (!(!((new BigNumber((_196_v15).length)).isLessThanOrEqualTo(new BigNumber((_module.__default.fm12(new BigNumber((_183_v5).length), _440_v193, _190_globalState)).length)))) || (_347_v134)) {
          let _441_v194;
          _441_v194 = _module.D2.create_DC5(_183_v5);
          let _442_v195;
          _442_v195 = _module.D4.create_DC14(_348_v135, _425_v180.f27, _441_v194);
          (_190_globalState).f23 = (_324_v111).isEqualTo((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(605), (_442_v195).dtor_cf29)));
          let _443_v196;
          let _444_v197;
          let _out22;
          let _out23;
          let _outcollector11 = _module.__default.m1(_347_v134, _425_v180.f27, _190_globalState);
          _out22 = _outcollector11[0];
          _out23 = _outcollector11[1];
          _443_v196 = _out22;
          _444_v197 = _out23;
          let _445_v198;
          let _nw66 = new _module.C0();
          _nw66.__ctor();
          _445_v198 = _nw66;
          let _rhs82 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt((((_348_v135).contains(_346_v133)) ? ((_348_v135).get(_346_v133)) : (_425_v180.f27)), _324_v111)));
          let _rhs83 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("rfeampgsi"), _183_v5), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(573)), ((_446_v10) => function (_447_i21) {
            return _446_v10;
          })(_187_v10)), (_441_v194).dtor_cf9));
          let _rhs84 = (_177_v1).plus(new BigNumber(-503));
          let _rhs85 = _445_v198;
          let _rhs86 = _445_v198;
          let _lhs58 = _190_globalState;
          _177_v1 = _rhs82;
          _183_v5 = _rhs83;
          _lhs58.f19 = _rhs84;
          _445_v198 = _rhs85;
          _445_v198 = _rhs86;
          (_190_globalState).f22 = (_425_v180).fm6(_dafny.Seq.of(_367_v152), _dafny.Seq.Concat(_183_v5, _dafny.Seq.UnicodeFromString("u")), _185_v8, _190_globalState);
          let _index59 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_186_v9).length));
          (_186_v9)[_index59] = !(!_dafny.Seq.contains(_183_v5, _187_v10));
          let _448_v199;
          _448_v199 = _dafny.Map.Empty.slice().updateUnsafe(_348_v135,_178_v2);
          let _449_v200;
          _449_v200 = _module.D2.create_DC8(_324_v111, _176_v0, _346_v133);
          let _index60 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_186_v9).length));
          let _rhs87 = _187_v10;
          let _rhs88 = _324_v111;
          let _rhs89 = new BigNumber((((_325_v112) ? (_448_v199) : (_448_v199))).length);
          let _rhs90 = (_449_v200).dtor_cf13;
          let _rhs91 = _177_v1;
          let _lhs59 = _190_globalState;
          let _lhs60 = _186_v9;
          let _lhs61 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_186_v9).length));
          let _lhs62 = _190_globalState;
          _187_v10 = _rhs87;
          _324_v111 = _rhs88;
          _lhs59.f2 = _rhs89;
          _lhs60[_lhs61] = _rhs90;
          _lhs62.f10 = _rhs91;
        } else {
          let _rhs92 = _425_v180;
          let _rhs93 = _341_v128;
          _425_v180 = _rhs92;
          _341_v128 = _rhs93;
          let _450_v201;
          _450_v201 = _module.D2.create_DC8(_346_v133, _325_v112, _425_v180.f27);
          let _451_v202;
          let _452_v203;
          let _out24;
          let _out25;
          let _outcollector12 = _module.__default.m1(false, (_450_v201).dtor_cf12, _190_globalState);
          _out24 = _outcollector12[0];
          _out25 = _outcollector12[1];
          _451_v202 = _out24;
          _452_v203 = _out25;
          (_190_globalState).f2 = new BigNumber((_183_v5).length);
          let _index61 = _module.__default.safeIndex(new BigNumber(60), new BigNumber((_186_v9).length));
          (_186_v9)[_index61] = _176_v0;
          let _index62 = _module.__default.safeIndex(new BigNumber(60), new BigNumber((_186_v9).length));
          (_186_v9)[_index62] = _452_v203;
          _195_v14 = (_195_v14).Merge(_195_v14);
        }
        let _453_v204;
        _453_v204 = _dafny.MultiSet.fromElements(_367_v152);
        let _454_v205;
        _454_v205 = _dafny.Map.Empty.slice().updateUnsafe(_425_v180,_453_v204);
        let _455_v206;
        _455_v206 = _dafny.Seq.of(((_dafny.MultiSet.FromArray(_178_v2)).update(_425_v180.f27, _module.__default.abs(_324_v111))).update(_425_v180.f27, _module.__default.abs(new BigNumber(154))), _348_v135);
        let _pat_let_tv17 = _455_v206;
        let _456_v207;
        _456_v207 = _dafny.MultiSet.fromElements(function (_pat_let8_0) {
          return function (_457_dt__update__tmp_h5) {
            return function (_pat_let9_0) {
              return function (_458_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_458_dt__update_hcf0_h0);
              }(_pat_let9_0);
            }(_pat_let_tv17);
          }(_pat_let8_0);
        }(_367_v152), _367_v152);
        let _459_v208;
        _459_v208 = _dafny.Seq.of(_module.__default.fm19(_190_globalState), (_456_v207));
        _454_v205 = (_454_v205).update(_425_v180, (_459_v208)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("xbifwaa")).length), new BigNumber((_459_v208).length))]);
      }
      let _460_v209;
      let _nw67 = new _module.C0();
      _nw67.__ctor();
      _460_v209 = _nw67;
      _460_v209 = _460_v209;
      let _461_v210;
      _461_v210 = _module.D3.create_DC11(new _dafny.CodePoint('r'.codePointAt(0)), false, _176_v0);
      let _hi1 = _177_v1;
      for (let _462_i22 = (_module.__default.fm1(!((_461_v210).dtor_cf20), _187_v10, new BigNumber((_dafny.MultiSet.fromElements(_346_v133)).cardinality()), _190_globalState)).minus(_324_v111); _462_i22.isLessThan(_hi1); _462_i22 = _462_i22.plus(_dafny.ONE)) {
        let _463_v211;
        let _nw68 = new _module.C0();
        _nw68.__ctor();
        _463_v211 = _nw68;
        _463_v211 = _463_v211;
        let _464_v212;
        let _465_v213;
        let _out26;
        let _out27;
        let _outcollector13 = _module.__default.m1(_176_v0, (_344_v131).dtor_cf17, _190_globalState);
        _out26 = _outcollector13[0];
        _out27 = _outcollector13[1];
        _464_v212 = _out26;
        _465_v213 = _out27;
        let _466_v214;
        _466_v214 = _dafny.Set.fromElements(_465_v213);
        let _467_v215;
        let _nw69 = new _module.C1();
        _nw69.__ctor(_466_v214, new BigNumber((_183_v5).length));
        _467_v215 = _nw69;
        _346_v133 = (_dafny.ZERO).minus((new BigNumber(578)).plus(_177_v1));
      }
      process.stdout.write(_dafny.toString(_176_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_177_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_178_v2, _dafny.Seq.of(new BigNumber(17)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_179_v3).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(new BigNumber(17))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v4)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v4)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v4)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v4)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v4)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v4)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v4)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v4)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_183_v5).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_184_v6).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("cb")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_v8).equals(_dafny.MultiSet.fromElements(true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v9)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v9)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v9)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v9)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v9)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v9)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v9)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v9)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v9)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v9)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v9)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_187_v10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v11).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v12).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(17),_dafny.MultiSet.fromElements(new _dafny.CodePoint('u'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_190_globalState).f0).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(new BigNumber(17))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_190_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_190_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_190_globalState).f6, _dafny.Seq.of(new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_190_globalState.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_190_globalState.f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f11)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f11)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f11)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f11)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f11)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f11)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f11)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f11)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f11)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f11)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f11)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f11)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f11)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f11)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f11)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f11)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f11)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f11)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f11)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f11)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_190_globalState).f12).equals(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("cb")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_190_globalState).f13).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_190_globalState.f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_globalState.f16).equals(_dafny.MultiSet.fromElements(false, false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_190_globalState).f17, _dafny.Seq.of(new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_190_globalState).f18)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_190_globalState).f18)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_190_globalState).f18)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_190_globalState).f18)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_190_globalState).f18)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_190_globalState).f18)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_190_globalState).f18)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_190_globalState).f18)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_190_globalState).f18)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_190_globalState).f18)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_190_globalState).f18)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_190_globalState.f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_190_globalState).f20).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_190_globalState).f21).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(17),_dafny.MultiSet.fromElements(new _dafny.CodePoint('u'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_190_globalState.f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_190_globalState.f23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_190_globalState).f24).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_190_globalState.f25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_194_v13).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_v14).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-17),true).updateUnsafe(new BigNumber(438),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v15).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_301_v90).equals(_dafny.Set.fromElements(new BigNumber(17)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_324_v111));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_325_v112));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_341_v128, _dafny.Seq.of(_dafny.MultiSet.fromElements(false, false, false, false, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_342_v129).equals(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(true, true, true), _dafny.MultiSet.fromElements(false, false, false, false, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_343_v130).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true, true, true),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_344_v131).dtor_cf16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_344_v131).dtor_cf17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_344_v131).dtor_cf18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_345_v132, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_346_v133));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_347_v134));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_348_v135).equals(_dafny.MultiSet.fromElements(new BigNumber(60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_349_v136).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(60)),new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_367_v152).dtor_cf0, _dafny.Seq.of(_dafny.MultiSet.fromElements(), _dafny.MultiSet.fromElements(new BigNumber(6), new BigNumber(119), new BigNumber(-456)), _dafny.MultiSet.fromElements(new BigNumber(9)), _dafny.MultiSet.fromElements(new BigNumber(508), new BigNumber(-985))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_461_v210).dtor_cf19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_461_v210).dtor_cf20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_461_v210).dtor_cf21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC2(cf2, cf3, cf4, cf5, cf6) {
      let $dt = new D0(1);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC3(cf7) {
      let $dt = new D0(2);
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(3);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC0() { return this.$tag === 3; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf2 === other.cf2 && _dafny.areEqual(this.cf3, other.cf3) && this.cf4 === other.cf4 && this.cf5 === other.cf5 && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf7, other.cf7);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D0.Default();
        }
      };
    }
  }

  $module.D1 = class D1 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC4(cf8) {
      let $dt = new D1(0);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf8, other.cf8);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return new _dafny.CodePoint('D'.codePointAt(0));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC6(cf10) {
      let $dt = new D2(0);
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC7(cf11) {
      let $dt = new D2(1);
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC8(cf12, cf13, cf14) {
      let $dt = new D2(2);
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC5(cf9) {
      let $dt = new D2(3);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get is_DC5() { return this.$tag === 3; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC5" + "(" + this.cf9.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf10 === other.cf10;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf11 === other.cf11;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf12, other.cf12) && this.cf13 === other.cf13 && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf9, other.cf9);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC10(cf16, cf17, cf18) {
      let $dt = new D3(0);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC11(cf19, cf20, cf21) {
      let $dt = new D3(1);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC12(cf22, cf23, cf24, cf25, cf26) {
      let $dt = new D3(2);
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC9(cf15) {
      let $dt = new D3(3);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get is_DC9() { return this.$tag === 3; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC12" + "(" + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf15) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf16 === other.cf16 && _dafny.areEqual(this.cf17, other.cf17) && this.cf18 === other.cf18;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf19, other.cf19) && this.cf20 === other.cf20 && this.cf21 === other.cf21;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf22 === other.cf22 && _dafny.areEqual(this.cf23, other.cf23) && this.cf24 === other.cf24 && this.cf25 === other.cf25 && this.cf26 === other.cf26;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf15, other.cf15);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(false, _dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC14(cf28, cf29, cf30) {
      let $dt = new D4(0);
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC13(cf27) {
      let $dt = new D4(1);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf27) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf28, other.cf28) && _dafny.areEqual(this.cf29, other.cf29) && _dafny.areEqual(this.cf30, other.cf30);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf27, other.cf27);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC14(_dafny.MultiSet.Empty, _dafny.ZERO, _module.D2.Default());
    }
    static Rtd() {
      return class {
        static get Default() {
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC15(cf31) {
      let $dt = new D5(0);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.MultiSet.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC17() {
      let $dt = new D6(0);
      return $dt;
    }
    static create_DC16(cf32) {
      let $dt = new D6(1);
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC18(cf33) {
      let $dt = new D6(2);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17";
      } else if (this.$tag === 1) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf32 === other.cf32;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf33, other.cf33);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC17();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC20(cf35, cf36, cf37, cf38, cf39) {
      let $dt = new D7(0);
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC21(cf40, cf41) {
      let $dt = new D7(1);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC19(cf34) {
      let $dt = new D7(2);
      $dt.cf34 = cf34;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf34() { return this.cf34; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf35) + ", " + this.cf36.toVerbatimString(true) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf34) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf35 === other.cf35 && _dafny.areEqual(this.cf36, other.cf36) && _dafny.areEqual(this.cf37, other.cf37) && this.cf38 === other.cf38 && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf40 === other.cf40 && this.cf41 === other.cf41;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf34, other.cf34);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC20(false, _dafny.Seq.UnicodeFromString(""), _dafny.ZERO, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC22(cf42) {
      let $dt = new D8(0);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf42) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf42, other.cf42);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Map.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f2 = _dafny.ZERO;
      this.f5 = _dafny.ZERO;
      this.f8 = _dafny.ZERO;
      this.f10 = _dafny.ZERO;
      this.f11 = [];
      this.f15 = _dafny.ZERO;
      this.f16 = _dafny.MultiSet.Empty;
      this.f19 = _dafny.ZERO;
      this.f22 = false;
      this.f23 = false;
      this.f25 = false;
      this._f0 = _dafny.Map.Empty;
      this._f1 = _dafny.ZERO;
      this._f3 = _dafny.ZERO;
      this._f4 = false;
      this._f6 = _dafny.Seq.UnicodeFromString("");
      this._f7 = _dafny.ZERO;
      this._f9 = false;
      this._f12 = _dafny.Set.Empty;
      this._f13 = _dafny.Seq.UnicodeFromString("");
      this._f14 = false;
      this._f17 = _dafny.Seq.UnicodeFromString("");
      this._f18 = [];
      this._f20 = _dafny.Seq.UnicodeFromString("");
      this._f21 = _dafny.Map.Empty;
      this._f24 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this).f8 = f8;
      (_this)._f9 = f9;
      (_this).f10 = f10;
      (_this).f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this).f15 = f15;
      (_this).f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this).f19 = f19;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      (_this).f22 = f22;
      (_this).f23 = f23;
      (_this)._f24 = f24;
      (_this).f25 = f25;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm4(p0, p1, globalState) {
      let _this = this;
      return (_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(152)), function (_468_i0) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      })).length), new BigNumber(860))).minus(new BigNumber(661));
    };
    fm7(p0, p1, p2, globalState) {
      let _this = this;
      return function () {
        let _coll9 = new _dafny.Map();
        for (const _compr_9 of (_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-108)), _dafny.Seq.of((_module.D0.create_DC1(new BigNumber(160))).dtor_cf1, new BigNumber(51), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(574))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length)))).Elements) {
          let _469_v0 = _compr_9;
          if (_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-108)), _dafny.Seq.of((_module.D0.create_DC1(new BigNumber(160))).dtor_cf1, new BigNumber(51), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(574))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length))), _469_v0)) {
            _coll9.push([(_469_v0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("nhtgllpi")).length)),!((new BigNumber((_dafny.Seq.UnicodeFromString("sntppqvt")).length)).isLessThan((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(742)), function (_470_i0) {
              return new _dafny.CodePoint('n'.codePointAt(0));
            })).length))))]);
          }
        }
        return _coll9;
      }();
    };
    fm8(p0, globalState) {
      let _this = this;
      return (new BigNumber((((true) ? (_dafny.Seq.of(new BigNumber(-934), new BigNumber(520))) : (_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("papx")).length)),false)).length))))).length)).plus((_dafny.ZERO).minus(((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(783)), function (_471_i0) {
        return _module.D0.create_DC1(new BigNumber((_dafny.Seq.UnicodeFromString("ua")).length));
      })).length))).minus(new BigNumber(826))));
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this.f27 = _dafny.ZERO;
      this._f26 = _dafny.Set.Empty;
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f26, f27) {
      let _this = this;
      (_this)._f26 = f26;
      (_this).f27 = f27;
      return;
    }
    fm4(p0, p1, globalState) {
      let _this = this;
      return _this.f27;
    };
    fm5(globalState) {
      let _this = this;
      return (((true) ? ((new _dafny.CodePoint('s'.codePointAt(0)))) : (new _dafny.CodePoint('h'.codePointAt(0)))));
    };
    fm6(p0, p1, p2, globalState) {
      let _this = this;
      return ((_this.f27).plus(_this.f27)).isLessThanOrEqualTo(_this.f27);
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _472_v0;
      _472_v0 = new _dafny.CodePoint('x'.codePointAt(0));
      let _473_v1;
      _473_v1 = _dafny.Map.Empty.slice().updateUnsafe(_this.f27,_472_v0);
      _473_v1 = (_473_v1).update(_module.__default.fm1(true, _472_v0, _this.f27, globalState), _472_v0);
      let _474_v2;
      let _nw70 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
      _474_v2 = _nw70;
      let _475_v3;
      _475_v3 = _dafny.Set.fromElements(_474_v2, _474_v2);
      if (((_475_v3).Intersect(_475_v3)).IsSubsetOf(_475_v3)) {
        let _476_v4;
        _476_v4 = _dafny.MultiSet.fromElements(new BigNumber(860));
        let _477_v5;
        _477_v5 = _dafny.Seq.of(_476_v4, _476_v4);
        let _478_v6;
        _478_v6 = _module.D0.create_DC0(_477_v5);
        let _479_v7;
        _479_v7 = _dafny.Seq.UnicodeFromString("ypmgf");
        let _480_v8;
        _480_v8 = _dafny.MultiSet.fromElements(p0, p0);
        let _481_v9;
        _481_v9 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _482_v10;
        _482_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm4(_this.f27, _module.__default.fm0((_this).fm6(_dafny.Seq.of(_478_v6), _479_v7, _480_v8, globalState), _481_v9, globalState), globalState),p0);
        let _index63 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_474_v2).length));
        (_474_v2)[_index63] = _this.f27;
        let _index64 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_474_v2).length));
        let _rhs94 = !(!(_dafny.Set.fromElements(_this.f27)).contains(new BigNumber(931))) || (p0);
        let _rhs95 = (((_dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f27, _this.f27, new BigNumber((_480_v8).cardinality())))).equals(_476_v4)) ? (_this.f27) : (_this.f27));
        let _rhs96 = _dafny.Map.Empty.slice().updateUnsafe(_this.f27,p0);
        let _rhs97 = new BigNumber(-345);
        let _lhs63 = globalState;
        let _lhs64 = globalState;
        let _lhs65 = _474_v2;
        let _lhs66 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_474_v2).length));
        _lhs63.f23 = _rhs94;
        _lhs64.f19 = _rhs95;
        _482_v10 = _rhs96;
        _lhs65[_lhs66] = _rhs97;
        let _483_v11;
        let _nw71 = Array((new BigNumber(5)).toNumber());
        _nw71[(_dafny.ZERO).toNumber()] = _479_v7;
        _nw71[(_dafny.ONE).toNumber()] = _479_v7;
        _nw71[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("wtybqap"), _module.__default.safeIndex(_this.f27, new BigNumber((_dafny.Seq.UnicodeFromString("wtybqap")).length)), _472_v0);
        _nw71[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_479_v7, _dafny.Seq.update(_479_v7, _module.__default.safeIndex((_474_v2)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_474_v2).length))], new BigNumber((_479_v7).length)), _472_v0));
        _nw71[(new BigNumber(4)).toNumber()] = _dafny.Seq.UnicodeFromString("vmqrsdga");
        _483_v11 = _nw71;
        _483_v11 = _483_v11;
        _479_v7 = _479_v7;
        (globalState).f23 = p0;
        let _484_v12;
        _484_v12 = _dafny.Map.Empty.slice().updateUnsafe(_this.f27,new BigNumber((_479_v7).length));
        _484_v12 = (_484_v12).update(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(154)), ((_485_v0) => function (_486_i0) {
          return _485_v0;
        })(_472_v0))).length), new BigNumber(((_this).f26).length));
      } else {
        let _487_v13;
        _487_v13 = _module.D0.create_DC2(p0, _475_v3, p0, p0, _this.f27);
        (globalState).f22 = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_487_v13).dtor_cf6,_this.f27)).length)).isEqualTo((_dafny.ZERO).minus(_this.f27));
        let _488_v14;
        let _nw72 = Array((new BigNumber(20)).toNumber()).fill(false);
        _488_v14 = _nw72;
        let _index65 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_488_v14).length));
        (_488_v14)[_index65] = p0;
        let _489_v15;
        _489_v15 = _dafny.Seq.UnicodeFromString("np");
        let _index66 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_488_v14).length));
        let _rhs98 = _488_v14;
        let _rhs99 = (new BigNumber((_475_v3).length)).isLessThanOrEqualTo(new BigNumber((_489_v15).length));
        let _lhs67 = _488_v14;
        let _lhs68 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_488_v14).length));
        _488_v14 = _rhs98;
        _lhs67[_lhs68] = _rhs99;
        (_this).f27 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_module.__default.fm1((_488_v14)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((_488_v14).length))], _472_v0, _this.f27, globalState)), _this.f27);
        let _490_v16;
        let _nw73 = Array((new BigNumber(13)).toNumber());
        _nw73[(_dafny.ZERO).toNumber()] = _472_v0;
        _nw73[(_dafny.ONE).toNumber()] = _472_v0;
        _nw73[(new BigNumber(2)).toNumber()] = _472_v0;
        _nw73[(new BigNumber(3)).toNumber()] = _472_v0;
        _nw73[(new BigNumber(4)).toNumber()] = (_this).fm5(globalState);
        _nw73[(new BigNumber(5)).toNumber()] = _472_v0;
        _nw73[(new BigNumber(6)).toNumber()] = _472_v0;
        _nw73[(new BigNumber(7)).toNumber()] = _472_v0;
        _nw73[(new BigNumber(8)).toNumber()] = _472_v0;
        _nw73[(new BigNumber(9)).toNumber()] = _472_v0;
        _nw73[(new BigNumber(10)).toNumber()] = _472_v0;
        _nw73[(new BigNumber(11)).toNumber()] = _472_v0;
        _nw73[(new BigNumber(12)).toNumber()] = _472_v0;
        _490_v16 = _nw73;
        let _index67 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_490_v16).length));
        (_490_v16)[_index67] = _472_v0;
        let _index68 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_490_v16).length));
        (_490_v16)[_index68] = ((p0) ? (new _dafny.CodePoint('l'.codePointAt(0))) : (_472_v0));
        let _491_v17;
        _491_v17 = _dafny.MultiSet.fromElements(!(p0), (_488_v14)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((_488_v14).length))], p0);
        (globalState).f15 = (((true) || (p0)) ? (_this.f27) : ((((_491_v17).contains(!((_488_v14)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((_488_v14).length))]))) ? ((_491_v17).get(!((_488_v14)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((_488_v14).length))]))) : (new BigNumber((_dafny.Seq.update(_489_v15, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("sc")).length), new BigNumber((_489_v15).length)), _472_v0)).length)))));
      }
      if (p0) {
        let _492_v18;
        _492_v18 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-390)), function (_493_i1) {
          return (_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f27));
        })));
        let _494_v19;
        _494_v19 = _dafny.Seq.of(new BigNumber(48));
        let _495_v20;
        _495_v20 = _dafny.Set.fromElements(_494_v19);
        _492_v18 = (_492_v18).update(p0, _495_v20);
        let _496_v21;
        _496_v21 = _dafny.Seq.UnicodeFromString("uflnveram");
        (globalState).f8 = (_this).fm4(new BigNumber((_496_v21).length), p0, globalState);
        let _497_v22;
        _497_v22 = _dafny.MultiSet.fromElements(_this.f27);
        let _498_v23;
        _498_v23 = _dafny.MultiSet.fromElements((((_497_v22).contains(_this.f27)) ? ((_497_v22).get(_this.f27)) : (new BigNumber((_496_v21).length))));
        let _499_v24;
        _499_v24 = _dafny.Seq.of(_498_v23);
        let _500_v25;
        _500_v25 = _module.D0.create_DC0(_dafny.Seq.Concat(_499_v24, _dafny.Seq.of((_499_v24)[_module.__default.safeIndex(_this.f27, new BigNumber((_499_v24).length))])));
        let _source5 = _500_v25;
        if (_source5.is_DC1) {
          let _501___mcc_h0 = (_source5).cf1;
          let _502_cf1 = _501___mcc_h0;
          let _503_v26;
          let _nw74 = Array((_dafny.ONE).toNumber()).fill(false);
          _503_v26 = _nw74;
          let _index69 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_503_v26).length));
          (_503_v26)[_index69] = p0;
          let _index70 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_503_v26).length));
          (_503_v26)[_index70] = p0;
          let _index71 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_503_v26).length));
          let _rhs100 = !((_503_v26)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_503_v26).length))]);
          let _rhs101 = (_503_v26)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_503_v26).length))];
          let _lhs69 = _503_v26;
          let _lhs70 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_503_v26).length));
          let _lhs71 = globalState;
          _lhs69[_lhs70] = _rhs100;
          _lhs71.f25 = _rhs101;
          let _index72 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_503_v26).length));
          (_503_v26)[_index72] = (((p0) && ((_503_v26)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_503_v26).length))])) ? ((_503_v26)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_503_v26).length))]) : (p0));
          let _504_v27;
          let _nw75 = new _module.C0();
          _nw75.__ctor();
          _504_v27 = _nw75;
          _504_v27 = _504_v27;
        } else if (_source5.is_DC2) {
          let _505___mcc_h1 = (_source5).cf2;
          let _506___mcc_h2 = (_source5).cf3;
          let _507___mcc_h3 = (_source5).cf4;
          let _508___mcc_h4 = (_source5).cf5;
          let _509___mcc_h5 = (_source5).cf6;
          let _510_cf6 = _509___mcc_h5;
          let _511_cf5 = _508___mcc_h4;
          let _512_cf4 = _507___mcc_h3;
          let _513_cf3 = _506___mcc_h2;
          let _514_cf2 = _505___mcc_h1;
          let _515_v28;
          _515_v28 = _dafny.Map.Empty.slice().updateUnsafe(p0,_496_v21);
          let _516_v29;
          _516_v29 = _dafny.Map.Empty.slice().updateUnsafe(_514_cf2,new BigNumber(((((_515_v28).contains(_511_cf5)) ? ((_515_v28).get(_511_cf5)) : (_dafny.Seq.UnicodeFromString("g")))).length));
          let _517_v30;
          _517_v30 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(879)), ((_518_v0) => function (_519_i2) {
            return _518_v0;
          })(_472_v0)), _module.__default.safeIndex(new BigNumber((_516_v29).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(879)), ((_520_v0) => function (_521_i2) {
            return _520_v0;
          })(_472_v0))).length)), _472_v0), _496_v21)).length))),_510_cf6);
          _517_v30 = (_517_v30).update((new BigNumber(611)).minus(new BigNumber((_494_v19).length)), new BigNumber(-674));
          let _522_v31;
          _522_v31 = _dafny.Seq.of(_500_v25, _500_v25);
          let _523_v32;
          _523_v32 = _dafny.MultiSet.fromElements(false, _511_cf5, false);
          let _524_v33;
          _524_v33 = _dafny.Map.Empty.slice().updateUnsafe(_514_cf2,(_this).fm6(_522_v31, _496_v21, _523_v32, globalState));
          let _525_v34;
          _525_v34 = _dafny.Seq.of(_474_v2);
          (globalState).f2 = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_524_v33).length)), (_this.f27).multipliedBy(new BigNumber((_525_v34).length)));
          (globalState).f10 = _this.f27;
          (globalState).f22 = _module.__default.fm0(!((_510_cf6).isEqualTo(_510_cf6)), _module.__default.fm9(new BigNumber(836), globalState), globalState);
        } else if (_source5.is_DC3) {
          let _526___mcc_h6 = (_source5).cf7;
          let _527_cf7 = _526___mcc_h6;
          (globalState).f2 = new BigNumber(744);
          (globalState).f2 = _this.f27;
          let _528_v35;
          let _nw76 = new _module.C0();
          _nw76.__ctor();
          _528_v35 = _nw76;
          let _529_v36;
          _529_v36 = _dafny.Seq.of(_494_v19, _494_v19, ((!(p0)) ? (_494_v19) : (_dafny.Seq.of(_this.f27))), _494_v19);
          (globalState).f5 = new BigNumber((_529_v36).length);
        } else {
          let _530___mcc_h7 = (_source5).cf0;
          let _531_cf0 = _530___mcc_h7;
          let _532_v37;
          let _nw77 = new _module.C0();
          _nw77.__ctor();
          _532_v37 = _nw77;
          let _533_v38;
          _533_v38 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          _533_v38 = (_533_v38).update(true, p0);
          let _534_v39;
          _534_v39 = new _dafny.CodePoint('p'.codePointAt(0));
          let _535_v40;
          _535_v40 = _dafny.Map.Empty.slice().updateUnsafe(p0,_472_v0);
          let _rhs102 = (((_535_v40).contains(p0)) ? ((_535_v40).get(p0)) : (_472_v0));
          let _rhs103 = _497_v22;
          _534_v39 = _rhs102;
          _498_v23 = _rhs103;
          let _536_v41;
          _536_v41 = _dafny.MultiSet.fromElements(_472_v0, _472_v0, _472_v0, _472_v0, _472_v0);
          let _rhs104 = p0;
          let _rhs105 = p0;
          let _rhs106 = _536_v41;
          let _lhs72 = globalState;
          let _lhs73 = globalState;
          _lhs72.f25 = _rhs104;
          _lhs73.f25 = _rhs105;
          _536_v41 = _rhs106;
        }
        let _537_v42;
        let _nw78 = Array((new BigNumber(8)).toNumber()).fill(false);
        _537_v42 = _nw78;
        let _538_v43;
        _538_v43 = _dafny.Set.fromElements(_537_v42);
        (globalState).f23 = ((_538_v43).Union(_538_v43)).IsDisjointFrom(_538_v43);
        let _539_v44;
        _539_v44 = _dafny.MultiSet.fromElements(p0);
        let _540_v45;
        _540_v45 = _dafny.Map.Empty.slice().updateUnsafe(p0,((p0) ? (new BigNumber((_539_v44).cardinality())) : (_this.f27)));
        _540_v45 = (_540_v45).update((_497_v22).IsDisjointFrom(_497_v22), _this.f27);
      } else {
        if (p0) {
          let _541_v46;
          let _nw79 = Array((new BigNumber(4)).toNumber());
          _541_v46 = _nw79;
          let _542_v47;
          _542_v47 = _dafny.Map.Empty.slice().updateUnsafe(p0,_541_v46);
          let _543_v48;
          _543_v48 = _dafny.MultiSet.fromElements(new BigNumber((_542_v47).length));
          _543_v48 = (_543_v48).Union(_543_v48);
          let _544_v49;
          let _nw80 = new _module.C0();
          _nw80.__ctor();
          _544_v49 = _nw80;
          let _545_v50;
          _545_v50 = _dafny.Seq.UnicodeFromString("l");
          _545_v50 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(416)), ((_546_v0) => function (_547_i3) {
            return _546_v0;
          })(_472_v0)), (_module.D2.create_DC5(_545_v50)).dtor_cf9);
          let _548_v51;
          let _nw81 = new _module.C0();
          _nw81.__ctor();
          _548_v51 = _nw81;
          let _549_v52;
          _549_v52 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _550_v53;
          _550_v53 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_549_v52).Merge(_549_v52)).length),(_this.f27).plus(_this.f27));
          _550_v53 = (_550_v53).update(new BigNumber(792), _this.f27);
        } else {
          let _551_v54;
          _551_v54 = _dafny.MultiSet.fromElements(_this.f27, _this.f27, _this.f27);
          (globalState).f25 = _module.__default.fm0((_551_v54).IsProperSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f27))), _dafny.Map.Empty.slice().updateUnsafe(p0,p0), globalState);
          let _552_v55;
          let _nw82 = Array((new BigNumber(3)).toNumber()).fill([]);
          _552_v55 = _nw82;
          let _553_v56;
          let _nw83 = Array((new BigNumber(3)).toNumber()).fill(false);
          _553_v56 = _nw83;
          let _index73 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_552_v55).length));
          (_552_v55)[_index73] = _553_v56;
          let _index74 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_552_v55).length));
          (_552_v55)[_index74] = _553_v56;
          let _554_v57;
          _554_v57 = _dafny.MultiSet.fromElements(p0);
          (globalState).f16 = ((_dafny.MultiSet.fromElements(false, false)).Difference(_554_v57)).Union(_554_v57);
          (globalState).f22 = p0;
          let _555_v58;
          _555_v58 = _dafny.Seq.of(p0);
          let _556_v59;
          _556_v59 = _dafny.Seq.of(_this.f27);
          let _arr0 = (_552_v55)[_module.__default.safeIndex(new BigNumber(99), new BigNumber((_552_v55).length))];
          let _index75 = _module.__default.safeIndex(new BigNumber(691), new BigNumber(((_552_v55)[_module.__default.safeIndex(new BigNumber(99), new BigNumber((_552_v55).length))]).length));
          _arr0[_index75] = (_555_v58)[_module.__default.safeIndex((_556_v59)[_module.__default.safeIndex(_this.f27, new BigNumber((_556_v59).length))], new BigNumber((_555_v58).length))];
          let _557_v60;
          _557_v60 = _dafny.Map.Empty.slice().updateUnsafe(p0,_472_v0);
          let _arr1 = (_552_v55)[_module.__default.safeIndex(new BigNumber(99), new BigNumber((_552_v55).length))];
          let _index76 = _module.__default.safeIndex(new BigNumber(691), new BigNumber(((_552_v55)[_module.__default.safeIndex(new BigNumber(99), new BigNumber((_552_v55).length))]).length));
          _arr1[_index76] = !(_557_v60).equals((_module.D3.create_DC9(_557_v60)).dtor_cf15);
        }
        let _558_v61;
        _558_v61 = _dafny.MultiSet.fromElements(p0, p0, p0, p0);
        let _559_v62;
        _559_v62 = _module.D4.create_DC13(_558_v61);
        (globalState).f16 = (_559_v62).dtor_cf27;
        let _560_v63;
        _560_v63 = _dafny.Seq.of(p0);
        let _561_v64;
        _561_v64 = _dafny.Seq.of((_560_v63)[_module.__default.safeIndex(_this.f27, new BigNumber((_560_v63).length))], p0, p0);
        let _562_v65;
        _562_v65 = _module.D2.create_DC7(p0);
        _561_v64 = (((_562_v65).dtor_cf11) ? (_dafny.Seq.update(_dafny.Seq.update(_560_v63, _module.__default.safeIndex(_module.__default.fm1(p0, _472_v0, _this.f27, globalState), new BigNumber((_560_v63).length)), false), _module.__default.safeIndex(_this.f27, new BigNumber((_dafny.Seq.update(_560_v63, _module.__default.safeIndex(_module.__default.fm1(p0, _472_v0, _this.f27, globalState), new BigNumber((_560_v63).length)), false)).length)), p0)) : (_561_v64));
        let _563_v66;
        _563_v66 = _dafny.Set.fromElements(_this.f27);
        let _564_v67;
        _564_v67 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _565_v68;
        _565_v68 = _dafny.Seq.UnicodeFromString("eelid");
        let _566_v69;
        _566_v69 = _dafny.Seq.of(_565_v68);
        let _567_v70;
        _567_v70 = _dafny.Seq.of(new BigNumber(((_566_v69)[_module.__default.safeIndex(_this.f27, new BigNumber((_566_v69).length))]).length));
        let _568_v71;
        _568_v71 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.of(_this.f27, _module.__default.fm1(p0, _472_v0, _this.f27, globalState), new BigNumber((_563_v66).length), new BigNumber(-710), new BigNumber((_module.__default.fm10(_this.f27, _this.f27, p0, _module.__default.fm0(p0, _564_v67, globalState), globalState)).length)), _567_v70),_module.D3.create_DC10(p0, _this.f27, p0));
        _568_v71 = _568_v71;
        (globalState).f8 = (_dafny.ZERO).minus(_this.f27);
      }
      let _569_v72;
      let _init8 = ((_570_p0) => function (_571_i4) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(806),_570_p0);
      })(p0);
      let _nw84 = Array((new BigNumber(7)).toNumber());
      for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw84.length); _i0_8++) {
        _nw84[_i0_8] = _init8(new BigNumber(_i0_8));
      }
      _569_v72 = _nw84;
      let _572_v73;
      _572_v73 = _dafny.Map.Empty.slice().updateUnsafe(_this.f27,p0);
      let _index77 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_569_v72).length));
      (_569_v72)[_index77] = _572_v73;
      let _index78 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_569_v72).length));
      (_569_v72)[_index78] = (_572_v73).Merge((_572_v73).Merge(_572_v73));
      let _573_v74;
      let _nw85 = new _module.C0();
      _nw85.__ctor();
      _573_v74 = _nw85;
      (globalState).f23 = p0;
      r0 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_this.f27, _this.f27));
      r1 = _this.f27;
      return [r0, r1];
    }
    get f26() {
      let _this = this;
      return _this._f26;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
