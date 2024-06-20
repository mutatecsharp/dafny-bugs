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
    static fm0(p0, globalState) {
      return (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.of(_module.D0.create_DC1(true, false, new BigNumber(197)))).length)).plus(_module.__default.safeDivisionInt((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("giqwu")).length))), new BigNumber((_dafny.Seq.of(new BigNumber(187), new BigNumber(69), new BigNumber(335))).length))));
    };
    static fm4(p0, p1, globalState) {
      return _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(new BigNumber(450)), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(!(false), true)).length), new BigNumber(-29)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(913)), function (_0_i0) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("ycbj")).length);
      })));
    };
    static fm8(p0, globalState) {
      if (true) {
        if (true) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('u'.codePointAt(0));
        }
      } else {
        return new _dafny.CodePoint('l'.codePointAt(0));
      }
    };
    static fm9(p0, p1, p2, globalState) {
      return _module.D2.create_DC6(new _dafny.CodePoint('l'.codePointAt(0)), (_dafny.ONE).minus(new BigNumber(202)));
    };
    static fm10(p0, p1, p2, globalState) {
      return _dafny.MultiSet.FromArray((((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),true)).length)).isEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(165)), function (_2_i0) {
        return new BigNumber((_dafny.Set.fromElements(true, true)).length);
      })).length))) ? (_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("khvhcxm")).length)), _dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(321),true)).length))))))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(550)), function (_1_i1) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(true, true)).length))).length));
      }))));
    };
    static fm11(p0, p1, p2, globalState) {
      return _dafny.Seq.UnicodeFromString("d");
    };
    static fm13(p0, p1, p2, globalState) {
      let _source0 = _module.D16.create_DC50(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),false), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),true)));
      if (_source0.is_DC51) {
        let _3___mcc_h0 = (_source0).cf94;
        let _4___mcc_h1 = (_source0).cf95;
        let _5___mcc_h2 = (_source0).cf96;
        let _6___mcc_h3 = (_source0).cf97;
        let _7_cf97 = _6___mcc_h3;
        let _8_cf96 = _5___mcc_h2;
        let _9_cf95 = _4___mcc_h1;
        let _10_cf94 = _3___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(356)), function (_11_i0) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("clkt"));
      } else {
        let _12___mcc_h4 = (_source0).cf93;
        let _13_cf93 = _12___mcc_h4;
        return _dafny.Seq.UnicodeFromString("may");
      }
    };
    static fm14(p0, p1, p2, p3, globalState) {
      return _module.D0.create_DC0(true);
    };
    static fm15(p0, p1, globalState) {
      let _source1 = _module.D8.create_DC26(!(!(false)), !(false), true, new BigNumber((_dafny.Set.fromElements(new BigNumber(-666))).length));
      if (_source1.is_DC26) {
        let _14___mcc_h0 = (_source1).cf45;
        let _15___mcc_h1 = (_source1).cf46;
        let _16___mcc_h2 = (_source1).cf47;
        let _17___mcc_h3 = (_source1).cf48;
        let _18_cf48 = _17___mcc_h3;
        let _19_cf47 = _16___mcc_h2;
        let _20_cf46 = _15___mcc_h1;
        let _21_cf45 = _14___mcc_h0;
        if (_19_cf47) {
          return _module.D0.create_DC1(_19_cf47, _19_cf47, _18_cf48);
        } else {
          return _module.D0.create_DC1(_20_cf46, _20_cf46, _18_cf48);
        }
      } else if (_source1.is_DC27) {
        let _22___mcc_h4 = (_source1).cf49;
        let _23___mcc_h5 = (_source1).cf50;
        let _24___mcc_h6 = (_source1).cf51;
        let _25___mcc_h7 = (_source1).cf52;
        let _26_cf52 = _25___mcc_h7;
        let _27_cf51 = _24___mcc_h6;
        let _28_cf50 = _23___mcc_h5;
        let _29_cf49 = _22___mcc_h4;
        return _module.D0.create_DC1(_29_cf49, _29_cf49, new BigNumber((function () {
  let _coll0 = new _dafny.Set();
  for (const _compr_0 of (function () {
    let _coll1 = new _dafny.Set();
    for (const _compr_1 of (function () {
      let _coll2 = new _dafny.Set();
      for (const _compr_2 of (function () {
        let _coll3 = new _dafny.Set();
        for (const _compr_3 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("pisbhaerb"), _dafny.Seq.UnicodeFromString("ejfkrj"), _dafny.Seq.UnicodeFromString("qqubr"))).Elements) {
          let _30_v0 = _compr_3;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("pisbhaerb"), _dafny.Seq.UnicodeFromString("ejfkrj"), _dafny.Seq.UnicodeFromString("qqubr")), _30_v0)) {
            _coll3.add(_30_v0);
          }
        }
        return _coll3;
      }()).Elements) {
        let _31_v1 = _compr_2;
        if ((function () {
          let _coll4 = new _dafny.Set();
          for (const _compr_4 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("pisbhaerb"), _dafny.Seq.UnicodeFromString("ejfkrj"), _dafny.Seq.UnicodeFromString("qqubr"))).Elements) {
            let _32_v0 = _compr_4;
            if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("pisbhaerb"), _dafny.Seq.UnicodeFromString("ejfkrj"), _dafny.Seq.UnicodeFromString("qqubr")), _32_v0)) {
              _coll4.add(_32_v0);
            }
          }
          return _coll4;
        }()).contains(_31_v1)) {
          _coll2.add(_31_v1);
        }
      }
      return _coll2;
    }()).Elements) {
      let _33_v2 = _compr_1;
      if ((function () {
        let _coll5 = new _dafny.Set();
        for (const _compr_5 of (function () {
          let _coll6 = new _dafny.Set();
          for (const _compr_6 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("pisbhaerb"), _dafny.Seq.UnicodeFromString("ejfkrj"), _dafny.Seq.UnicodeFromString("qqubr"))).Elements) {
            let _34_v0 = _compr_6;
            if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("pisbhaerb"), _dafny.Seq.UnicodeFromString("ejfkrj"), _dafny.Seq.UnicodeFromString("qqubr")), _34_v0)) {
              _coll6.add(_34_v0);
            }
          }
          return _coll6;
        }()).Elements) {
          let _35_v1 = _compr_5;
          if ((function () {
            let _coll7 = new _dafny.Set();
            for (const _compr_7 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("pisbhaerb"), _dafny.Seq.UnicodeFromString("ejfkrj"), _dafny.Seq.UnicodeFromString("qqubr"))).Elements) {
              let _36_v0 = _compr_7;
              if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("pisbhaerb"), _dafny.Seq.UnicodeFromString("ejfkrj"), _dafny.Seq.UnicodeFromString("qqubr")), _36_v0)) {
                _coll7.add(_36_v0);
              }
            }
            return _coll7;
          }()).contains(_35_v1)) {
            _coll5.add(_35_v1);
          }
        }
        return _coll5;
      }()).contains(_33_v2)) {
        _coll1.add(_33_v2);
      }
    }
    return _coll1;
  }()).Elements) {
    let _37_v3 = _compr_0;
    if ((function () {
      let _coll8 = new _dafny.Set();
      for (const _compr_8 of (function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of (function () {
          let _coll10 = new _dafny.Set();
          for (const _compr_10 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("pisbhaerb"), _dafny.Seq.UnicodeFromString("ejfkrj"), _dafny.Seq.UnicodeFromString("qqubr"))).Elements) {
            let _38_v0 = _compr_10;
            if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("pisbhaerb"), _dafny.Seq.UnicodeFromString("ejfkrj"), _dafny.Seq.UnicodeFromString("qqubr")), _38_v0)) {
              _coll10.add(_38_v0);
            }
          }
          return _coll10;
        }()).Elements) {
          let _39_v1 = _compr_9;
          if ((function () {
            let _coll11 = new _dafny.Set();
            for (const _compr_11 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("pisbhaerb"), _dafny.Seq.UnicodeFromString("ejfkrj"), _dafny.Seq.UnicodeFromString("qqubr"))).Elements) {
              let _40_v0 = _compr_11;
              if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("pisbhaerb"), _dafny.Seq.UnicodeFromString("ejfkrj"), _dafny.Seq.UnicodeFromString("qqubr")), _40_v0)) {
                _coll11.add(_40_v0);
              }
            }
            return _coll11;
          }()).contains(_39_v1)) {
            _coll9.add(_39_v1);
          }
        }
        return _coll9;
      }()).Elements) {
        let _41_v2 = _compr_8;
        if ((function () {
          let _coll12 = new _dafny.Set();
          for (const _compr_12 of (function () {
            let _coll13 = new _dafny.Set();
            for (const _compr_13 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("pisbhaerb"), _dafny.Seq.UnicodeFromString("ejfkrj"), _dafny.Seq.UnicodeFromString("qqubr"))).Elements) {
              let _42_v0 = _compr_13;
              if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("pisbhaerb"), _dafny.Seq.UnicodeFromString("ejfkrj"), _dafny.Seq.UnicodeFromString("qqubr")), _42_v0)) {
                _coll13.add(_42_v0);
              }
            }
            return _coll13;
          }()).Elements) {
            let _43_v1 = _compr_12;
            if ((function () {
              let _coll14 = new _dafny.Set();
              for (const _compr_14 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("pisbhaerb"), _dafny.Seq.UnicodeFromString("ejfkrj"), _dafny.Seq.UnicodeFromString("qqubr"))).Elements) {
                let _44_v0 = _compr_14;
                if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("pisbhaerb"), _dafny.Seq.UnicodeFromString("ejfkrj"), _dafny.Seq.UnicodeFromString("qqubr")), _44_v0)) {
                  _coll14.add(_44_v0);
                }
              }
              return _coll14;
            }()).contains(_43_v1)) {
              _coll12.add(_43_v1);
            }
          }
          return _coll12;
        }()).contains(_41_v2)) {
          _coll8.add(_41_v2);
        }
      }
      return _coll8;
    }()).contains(_37_v3)) {
      _coll0.add(_37_v3);
    }
  }
  return _coll0;
}()).length));
      } else if (_source1.is_DC28) {
        let _45___mcc_h8 = (_source1).cf53;
        let _46___mcc_h9 = (_source1).cf54;
        let _47_cf54 = _46___mcc_h9;
        let _48_cf53 = _45___mcc_h8;
        return _module.D0.create_DC1(_48_cf53, _48_cf53, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_47_cf54,_48_cf53)).length));
      } else if (_source1.is_DC25) {
        let _49___mcc_h10 = (_source1).cf44;
        let _50_cf44 = _49___mcc_h10;
        if (false) {
          return _module.D0.create_DC1(false, false, new BigNumber(770));
        } else {
          return _module.D0.create_DC1(true, false, new BigNumber((_dafny.Seq.UnicodeFromString("dbqyvcng")).length));
        }
      } else {
        let _51___mcc_h11 = (_source1).cf55;
        let _52_cf55 = _51___mcc_h11;
        return _module.D0.create_DC1(false, true, new BigNumber(859));
      }
    };
    static fm17(p0, p1, globalState) {
      return _module.D0.create_DC1(true, (_dafny.MultiSet.fromElements(new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)))).IsSubsetOf(_dafny.MultiSet.fromElements(new _dafny.CodePoint('i'.codePointAt(0)))), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(125)), _dafny.Seq.of(new BigNumber(294)))).length));
    };
    static fm18(p0, p1, globalState) {
      return (_module.D18.create_DC55((_module.D18.create_DC55(_dafny.Seq.UnicodeFromString("qjvd"), true)).dtor_cf104, true)).dtor_cf104;
    };
    static fm19(p0, p1, globalState) {
      let _source2 = _module.D13.create_DC43(_module.D13.create_DC40(_dafny.MultiSet.FromArray(_dafny.Seq.of(false, true, false))));
      if (_source2.is_DC41) {
        let _53___mcc_h0 = (_source2).cf76;
        let _54_cf76 = _53___mcc_h0;
        return _module.D4.create_DC11(_dafny.Seq.of(false, true, false, false));
      } else if (_source2.is_DC42) {
        let _55___mcc_h1 = (_source2).cf77;
        let _56___mcc_h2 = (_source2).cf78;
        let _57___mcc_h3 = (_source2).cf79;
        let _58___mcc_h4 = (_source2).cf80;
        let _59___mcc_h5 = (_source2).cf81;
        let _60_cf81 = _59___mcc_h5;
        let _61_cf80 = _58___mcc_h4;
        let _62_cf79 = _57___mcc_h3;
        let _63_cf78 = _56___mcc_h2;
        let _64_cf77 = _55___mcc_h1;
        return _module.D4.create_DC11(_64_cf77);
      } else if (_source2.is_DC40) {
        let _65___mcc_h6 = (_source2).cf75;
        let _66_cf75 = _65___mcc_h6;
        return _module.D4.create_DC11(_dafny.Seq.of(true, true));
      } else {
        let _67___mcc_h7 = (_source2).cf82;
        let _68_cf82 = _67___mcc_h7;
        return _module.D4.create_DC11(_dafny.Seq.of(false));
      }
    };
    static fm20(p0, globalState) {
      return _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qcx"), _dafny.Seq.UnicodeFromString("yykt")), _dafny.Seq.UnicodeFromString("maxyot"));
    };
    static fm21(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("wcathxy"), _dafny.Seq.UnicodeFromString("ok"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(626)), function (_69_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("epmmm"), _dafny.Seq.UnicodeFromString("lvjl")), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-381)), function (_70_i1) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(927)), function (_71_i2) {
          return new _dafny.CodePoint('b'.codePointAt(0));
        });
      }));
    };
    static fm22(p0, globalState) {
      return new _dafny.CodePoint('x'.codePointAt(0));
    };
    static fm23(globalState) {
      return _module.D4.create_DC12();
    };
    static fm24(p0, globalState) {
      return _module.D2.create_DC7(new BigNumber(46), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(940)), function (_72_i0) {
  return new BigNumber(415);
})).length), (_module.D2.create_DC7(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(69),true)).length), new BigNumber(258), _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(true, false)).cardinality()), new BigNumber(914))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(590)), function (_73_i1) {
  return new _dafny.CodePoint('v'.codePointAt(0));
})).length))).dtor_cf18, new BigNumber(683), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("umst"), _dafny.Seq.UnicodeFromString("h"))).length));
    };
    static fm25(p0, globalState) {
      let _source3 = _module.D1.create_DC4(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(122),!(true)), true, new BigNumber((_dafny.Seq.of(new BigNumber(88), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.MultiSet.fromElements(true, true, false)).cardinality()))).length))).length), true);
      if (_source3.is_DC4) {
        let _74___mcc_h0 = (_source3).cf9;
        let _75___mcc_h1 = (_source3).cf10;
        let _76___mcc_h2 = (_source3).cf11;
        let _77___mcc_h3 = (_source3).cf12;
        let _78_cf12 = _77___mcc_h3;
        let _79_cf11 = _76___mcc_h2;
        let _80_cf10 = _75___mcc_h1;
        let _81_cf9 = _74___mcc_h0;
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_78_cf12, _78_cf12)).length),_79_cf11)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(851),_79_cf11));
      } else {
        let _82___mcc_h4 = (_source3).cf8;
        let _83_cf8 = _82___mcc_h4;
        return (function () {
          let _coll15 = new _dafny.Map();
          for (const _compr_15 of _dafny.IntegerRange(new BigNumber(392), new BigNumber(89))) {
            let _84_v0 = _compr_15;
            if (((new BigNumber(392)).isLessThanOrEqualTo(_84_v0)) && ((_84_v0).isLessThan(new BigNumber(89)))) {
              _coll15.push([_module.__default.safeModuloInt(_84_v0, new BigNumber((_dafny.Seq.UnicodeFromString("yy")).length)),new BigNumber(803)]);
            }
          }
          return _coll15;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())),(_module.D5.create_DC17(new BigNumber(258))).dtor_cf35));
      }
    };
    static fm26(p0, p1, p2, p3, globalState) {
      let _source4 = _module.D6.create_DC20(!(true), (_module.D8.create_DC27(true, true, new BigNumber((_dafny.Set.fromElements(false)).length), new BigNumber((_dafny.Seq.UnicodeFromString("u")).length))).dtor_cf49);
      if (_source4.is_DC20) {
        let _85___mcc_h0 = (_source4).cf38;
        let _86___mcc_h1 = (_source4).cf39;
        let _87_cf39 = _86___mcc_h1;
        let _88_cf38 = _85___mcc_h0;
        return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-70), new BigNumber(693), (_dafny.ZERO).minus(new BigNumber(-986)))).length));
      } else if (_source4.is_DC19) {
        let _89___mcc_h2 = (_source4).cf37;
        let _90_cf37 = _89___mcc_h2;
        return _dafny.MultiSet.fromElements(new BigNumber(-916), new BigNumber(212));
      } else {
        let _91___mcc_h3 = (_source4).cf40;
        let _92_cf40 = _91___mcc_h3;
        return _dafny.MultiSet.fromElements(new BigNumber(772));
      }
    };
    static fm27(p0, p1, globalState) {
      let _source5 = _module.D18.create_DC53(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("ujbhhlxdv")));
      if (_source5.is_DC54) {
        let _93___mcc_h0 = (_source5).cf100;
        let _94___mcc_h1 = (_source5).cf101;
        let _95___mcc_h2 = (_source5).cf102;
        let _96___mcc_h3 = (_source5).cf103;
        let _97_cf103 = _96___mcc_h3;
        let _98_cf102 = _95___mcc_h2;
        let _99_cf101 = _94___mcc_h1;
        let _100_cf100 = _93___mcc_h0;
        return _module.D5.create_DC17(_98_cf102);
      } else if (_source5.is_DC55) {
        let _101___mcc_h4 = (_source5).cf104;
        let _102___mcc_h5 = (_source5).cf105;
        let _103_cf105 = _102___mcc_h5;
        let _104_cf104 = _101___mcc_h4;
        return _module.D5.create_DC17(new BigNumber((_dafny.Seq.of(_103_cf105, _103_cf105, _103_cf105, _103_cf105, _103_cf105)).length));
      } else {
        let _105___mcc_h6 = (_source5).cf99;
        let _106_cf99 = _105___mcc_h6;
        return _module.D5.create_DC17(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(true, true)).length), new BigNumber(-18), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(75))).length),!(false))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(!(true), false)).length)))).length));
      }
    };
    static fm28(p0, p1, globalState) {
      return _module.D8.create_DC26(true, !((_dafny.MultiSet.fromElements(new BigNumber(449))).IsDisjointFrom(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-443))).length)))), (_dafny.Set.fromElements(true)).IsProperSubsetOf(_dafny.Set.fromElements(!(false), false)), new BigNumber(285));
    };
    static fm29(globalState) {
      return _dafny.Set.fromElements(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("b")).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false)).length), new BigNumber((_dafny.Seq.UnicodeFromString("whjhxs")).length), new BigNumber(378))).cardinality())), new BigNumber(((_dafny.Set.fromElements(!(false))).Difference(_dafny.Set.fromElements(false, false, true))).length), (new BigNumber((_dafny.Set.fromElements(false)).length)).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("bcax")).length),function () {
        let _coll16 = new _dafny.Map();
        for (const _compr_16 of _dafny.IntegerRange(new BigNumber(-957), new BigNumber(399))) {
          let _107_v0 = _compr_16;
          if (((new BigNumber(-957)).isLessThanOrEqualTo(_107_v0)) && ((_107_v0).isLessThan(new BigNumber(399)))) {
            _coll16.push([_module.__default.safeDivisionInt(_107_v0, new BigNumber((_dafny.Seq.of(true)).length)),false]);
          }
        }
        return _coll16;
      }())).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_module.D4.create_DC13(false, _dafny.MultiSet.fromElements(new _dafny.CodePoint('q'.codePointAt(0))), false)).dtor_cf31,new BigNumber(-621))).length), (new BigNumber((_dafny.Seq.UnicodeFromString("xb")).length)).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length))));
    };
    static fm30(p0, p1, p2, p3, globalState) {
      return _module.D4.create_DC13((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(853)), function (_108_i0) {
  return new _dafny.CodePoint('r'.codePointAt(0));
})).length)).isLessThanOrEqualTo(new BigNumber(-795)), _dafny.MultiSet.fromElements(new _dafny.CodePoint('k'.codePointAt(0))), _dafny.areEqual(_dafny.Seq.UnicodeFromString("ljkdpw"), _dafny.Seq.UnicodeFromString("hrxq")));
    };
    static fm31(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(new BigNumber(612)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(35)), function (_109_i0) {
        return new BigNumber(207);
      }))).length))))), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("k")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),true)).length), new BigNumber(-416))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(978), new BigNumber((_dafny.Seq.UnicodeFromString("cltxgkd")).length)), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(648), new BigNumber(-694), new BigNumber(372), new BigNumber(280))).length))));
    };
    static fm32(p0, p1, p2, globalState) {
      return _dafny.Seq.of(true, ((false) ? (false) : (true)), true, true);
    };
    static fm33(p0, p1, globalState) {
      return _module.D8.create_DC29(_module.D8.create_DC28((_module.D19.create_DC58(false)).dtor_cf107, new BigNumber((_dafny.Seq.of(_dafny.Set.fromElements(false), _dafny.Set.fromElements(!(false)))).length)));
    };
    static fm34(p0, p1, globalState) {
      if (!(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("s")).length)))).contains(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("xeh")).length)), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_module.D10.create_DC33(false), _module.D10.create_DC33(true), _module.D10.create_DC33(!(true)))).length))).cardinality()))).length))) {
        return (_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(true, !(false), false)).length))).cardinality()))))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(987)), function (_110_i0) {
          return new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(251), new BigNumber((_dafny.Seq.of(!(false))).length))).cardinality()), new BigNumber((_dafny.Set.fromElements(!(false))).length))).cardinality());
        }), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(422))).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(677)), function (_111_i1) {
          return new BigNumber(691);
        }), _dafny.Seq.of(new BigNumber(672)), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('u'.codePointAt(0)),_module.D1.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(54)), function (_112_i2) {
  return _module.D0.create_DC1(!(!(false)), true, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_module.D4.create_DC12())).length))).cardinality()));
})))).length),new BigNumber(911))).length)))));
      } else {
        return (_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.of(new BigNumber((function () {
          let _coll17 = new _dafny.Map();
          for (const _compr_17 of _dafny.IntegerRange(new BigNumber(389), new BigNumber(925))) {
            let _113_v0 = _compr_17;
            if (((new BigNumber(389)).isLessThanOrEqualTo(_113_v0)) && ((_113_v0).isLessThan(new BigNumber(925)))) {
              _coll17.push([_module.__default.safeModuloInt(_113_v0, new BigNumber((_dafny.Set.fromElements(new BigNumber(-405), new BigNumber(630), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
                let _coll18 = new _dafny.Map();
                for (const _compr_18 of _dafny.IntegerRange(new BigNumber(-370), new BigNumber(149))) {
                  let _114_v1 = _compr_18;
                  if (((new BigNumber(-370)).isLessThanOrEqualTo(_114_v1)) && ((_114_v1).isLessThan(new BigNumber(149)))) {
                    _coll18.push([(_114_v1).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("jegrswki")).length))).length)),false]);
                  }
                }
                return _coll18;
              }()).length),false)).length), new BigNumber((_dafny.Seq.of(_dafny.Seq.of(false))).length))).length)),new _dafny.CodePoint('p'.codePointAt(0))]);
            }
          }
          return _coll17;
        }()).length))))).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(178), new BigNumber(522), new BigNumber(647), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-523),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(true)).length))).length))).length))).length)), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false)).length), new BigNumber(426)), _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),(_dafny.ZERO).minus(new BigNumber(-36)))).length))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(816)), function (_115_i3) {
          return new BigNumber(321);
        }), _dafny.Seq.of(new BigNumber(-821))));
      }
    };
    static fm35(p0, p1, globalState) {
      if (false) {
        if (true) {
          return _module.D8.create_DC25(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(554)), function (_116_i0) {
  return new _dafny.CodePoint('m'.codePointAt(0));
}),new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())));
        } else {
          return _module.D8.create_DC25(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("oth"),new BigNumber(247)));
        }
      } else {
        return _module.D8.create_DC25(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("bg"),(_dafny.ZERO).minus(new BigNumber(-997))));
      }
    };
    static fm36(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(!(false)),false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(false),true));
    };
    static fm37(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(false, (new BigNumber(923)).isLessThanOrEqualTo(new BigNumber(-539)), true, !(false));
    };
    static fm38(p0, p1, p2, p3, globalState) {
      return _module.D3.create_DC10(true, new _dafny.CodePoint('v'.codePointAt(0)), false, true, false);
    };
    static fm39(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements(new BigNumber(587), new BigNumber((_dafny.Set.fromElements(true, !(false))).length), new BigNumber(103))).length));
    };
    static fm40(globalState) {
      if (_dafny.areEqual(_dafny.Seq.UnicodeFromString("udj"), _dafny.Seq.UnicodeFromString("jllaki"))) {
        return _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-482),new BigNumber(297)));
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.of(function () {
          let _coll19 = new _dafny.Map();
          for (const _compr_19 of _dafny.IntegerRange(new BigNumber(836), new BigNumber(538))) {
            let _117_v0 = _compr_19;
            if (((new BigNumber(836)).isLessThanOrEqualTo(_117_v0)) && ((_117_v0).isLessThan(new BigNumber(538)))) {
              _coll19.push([(_117_v0).plus(new BigNumber((_dafny.Set.fromElements(false, false, true)).length)),new BigNumber((_dafny.Seq.of(!(false), false)).length)]);
            }
          }
          return _coll19;
        }()), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_module.D16.create_DC51(false, _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.MultiSet.fromElements(false, false)).cardinality()))), true, new BigNumber((_dafny.Seq.of(false)).length))).dtor_cf97,new BigNumber((_dafny.Seq.of(new BigNumber(-228))).length))));
      }
    };
    static fm41(p0, p1, p2, globalState) {
      let _source6 = _module.D2.create_DC8(_module.D2.create_DC7(new BigNumber((_dafny.Seq.of(!(true), false)).length), new BigNumber(475), _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber(941), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(197))).cardinality()))), new BigNumber(189), new BigNumber((_dafny.MultiSet.fromElements(false, true)).cardinality())));
      if (_source6.is_DC6) {
        let _118___mcc_h0 = (_source6).cf14;
        let _119___mcc_h1 = (_source6).cf15;
        let _120_cf15 = _119___mcc_h1;
        let _121_cf14 = _118___mcc_h0;
        return _module.D5.create_DC15(_module.D4.create_DC11(_dafny.Seq.of(true)), true);
      } else if (_source6.is_DC7) {
        let _122___mcc_h2 = (_source6).cf16;
        let _123___mcc_h3 = (_source6).cf17;
        let _124___mcc_h4 = (_source6).cf18;
        let _125___mcc_h5 = (_source6).cf19;
        let _126___mcc_h6 = (_source6).cf20;
        let _127_cf20 = _126___mcc_h6;
        let _128_cf19 = _125___mcc_h5;
        let _129_cf18 = _124___mcc_h4;
        let _130_cf17 = _123___mcc_h3;
        let _131_cf16 = _122___mcc_h2;
        return _module.D5.create_DC15(_module.D4.create_DC11(_dafny.Seq.of(false, true)), !(false));
      } else if (_source6.is_DC5) {
        let _132___mcc_h7 = (_source6).cf13;
        let _133_cf13 = _132___mcc_h7;
        return _module.D5.create_DC15(_module.D4.create_DC11(_dafny.Seq.of(false, true)), false);
      } else {
        let _134___mcc_h8 = (_source6).cf21;
        let _135_cf21 = _134___mcc_h8;
        return _module.D5.create_DC15(_module.D4.create_DC11(_dafny.Seq.of(false)), true);
      }
    };
    static fm42(p0, p1, p2, globalState) {
      return _dafny.Seq.of((_module.D20.create_DC60(_dafny.Set.fromElements(new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0))))).dtor_cf109, _dafny.Set.fromElements(new _dafny.CodePoint('s'.codePointAt(0))), (_dafny.Set.fromElements(new _dafny.CodePoint('a'.codePointAt(0)))).Union(_dafny.Set.fromElements(new _dafny.CodePoint('q'.codePointAt(0)))), function () {
        let _coll20 = new _dafny.Set();
        for (const _compr_20 of (_dafny.Set.fromElements(new _dafny.CodePoint('r'.codePointAt(0)))).Elements) {
          let _136_v0 = _compr_20;
          if ((_dafny.Set.fromElements(new _dafny.CodePoint('r'.codePointAt(0)))).contains(_136_v0)) {
            _coll20.add(_136_v0);
          }
        }
        return _coll20;
      }());
    };
    static m6(globalState) {
      let r0 = false;
      let r1 = _dafny.Map.Empty;
      let _137_v0;
      _137_v0 = new BigNumber(387);
      let _138_v1;
      _138_v1 = true;
      let _139_v2;
      _139_v2 = _dafny.Map.Empty.slice().updateUnsafe(_138_v1,_137_v0);
      let _140_v3;
      let _nw0 = Array((_dafny.ONE).toNumber());
      _nw0[(_dafny.ZERO).toNumber()] = _139_v2;
      _140_v3 = _nw0;
      let _141_v4;
      let _init0 = ((_142_v1) => function (_143_i1) {
        return _142_v1;
      })(_138_v1);
      let _nw1 = Array((new BigNumber(4)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
        _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _141_v4 = _nw1;
      let _144_v5;
      let _nw2 = new _module.C2();
      _nw2.__ctor(true, _137_v0, _140_v3, _138_v1, _141_v4);
      _144_v5 = _nw2;
      let _145_v6;
      _145_v6 = _dafny.MultiSet.fromElements(_144_v5, _144_v5, _144_v5, _144_v5);
      let _146_v7;
      _146_v7 = _module.D7.create_DC23(_137_v0);
      let _hi0 = ((_146_v7).dtor_cf42).minus(_137_v0);
      for (let _147_i0 = new BigNumber((_145_v6).cardinality()); _147_i0.isLessThan(_hi0); _147_i0 = _147_i0.plus(_dafny.ONE)) {
        let _148_v8;
        let _init1 = function (_149_i2) {
          return _module.__default.safeModuloInt(_149_i2, new BigNumber((_dafny.Seq.UnicodeFromString("ussfu")).length));
        };
        let _nw3 = Array((new BigNumber(27)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw3.length); _i0_1++) {
          _nw3[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _148_v8 = _nw3;
        _148_v8 = _148_v8;
        _138_v1 = _138_v1;
        let _index0 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length));
        (_148_v8)[_index0] = new BigNumber(-482);
        let _index1 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length));
        (_148_v8)[_index1] = _module.__default.safeModuloInt(_147_i0, _137_v0);
        let _150_v9;
        _150_v9 = _dafny.MultiSet.fromElements((_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))], new BigNumber(887));
        let _151_v10;
        _151_v10 = _dafny.MultiSet.fromElements(new BigNumber((_150_v9).cardinality()), _147_i0);
        let _152_v11;
        _152_v11 = _dafny.MultiSet.fromElements(_151_v10, _151_v10);
        let _153_v12;
        _153_v12 = _module.D2.create_DC7((_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))], (_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))], _152_v11, (_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))], new BigNumber(-258));
        let _154_v13;
        _154_v13 = _module.D2.create_DC8(_153_v12);
        let _source7 = ((_138_v1) ? (((_138_v1) ? (_154_v13) : (_154_v13))) : (_154_v13));
        if (_source7.is_DC6) {
          let _155___mcc_h0 = (_source7).cf14;
          let _156___mcc_h1 = (_source7).cf15;
          let _157_cf15 = _156___mcc_h1;
          let _158_cf14 = _155___mcc_h0;
          _148_v8 = ((_138_v1) ? (_148_v8) : (_148_v8));
          _157_cf15 = (_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))];
          _137_v0 = _module.__default.safeDivisionInt((_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))], _137_v0);
          let _159_v14;
          let _nw4 = Array((new BigNumber(3)).toNumber());
          _159_v14 = _nw4;
          let _index2 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_159_v14).length));
          (_159_v14)[_index2] = _144_v5;
          let _index3 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_159_v14).length));
          (_159_v14)[_index3] = _144_v5;
        } else if (_source7.is_DC7) {
          let _160___mcc_h2 = (_source7).cf16;
          let _161___mcc_h3 = (_source7).cf17;
          let _162___mcc_h4 = (_source7).cf18;
          let _163___mcc_h5 = (_source7).cf19;
          let _164___mcc_h6 = (_source7).cf20;
          let _165_cf20 = _164___mcc_h6;
          let _166_cf19 = _163___mcc_h5;
          let _167_cf18 = _162___mcc_h4;
          let _168_cf17 = _161___mcc_h3;
          let _169_cf16 = _160___mcc_h2;
          let _170_v15;
          _170_v15 = new _dafny.CodePoint('i'.codePointAt(0));
          _170_v15 = _170_v15;
          let _171_v16;
          _171_v16 = _dafny.Seq.of(_138_v1);
          let _172_v17;
          _172_v17 = _dafny.Seq.of(_137_v0, _137_v0);
          _166_cf19 = _module.__default.safeDivisionInt(new BigNumber(636), ((_138_v1) ? (_137_v0) : (new BigNumber((_module.__default.fm26(_138_v1, _138_v1, _module.__default.fm41(_170_v15, _171_v16, _172_v17, globalState), _166_cf19, globalState)).cardinality()))));
          _168_cf17 = ((_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))]).multipliedBy(_168_cf17);
          let _173_v18;
          let _nw5 = new _module.C0();
          _nw5.__ctor();
          _173_v18 = _nw5;
        } else if (_source7.is_DC5) {
          let _174___mcc_h7 = (_source7).cf13;
          let _175_cf13 = _174___mcc_h7;
          let _176_v19;
          _176_v19 = _dafny.Map.Empty.slice().updateUnsafe((_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))],_138_v1);
          let _177_v20;
          _177_v20 = _module.D1.create_DC4(_176_v19, !(_138_v1), (_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))], _138_v1);
          let _index4 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length));
          (_148_v8)[_index4] = (_177_v20).dtor_cf11;
          let _178_v21;
          _178_v21 = _module.D6.create_DC20(_138_v1, _138_v1);
          let _179_v22;
          _179_v22 = _dafny.Set.fromElements(_178_v21, _178_v21);
          let _180_v23;
          _180_v23 = _dafny.Seq.of(_179_v22);
          let _181_v24;
          let _nw6 = new _module.C0();
          _nw6.__ctor();
          _181_v24 = _nw6;
          let _index5 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length));
          let _rhs0 = _180_v23;
          let _rhs1 = new BigNumber((((_139_v2).Merge(_139_v2)).Merge(_139_v2)).length);
          let _rhs2 = _148_v8;
          let _rhs3 = (_dafny.ZERO).minus((_137_v0).minus((_dafny.ZERO).minus((_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))])));
          let _rhs4 = _module.__default.safeModuloInt(((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_181_v24)).length))).plus(_147_i0), (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))], new BigNumber((_dafny.Set.fromElements(_147_i0, _137_v0)).length))));
          let _lhs0 = _148_v8;
          let _lhs1 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length));
          _180_v23 = _rhs0;
          _137_v0 = _rhs1;
          _148_v8 = _rhs2;
          _137_v0 = _rhs3;
          _lhs0[_lhs1] = _rhs4;
          let _182_v25;
          _182_v25 = _dafny.Set.fromElements((_175_cf13)[_module.__default.safeIndex(new BigNumber(-185), new BigNumber((_175_cf13).length))]);
          _138_v1 = (new BigNumber((_175_cf13).length)).isLessThan((new BigNumber((_182_v25).length)).minus(_137_v0));
          let _183_v26;
          _183_v26 = _dafny.MultiSet.fromElements(_138_v1, _138_v1, _138_v1, _138_v1);
          let _184_v27;
          _184_v27 = _module.D8.create_DC26(_138_v1, _dafny.Seq.IsPrefixOf(_175_cf13, _175_cf13), (new BigNumber((_183_v26).cardinality())).isLessThan((_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))]), (_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))]);
          _184_v27 = _184_v27;
        } else {
          let _185___mcc_h8 = (_source7).cf21;
          let _186_cf21 = _185___mcc_h8;
          let _187_v28;
          _187_v28 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.of((_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))]), _module.__default.safeIndex((_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))], new BigNumber((_dafny.Seq.of((_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))])).length)), _147_i0), _module.__default.safeIndex(_147_i0, new BigNumber((_dafny.Seq.update(_dafny.Seq.of((_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))]), _module.__default.safeIndex((_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))], new BigNumber((_dafny.Seq.of((_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))])).length)), _147_i0)).length)), (_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))])).length));
          _137_v0 = ((_module.__default.fm4((_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))], _187_v28, globalState)) ? ((_148_v8)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length))]) : ((_144_v5).fm2(globalState)));
          let _188_v29;
          _188_v29 = _dafny.Set.fromElements(false);
          let _index6 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_148_v8).length));
          (_148_v8)[_index6] = (_137_v0).plus(_module.__default.safeModuloInt(new BigNumber((_188_v29).length), new BigNumber((_dafny.Seq.of(_138_v1)).length)));
          let _189_v30;
          _189_v30 = _module.D8.create_DC28(_138_v1, new BigNumber(-658));
          let _190_v31;
          _190_v31 = _module.D8.create_DC29(_189_v30);
          _190_v31 = _190_v31;
          let _191_v32;
          _191_v32 = _dafny.Seq.UnicodeFromString("yvdvmxg");
          let _192_v33;
          let _nw7 = Array((new BigNumber(18)).toNumber());
          _nw7[(_dafny.ZERO).toNumber()] = _191_v32;
          _nw7[(_dafny.ONE).toNumber()] = _191_v32;
          _nw7[(new BigNumber(2)).toNumber()] = _191_v32;
          _nw7[(new BigNumber(3)).toNumber()] = _191_v32;
          _nw7[(new BigNumber(4)).toNumber()] = _191_v32;
          _nw7[(new BigNumber(5)).toNumber()] = _191_v32;
          _nw7[(new BigNumber(6)).toNumber()] = _191_v32;
          _nw7[(new BigNumber(7)).toNumber()] = _191_v32;
          _nw7[(new BigNumber(8)).toNumber()] = _191_v32;
          _nw7[(new BigNumber(9)).toNumber()] = _191_v32;
          _nw7[(new BigNumber(10)).toNumber()] = _191_v32;
          _nw7[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("eiaukerpe");
          _nw7[(new BigNumber(12)).toNumber()] = _191_v32;
          _nw7[(new BigNumber(13)).toNumber()] = _191_v32;
          _nw7[(new BigNumber(14)).toNumber()] = _191_v32;
          _nw7[(new BigNumber(15)).toNumber()] = _191_v32;
          _nw7[(new BigNumber(16)).toNumber()] = _191_v32;
          _nw7[(new BigNumber(17)).toNumber()] = _191_v32;
          _192_v33 = _nw7;
          let _193_v34;
          _193_v34 = _module.D19.create_DC56(_192_v33);
          let _pat_let_tv0 = _192_v33;
          let _194_v35;
          let _nw8 = Array((new BigNumber(2)).toNumber());
          _nw8[(_dafny.ZERO).toNumber()] = (function (_pat_let0_0) {
            return function (_195_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_196_dt__update_hcf106_h0) {
                  return _module.D19.create_DC56(_196_dt__update_hcf106_h0);
                }(_pat_let1_0);
              }(_pat_let_tv0);
            }(_pat_let0_0);
          }(_193_v34)).dtor_cf106;
          _nw8[(_dafny.ONE).toNumber()] = _192_v33;
          _194_v35 = _nw8;
          let _index7 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_194_v35).length));
          (_194_v35)[_index7] = _192_v33;
          let _index8 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_194_v35).length));
          let _init2 = ((_197_v32) => function (_198_i3) {
            return _197_v32;
          })(_191_v32);
          let _nw9 = Array((new BigNumber(28)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw9.length); _i0_2++) {
            _nw9[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          (_194_v35)[_index8] = _nw9;
        }
      }
      let _199_v36;
      _199_v36 = _dafny.MultiSet.fromElements(!(_138_v1), _138_v1);
      let _200_v37;
      _200_v37 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements(_138_v1)).equals(_199_v36),_138_v1);
      _200_v37 = (_200_v37).update(true, false);
      let _hi1 = _137_v0;
      for (let _201_i4 = _137_v0; _201_i4.isLessThan(_hi1); _201_i4 = _201_i4.plus(_dafny.ONE)) {
        if (_138_v1) {
          let _202_v38;
          let _nw10 = Array((new BigNumber(5)).toNumber());
          _nw10[(_dafny.ZERO).toNumber()] = _144_v5.f2;
          _nw10[(_dafny.ONE).toNumber()] = _144_v5.f2;
          _nw10[(new BigNumber(2)).toNumber()] = _141_v4;
          _nw10[(new BigNumber(3)).toNumber()] = _144_v5.f2;
          _nw10[(new BigNumber(4)).toNumber()] = _141_v4;
          _202_v38 = _nw10;
          let _203_v39;
          _203_v39 = _module.D15.create_DC47(_202_v38);
          let _204_v40;
          _204_v40 = _module.D15.create_DC49(_203_v39);
          let _205_v41;
          _205_v41 = _module.D15.create_DC49(_203_v39);
          let _206_v42;
          _206_v42 = _dafny.Seq.of(_205_v41);
          let _207_v43;
          let _nw11 = new _module.C4();
          _nw11.__ctor(!_dafny.areEqual(_206_v42, _206_v42));
          _207_v43 = _nw11;
          r0 = (_207_v43).f6;
          let _208_v44;
          _208_v44 = new _dafny.CodePoint('t'.codePointAt(0));
          let _209_v45;
          _209_v45 = _dafny.Seq.of(_208_v44);
          let _210_v46;
          _210_v46 = _module.D13.create_DC40(_dafny.MultiSet.fromElements((_207_v43).fm6(_209_v45, (_207_v43).f6, globalState), _138_v1));
          let _211_v47;
          _211_v47 = _dafny.Set.fromElements(_201_i4);
          r0 = !(_module.__default.fm4(new BigNumber(((_210_v46).dtor_cf75).cardinality()), _211_v47, globalState));
          (globalState).f1 = _144_v5.f2;
          let _212_v48;
          let _nw12 = new _module.C2();
          _nw12.__ctor((_207_v43).f6, _201_i4, _140_v3, _138_v1, _144_v5.f2);
          _212_v48 = _nw12;
          let _213_v49;
          let _nw13 = Array((new BigNumber(17)).toNumber());
          _nw13[(_dafny.ZERO).toNumber()] = _212_v48;
          _nw13[(_dafny.ONE).toNumber()] = _212_v48;
          _nw13[(new BigNumber(2)).toNumber()] = _212_v48;
          _nw13[(new BigNumber(3)).toNumber()] = _212_v48;
          _nw13[(new BigNumber(4)).toNumber()] = _212_v48;
          _nw13[(new BigNumber(5)).toNumber()] = _212_v48;
          _nw13[(new BigNumber(6)).toNumber()] = _212_v48;
          _nw13[(new BigNumber(7)).toNumber()] = _212_v48;
          _nw13[(new BigNumber(8)).toNumber()] = (((_207_v43).f6) ? (_212_v48) : (_212_v48));
          _nw13[(new BigNumber(9)).toNumber()] = _212_v48;
          _nw13[(new BigNumber(10)).toNumber()] = _212_v48;
          _nw13[(new BigNumber(11)).toNumber()] = _212_v48;
          _nw13[(new BigNumber(12)).toNumber()] = _212_v48;
          _nw13[(new BigNumber(13)).toNumber()] = _212_v48;
          _nw13[(new BigNumber(14)).toNumber()] = _212_v48;
          _nw13[(new BigNumber(15)).toNumber()] = _212_v48;
          _nw13[(new BigNumber(16)).toNumber()] = _212_v48;
          _213_v49 = _nw13;
          let _index9 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_213_v49).length));
          (_213_v49)[_index9] = _212_v48;
          let _214_v50;
          _214_v50 = _dafny.Set.fromElements(_138_v1);
          let _index10 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_213_v49).length));
          let _rhs5 = _201_i4;
          let _rhs6 = _212_v48;
          let _rhs7 = ((_212_v48.f7) ? ((true) || (true)) : ((_214_v50).IsDisjointFrom(_214_v50)));
          let _rhs8 = _dafny.Seq.Concat(_209_v45, _209_v45);
          let _lhs2 = _213_v49;
          let _lhs3 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_213_v49).length));
          _137_v0 = _rhs5;
          _lhs2[_lhs3] = _rhs6;
          _138_v1 = _rhs7;
          _209_v45 = _rhs8;
        } else {
          let _215_v51;
          let _nw14 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.of());
          _215_v51 = _nw14;
          let _216_v52;
          _216_v52 = new _dafny.CodePoint('p'.codePointAt(0));
          let _217_v53;
          _217_v53 = _dafny.Set.fromElements(_216_v52, _216_v52, _216_v52);
          let _218_v54;
          _218_v54 = _dafny.Seq.of(_217_v53);
          let _index11 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((_215_v51).length));
          (_215_v51)[_index11] = _218_v54;
          let _index12 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((_215_v51).length));
          (_215_v51)[_index12] = _dafny.Seq.Concat(_218_v54, _module.__default.fm42(_137_v0, _138_v1, _138_v1, globalState));
          _138_v1 = !(_138_v1);
          _216_v52 = _216_v52;
          _137_v0 = _137_v0;
          let _219_v55;
          _219_v55 = _dafny.Seq.UnicodeFromString("fc");
          r0 = (new BigNumber(920)).isLessThanOrEqualTo(_module.__default.safeModuloInt(new BigNumber((_219_v55).length), _137_v0));
        }
        let _220_v56;
        let _init3 = ((_221_i4, _222_v1) => function (_223_i5) {
          return _dafny.Seq.Concat(_dafny.Seq.of(_221_i4), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(_222_v1, _222_v1)).length)));
        })(_201_i4, _138_v1);
        let _nw15 = Array((new BigNumber(21)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw15.length); _i0_3++) {
          _nw15[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _220_v56 = _nw15;
        let _224_v57;
        _224_v57 = _dafny.Seq.of(_201_i4);
        let _index13 = _module.__default.safeIndex(new BigNumber(671), new BigNumber((_220_v56).length));
        (_220_v56)[_index13] = _224_v57;
        let _225_v58;
        let _nw16 = Array((new BigNumber(8)).toNumber());
        _225_v58 = _nw16;
        let _index14 = _module.__default.safeIndex(new BigNumber(671), new BigNumber((_220_v56).length));
        let _rhs9 = _138_v1;
        let _rhs10 = _224_v57;
        let _rhs11 = _225_v58;
        let _lhs4 = _220_v56;
        let _lhs5 = _module.__default.safeIndex(new BigNumber(671), new BigNumber((_220_v56).length));
        r0 = _rhs9;
        _lhs4[_lhs5] = _rhs10;
        _225_v58 = _rhs11;
        _138_v1 = _138_v1;
        let _226_v59;
        _226_v59 = _module.D14.create_DC45(_138_v1);
        _226_v59 = _module.D14.create_DC45(_138_v1);
      }
      let _227_v60;
      let _nw17 = Array((new BigNumber(16)).toNumber());
      _227_v60 = _nw17;
      let _228_v61;
      let _nw18 = new _module.C0();
      _nw18.__ctor();
      _228_v61 = _nw18;
      let _index15 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_227_v60).length));
      (_227_v60)[_index15] = _228_v61;
      let _229_v62;
      _229_v62 = _module.D8.create_DC27(_138_v1, _138_v1, new BigNumber((_dafny.Seq.UnicodeFromString("begh")).length), (_137_v0).minus(new BigNumber(416)));
      let _230_v63;
      let _init4 = function (_231_i6) {
        return (_231_i6).multipliedBy(new BigNumber(430));
      };
      let _nw19 = Array((new BigNumber(19)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw19.length); _i0_4++) {
        _nw19[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _230_v63 = _nw19;
      let _index16 = _module.__default.safeIndex(new BigNumber(503), new BigNumber((_230_v63).length));
      (_230_v63)[_index16] = _137_v0;
      let _232_v64;
      _232_v64 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(194),_141_v4);
      let _index17 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_227_v60).length));
      let _index18 = _module.__default.safeIndex(new BigNumber(503), new BigNumber((_230_v63).length));
      let _rhs12 = _228_v61;
      let _rhs13 = _229_v62;
      let _rhs14 = _module.__default.safeDivisionInt(new BigNumber(((_232_v64).update(new BigNumber((_module.__default.fm18(_138_v1, _dafny.Map.Empty.slice().updateUnsafe(_137_v0,true), globalState)).length), _144_v5.f2)).length), (_dafny.ZERO).minus(_137_v0));
      let _lhs6 = _227_v60;
      let _lhs7 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_227_v60).length));
      let _lhs8 = _230_v63;
      let _lhs9 = _module.__default.safeIndex(new BigNumber(503), new BigNumber((_230_v63).length));
      _lhs6[_lhs7] = _rhs12;
      _229_v62 = _rhs13;
      _lhs8[_lhs9] = _rhs14;
      let _233_v65;
      _233_v65 = new _dafny.CodePoint('c'.codePointAt(0));
      _233_v65 = _233_v65;
      let _index19 = _module.__default.safeIndex(new BigNumber(503), new BigNumber((_230_v63).length));
      (_230_v63)[_index19] = (_230_v63)[_module.__default.safeIndex(new BigNumber(503), new BigNumber((_230_v63).length))];
      let _234_v66;
      _234_v66 = _module.D3.create_DC10(_138_v1, _233_v65, _138_v1, _138_v1, _138_v1);
      let _235_v67;
      _235_v67 = _dafny.Map.Empty.slice().updateUnsafe(_234_v66,new BigNumber(324));
      let _pat_let_tv1 = _138_v1;
      let _pat_let_tv2 = _138_v1;
      r0 = (function () {
        let _coll21 = new _dafny.Set();
        for (const _compr_21 of (_235_v67).Keys.Elements) {
          let _236_v68 = _compr_21;
          if ((_235_v67).contains(_236_v68)) {
            _coll21.add(_236_v68);
          }
        }
        return _coll21;
      }()).contains(function (_pat_let2_0) {
        return function (_237_dt__update__tmp_h1) {
          return function (_pat_let3_0) {
            return function (_238_dt__update_hcf26_h0) {
              return function (_pat_let4_0) {
                return function (_239_dt__update_hcf23_h0) {
                  return _module.D3.create_DC10(_239_dt__update_hcf23_h0, (_237_dt__update__tmp_h1).dtor_cf24, (_237_dt__update__tmp_h1).dtor_cf25, _238_dt__update_hcf26_h0, (_237_dt__update__tmp_h1).dtor_cf27);
                }(_pat_let4_0);
              }(_pat_let_tv2);
            }(_pat_let3_0);
          }(_pat_let_tv1);
        }(_pat_let2_0);
      }(_234_v66));
      let _240_v70;
      _240_v70 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true),_139_v2);
      r1 = function () {
        let _coll22 = new _dafny.Map();
        for (const _compr_22 of (_240_v70).Keys.Elements) {
          let _241_v69 = _compr_22;
          if ((_240_v70).contains(_241_v69)) {
            _coll22.push([_241_v69,(_138_v1) || (true)]);
          }
        }
        return _coll22;
      }();
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _242_v0;
      _242_v0 = true;
      let _243_v1;
      _243_v1 = _dafny.Set.fromElements(_242_v0, _242_v0, _242_v0, _242_v0);
      let _244_v2;
      let _init5 = function (_245_i0) {
        return false;
      };
      let _nw20 = Array((new BigNumber(14)).toNumber());
      for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw20.length); _i0_5++) {
        _nw20[_i0_5] = _init5(new BigNumber(_i0_5));
      }
      _244_v2 = _nw20;
      let _246_globalState;
      let _nw21 = new _module.GlobalState();
      _nw21.__ctor(_243_v1, _244_v2);
      _246_globalState = _nw21;
      let _247_v3;
      _247_v3 = new BigNumber(923);
      let _hi2 = _247_v3;
      for (let _248_i1 = _247_v3; _248_i1.isLessThan(_hi2); _248_i1 = _248_i1.plus(_dafny.ONE)) {
        let _249_v4;
        _249_v4 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_247_v3),_248_i1);
        let _250_v5;
        _250_v5 = _dafny.Seq.of(_module.__default.fm0((_dafny.ZERO).minus(_247_v3), _246_globalState), _247_v3, _247_v3, (((_249_v4).contains(_247_v3)) ? ((_249_v4).get(_247_v3)) : (_247_v3)), _247_v3);
        _247_v3 = new BigNumber((_250_v5).length);
        let _251_v6;
        let _nw22 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
        _251_v6 = _nw22;
        let _252_v7;
        _252_v7 = _dafny.Map.Empty.slice().updateUnsafe(_251_v6,new BigNumber((_dafny.Seq.of(_248_i1)).length));
        _242_v0 = ((_dafny.ZERO).minus(_module.__default.safeModuloInt((((_249_v4).contains((_dafny.ZERO).minus(new BigNumber(-586)))) ? ((_249_v4).get((_dafny.ZERO).minus(new BigNumber(-586)))) : (_247_v3)), _248_i1))).isLessThan((_247_v3).multipliedBy(new BigNumber(((_252_v7).update(_251_v6, new BigNumber(464))).length)));
        let _253_v8;
        _253_v8 = _dafny.Seq.UnicodeFromString("rcrmu");
        if (_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("dbqp"), _253_v8)) {
          let _254_v9;
          let _nw23 = new _module.C3();
          _nw23.__ctor(_244_v2);
          _254_v9 = _nw23;
          let _255_v10;
          _255_v10 = _dafny.Seq.of(!(_242_v0), _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("vbki"), _253_v8), _242_v0);
          let _256_v11;
          _256_v11 = _dafny.Set.fromElements(_247_v3, _247_v3);
          let _257_v12;
          _257_v12 = _dafny.MultiSet.fromElements(_242_v0, _242_v0, _242_v0, _242_v0, false);
          let _258_v13;
          _258_v13 = _dafny.Map.Empty.slice().updateUnsafe(_247_v3,_module.__default.fm4(_248_i1, _256_v11, _246_globalState));
          let _259_v14;
          let _nw24 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
          _259_v14 = _nw24;
          let _260_v15;
          _260_v15 = _dafny.Seq.of(_244_v2, _244_v2);
          let _261_v16;
          let _nw25 = new _module.C2();
          _nw25.__ctor(!(!((_module.__default.fm29(_246_globalState)).IsProperSubsetOf(_256_v11))), (((_257_v12).contains(!(_242_v0))) ? ((_257_v12).get(!(_242_v0))) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_module.D1.create_DC4(_258_v13, _242_v0, new BigNumber((_256_v11).length), true)).dtor_cf11,_247_v3)).length))), _259_v14, (_256_v11).IsSubsetOf(_256_v11), (_260_v15)[_module.__default.safeIndex(_247_v3, new BigNumber((_260_v15).length))]);
          _261_v16 = _nw25;
          let _262_v17;
          _262_v17 = new _dafny.CodePoint('c'.codePointAt(0));
          let _rhs15 = (_dafny.ZERO).minus(_247_v3);
          let _rhs16 = _dafny.Seq.of(_dafny.areEqual(_250_v5, _250_v5), !(_261_v16.f7), (_module.__default.fm38(new BigNumber(301), _242_v0, _253_v8, _262_v17, _246_globalState)).dtor_cf23);
          let _rhs17 = (_250_v5)[_module.__default.safeIndex(_247_v3, new BigNumber((_250_v5).length))];
          let _rhs18 = _261_v16;
          _247_v3 = _rhs15;
          _255_v10 = _rhs16;
          _247_v3 = _rhs17;
          _261_v16 = _rhs18;
          let _263_v22;
          let _nw26 = Array((new BigNumber(13)).toNumber());
          _nw26[(_dafny.ZERO).toNumber()] = _258_v13;
          _nw26[(_dafny.ONE).toNumber()] = _258_v13;
          _nw26[(new BigNumber(2)).toNumber()] = _258_v13;
          _nw26[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_248_i1,_242_v0);
          _nw26[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_253_v8).length),(_255_v10)[_module.__default.safeIndex(_247_v3, new BigNumber((_255_v10).length))]);
          _nw26[(new BigNumber(5)).toNumber()] = _258_v13;
          _nw26[(new BigNumber(6)).toNumber()] = function () {
            let _coll23 = new _dafny.Map();
            for (const _compr_23 of _dafny.IntegerRange(new BigNumber(54), new BigNumber(-989))) {
              let _264_v18 = _compr_23;
              if (((new BigNumber(54)).isLessThanOrEqualTo(_264_v18)) && ((_264_v18).isLessThan(new BigNumber(-989)))) {
                _coll23.push([(_264_v18).plus(_247_v3),_261_v16.f7]);
              }
            }
            return _coll23;
          }();
          _nw26[(new BigNumber(7)).toNumber()] = _258_v13;
          _nw26[(new BigNumber(8)).toNumber()] = function () {
            let _coll24 = new _dafny.Map();
            for (const _compr_24 of _dafny.IntegerRange(new BigNumber(-130), new BigNumber(973))) {
              let _265_v19 = _compr_24;
              if (((new BigNumber(-130)).isLessThanOrEqualTo(_265_v19)) && ((_265_v19).isLessThan(new BigNumber(973)))) {
                _coll24.push([_module.__default.safeDivisionInt(_265_v19, _248_i1),false]);
              }
            }
            return _coll24;
          }();
          _nw26[(new BigNumber(9)).toNumber()] = (_258_v13).Merge(_258_v13);
          _nw26[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_248_i1,_261_v16.f7);
          _nw26[(new BigNumber(11)).toNumber()] = (_258_v13).Merge(function () {
            let _coll25 = new _dafny.Map();
            for (const _compr_25 of _dafny.IntegerRange(new BigNumber(-922), new BigNumber(-131))) {
              let _266_v20 = _compr_25;
              if (((new BigNumber(-922)).isLessThanOrEqualTo(_266_v20)) && ((_266_v20).isLessThan(new BigNumber(-131)))) {
                _coll25.push([(_266_v20).minus(new BigNumber(28)),_261_v16.f7]);
              }
            }
            return _coll25;
          }());
          _nw26[(new BigNumber(12)).toNumber()] = function () {
            let _coll26 = new _dafny.Map();
            for (const _compr_26 of _dafny.IntegerRange(new BigNumber(558), new BigNumber(765))) {
              let _267_v21 = _compr_26;
              if (((new BigNumber(558)).isLessThanOrEqualTo(_267_v21)) && ((_267_v21).isLessThan(new BigNumber(765)))) {
                _coll26.push([(_267_v21).plus(_248_i1),(_255_v10)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_250_v5)).cardinality()), new BigNumber((_255_v10).length))]]);
              }
            }
            return _coll26;
          }();
          _263_v22 = _nw26;
          let _268_v23;
          _268_v23 = _module.D1.create_DC4(_258_v13, _242_v0, _248_i1, _242_v0);
          let _index20 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_263_v22).length));
          (_263_v22)[_index20] = (_268_v23).dtor_cf9;
          let _269_v24;
          let _nw27 = Array((new BigNumber(9)).toNumber());
          _nw27[(_dafny.ZERO).toNumber()] = _253_v8;
          _nw27[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("oqt");
          _nw27[(new BigNumber(2)).toNumber()] = _253_v8;
          _nw27[(new BigNumber(3)).toNumber()] = _253_v8;
          _nw27[(new BigNumber(4)).toNumber()] = _253_v8;
          _nw27[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("cpvuli");
          _nw27[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_253_v8, _dafny.Seq.UnicodeFromString("v"));
          _nw27[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(521)), ((_270_v17) => function (_271_i2) {
            return _270_v17;
          })(_262_v17));
          _nw27[(new BigNumber(8)).toNumber()] = _253_v8;
          _269_v24 = _nw27;
          let _index21 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_269_v24).length));
          (_269_v24)[_index21] = _dafny.Seq.UnicodeFromString("fipalxnh");
          let _272_v26;
          _272_v26 = _module.D2.create_DC5(_dafny.Seq.of(_262_v17, _262_v17, _262_v17));
          let _index22 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_263_v22).length));
          let _index23 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_269_v24).length));
          let _rhs19 = function () {
            let _coll27 = new _dafny.Map();
            for (const _compr_27 of _dafny.IntegerRange(new BigNumber(65), new BigNumber(555))) {
              let _273_v25 = _compr_27;
              if (((new BigNumber(65)).isLessThanOrEqualTo(_273_v25)) && ((_273_v25).isLessThan(new BigNumber(555)))) {
                _coll27.push([(_273_v25).plus((_261_v16).f8),(((_258_v13).contains(_247_v3)) ? ((_258_v13).get(_247_v3)) : (_261_v16.f7))]);
              }
            }
            return _coll27;
          }();
          let _rhs20 = _254_v9;
          let _rhs21 = (_272_v26).dtor_cf13;
          let _rhs22 = (_module.__default.safeModuloInt(new BigNumber(616), new BigNumber((_250_v5).length))).plus((_261_v16).f8);
          let _lhs10 = _263_v22;
          let _lhs11 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_263_v22).length));
          let _lhs12 = _269_v24;
          let _lhs13 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_269_v24).length));
          _lhs10[_lhs11] = _rhs19;
          _254_v9 = _rhs20;
          _lhs12[_lhs13] = _rhs21;
          _247_v3 = _rhs22;
          let _274_v27;
          let _nw28 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
          _274_v27 = _nw28;
          let _275_v28;
          let _276_v29;
          let _out0;
          let _out1;
          let _outcollector0 = (_261_v16).m1(_247_v3, _262_v17, _269_v24, _274_v27, _246_globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _275_v28 = _out0;
          _276_v29 = _out1;
          _276_v29 = !(!(!(_276_v29) || (_module.__default.fm20(_242_v0, _246_globalState))));
        } else {
          _247_v3 = _module.__default.safeDivisionInt(new BigNumber(-2), _247_v3);
          let _277_v30;
          let _278_v31;
          let _out2;
          let _out3;
          let _outcollector1 = _module.__default.m6(_246_globalState);
          _out2 = _outcollector1[0];
          _out3 = _outcollector1[1];
          _277_v30 = _out2;
          _278_v31 = _out3;
          let _index24 = _module.__default.safeIndex(new BigNumber(454), new BigNumber((_244_v2).length));
          (_244_v2)[_index24] = _277_v30;
          let _279_v32;
          let _init6 = ((_280_v30, _281_i1) => function (_282_i3) {
            return _dafny.Map.Empty.slice().updateUnsafe(_280_v30,_281_i1);
          })(_277_v30, _248_i1);
          let _nw29 = Array((new BigNumber(2)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw29.length); _i0_6++) {
            _nw29[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _279_v32 = _nw29;
          let _283_v33;
          let _nw30 = new _module.C5();
          _nw30.__ctor(_277_v30, _244_v2, _279_v32, _242_v0);
          _283_v33 = _nw30;
          let _284_v34;
          _284_v34 = _dafny.Set.fromElements(_283_v33, _283_v33);
          let _index25 = _module.__default.safeIndex(new BigNumber(454), new BigNumber((_244_v2).length));
          (_244_v2)[_index25] = !(((_284_v34).Union(_284_v34)).IsSubsetOf((_284_v34).Difference(_284_v34)));
          let _285_v35;
          _285_v35 = _dafny.Seq.of((_283_v33).f5);
          _247_v3 = (_247_v3).minus(new BigNumber((_285_v35).length));
          let _286_v36;
          _286_v36 = _dafny.MultiSet.fromElements(_251_v6);
          let _287_v37;
          let _nw31 = Array((new BigNumber(13)).toNumber());
          _287_v37 = _nw31;
          let _288_v38;
          _288_v38 = _dafny.Map.Empty.slice().updateUnsafe(_287_v37,_247_v3);
          _247_v3 = new BigNumber(((((_286_v36).IsDisjointFrom(_286_v36)) ? ((_288_v38).update(_287_v37, _248_i1)) : (((_242_v0) ? (_288_v38) : (_288_v38))))).length);
        }
        let _289_v39;
        _289_v39 = _dafny.Map.Empty.slice().updateUnsafe(_244_v2,_249_v4);
        let _290_v40;
        let _init7 = ((_291_v0, _292_v3) => function (_293_i4) {
          return _dafny.Map.Empty.slice().updateUnsafe(_291_v0,_292_v3);
        })(_242_v0, _247_v3);
        let _nw32 = Array((new BigNumber(5)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw32.length); _i0_7++) {
          _nw32[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _290_v40 = _nw32;
        let _294_v41;
        let _nw33 = new _module.C1();
        _nw33.__ctor(_290_v40, _242_v0, _244_v2);
        _294_v41 = _nw33;
        let _295_v42;
        _295_v42 = _module.D10.create_DC31(_294_v41);
        let _rhs23 = _289_v39;
        let _rhs24 = (new BigNumber(669)).isLessThanOrEqualTo(_248_i1);
        let _rhs25 = (_295_v42).dtor_cf57;
        _289_v39 = _rhs23;
        _242_v0 = _rhs24;
        _294_v41 = _rhs25;
      }
      let _296_v43;
      let _297_v44;
      let _out4;
      let _out5;
      let _outcollector2 = _module.__default.m6(_246_globalState);
      _out4 = _outcollector2[0];
      _out5 = _outcollector2[1];
      _296_v43 = _out4;
      _297_v44 = _out5;
      let _298_v45;
      let _nw34 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
      _298_v45 = _nw34;
      let _299_v46;
      _299_v46 = _module.D5.create_DC14(_298_v45);
      let _source8 = _299_v46;
      if (_source8.is_DC15) {
        let _300___mcc_h0 = (_source8).cf33;
        let _301___mcc_h1 = (_source8).cf34;
        let _302_cf34 = _301___mcc_h1;
        let _303_cf33 = _300___mcc_h0;
        _243_v1 = _dafny.Set.fromElements(_302_cf34, _302_cf34, _242_v0);
        _247_v3 = _module.__default.fm0(_247_v3, _246_globalState);
        let _304_v47;
        let _nw35 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.of());
        _304_v47 = _nw35;
        let _305_v48;
        _305_v48 = _dafny.Seq.UnicodeFromString("u");
        let _306_v49;
        _306_v49 = _dafny.Seq.of(_305_v48);
        let _index26 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_304_v47).length));
        (_304_v47)[_index26] = _dafny.Seq.Concat(_306_v49, _dafny.Seq.Create(_module.__default.abs(new BigNumber(748)), ((_307_v48) => function (_308_i5) {
          return _307_v48;
        })(_305_v48)));
        let _309_v50;
        _309_v50 = new _dafny.CodePoint('e'.codePointAt(0));
        let _310_v51;
        let _nw36 = new _module.C3();
        _nw36.__ctor(_244_v2);
        _310_v51 = _nw36;
        let _311_v52;
        _311_v52 = _dafny.Map.Empty.slice().updateUnsafe(_305_v48,_310_v51);
        let _index27 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_304_v47).length));
        let _rhs26 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(561)), ((_312_v49, _313_v3, _314_v48) => function (_315_i6) {
          return _dafny.Seq.update(_dafny.Seq.Concat((_312_v49)[_module.__default.safeIndex(_313_v3, new BigNumber((_312_v49).length))], _314_v48), _module.__default.safeIndex(_313_v3, new BigNumber((_dafny.Seq.Concat((_312_v49)[_module.__default.safeIndex(_313_v3, new BigNumber((_312_v49).length))], _314_v48)).length)), new _dafny.CodePoint('q'.codePointAt(0)));
        })(_306_v49, _247_v3, _305_v48));
        let _rhs27 = _309_v50;
        let _rhs28 = _296_v43;
        let _rhs29 = _dafny.Map.Empty.slice().updateUnsafe(_305_v48,_310_v51);
        let _lhs14 = _304_v47;
        let _lhs15 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_304_v47).length));
        _lhs14[_lhs15] = _rhs26;
        _309_v50 = _rhs27;
        _296_v43 = _rhs28;
        _311_v52 = _rhs29;
        _247_v3 = new BigNumber((_dafny.Seq.Concat((_304_v47)[_module.__default.safeIndex(new BigNumber(52), new BigNumber((_304_v47).length))], _module.__default.fm21(_246_globalState))).length);
      } else if (_source8.is_DC16) {
        let _316_v53;
        _316_v53 = _dafny.MultiSet.fromElements(new BigNumber(729), _247_v3, _247_v3);
        let _317_v54;
        _317_v54 = _dafny.Seq.of(true, _296_v43, _296_v43, _242_v0);
        let _318_v55;
        _318_v55 = _module.D4.create_DC11(_317_v54);
        let _319_v56;
        _319_v56 = _dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_247_v3)).cardinality()), (_dafny.ZERO).minus(_247_v3));
        let _320_v57;
        _320_v57 = _dafny.Seq.UnicodeFromString("yfwppgog");
        let _321_v58;
        _321_v58 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm4(new BigNumber(243), _319_v56, _246_globalState),new BigNumber((_320_v57).length));
        let _322_v59;
        _322_v59 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-66),_242_v0);
        let _323_v60;
        _323_v60 = _dafny.Seq.of(_244_v2);
        let _324_v61;
        _324_v61 = _module.D1.create_DC4(_322_v59, _296_v43, new BigNumber((_323_v60).length), _296_v43);
        let _325_v62;
        let _nw37 = Array((new BigNumber(22)).toNumber());
        _nw37[(_dafny.ZERO).toNumber()] = (((_316_v53).contains(new BigNumber(-547))) ? ((_316_v53).get(new BigNumber(-547))) : (_247_v3));
        _nw37[(_dafny.ONE).toNumber()] = _module.__default.safeDivisionInt(_247_v3, new BigNumber(165));
        _nw37[(new BigNumber(2)).toNumber()] = (((_316_v53).contains(_247_v3)) ? ((_316_v53).get(_247_v3)) : (_247_v3));
        _nw37[(new BigNumber(3)).toNumber()] = _247_v3;
        _nw37[(new BigNumber(4)).toNumber()] = _247_v3;
        _nw37[(new BigNumber(5)).toNumber()] = _247_v3;
        _nw37[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt(_247_v3, _247_v3);
        _nw37[(new BigNumber(7)).toNumber()] = _247_v3;
        _nw37[(new BigNumber(8)).toNumber()] = _247_v3;
        _nw37[(new BigNumber(9)).toNumber()] = _247_v3;
        _nw37[(new BigNumber(10)).toNumber()] = new BigNumber(909);
        _nw37[(new BigNumber(11)).toNumber()] = _247_v3;
        _nw37[(new BigNumber(12)).toNumber()] = new BigNumber(((_318_v55).dtor_cf28).length);
        _nw37[(new BigNumber(13)).toNumber()] = (_module.__default.fm0(_247_v3, _246_globalState)).multipliedBy(_247_v3);
        _nw37[(new BigNumber(14)).toNumber()] = _247_v3;
        _nw37[(new BigNumber(15)).toNumber()] = _247_v3;
        _nw37[(new BigNumber(16)).toNumber()] = _247_v3;
        _nw37[(new BigNumber(17)).toNumber()] = new BigNumber((_dafny.Set.fromElements(_247_v3)).length);
        _nw37[(new BigNumber(18)).toNumber()] = _247_v3;
        _nw37[(new BigNumber(19)).toNumber()] = new BigNumber(((_321_v58).Merge(_321_v58)).length);
        _nw37[(new BigNumber(20)).toNumber()] = _247_v3;
        _nw37[(new BigNumber(21)).toNumber()] = (_324_v61).dtor_cf11;
        _325_v62 = _nw37;
        _325_v62 = _325_v62;
        let _326_v63;
        let _nw38 = new _module.C3();
        _nw38.__ctor(_244_v2);
        _326_v63 = _nw38;
        let _327_v64;
        _327_v64 = _dafny.Map.Empty.slice().updateUnsafe(_326_v63,new BigNumber(885));
        _327_v64 = (_327_v64).update(_326_v63, _247_v3);
        if (_242_v0) {
          let _index28 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_244_v2).length));
          (_244_v2)[_index28] = true;
          let _index29 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_244_v2).length));
          (_244_v2)[_index29] = _296_v43;
          let _328_v65;
          let _init8 = function (_329_i7) {
            return _dafny.Seq.UnicodeFromString("uyjkwf");
          };
          let _nw39 = Array((new BigNumber(3)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw39.length); _i0_8++) {
            _nw39[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _328_v65 = _nw39;
          _328_v65 = _328_v65;
          let _330_v66;
          let _init9 = ((_331_v0) => function (_332_i8) {
            return _331_v0;
          })(_242_v0);
          let _nw40 = Array((new BigNumber(3)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw40.length); _i0_9++) {
            _nw40[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _330_v66 = _nw40;
          let _333_v67;
          let _nw41 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Map.Empty);
          _333_v67 = _nw41;
          let _334_v68;
          let _nw42 = new _module.C5();
          _nw42.__ctor((_244_v2)[_module.__default.safeIndex(new BigNumber(800), new BigNumber((_244_v2).length))], _330_v66, _333_v67, (_243_v1).IsProperSubsetOf(_243_v1));
          _334_v68 = _nw42;
          let _335_v69;
          _335_v69 = _module.D8.create_DC26(false, (_334_v68).f5, _296_v43, _247_v3);
          let _336_v70;
          _336_v70 = _module.D8.create_DC29(_335_v69);
          _336_v70 = _336_v70;
          let _337_v71;
          _337_v71 = new _dafny.CodePoint('x'.codePointAt(0));
          let _338_v72;
          _338_v72 = _dafny.Map.Empty.slice().updateUnsafe(_337_v71,new _dafny.CodePoint('h'.codePointAt(0)));
          _296_v43 = _dafny.Seq.contains(_320_v57, (((_338_v72).contains(_337_v71)) ? ((_338_v72).get(_337_v71)) : (_337_v71)));
        } else {
          _247_v3 = (_247_v3).minus(new BigNumber((_320_v57).length));
          let _339_v73;
          _339_v73 = new _dafny.CodePoint('y'.codePointAt(0));
          _320_v57 = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("rasbc"), _module.__default.safeIndex(_247_v3, new BigNumber((_dafny.Seq.UnicodeFromString("rasbc")).length)), _339_v73);
          _296_v43 = _296_v43;
          let _340_v74;
          let _nw43 = new _module.C4();
          _nw43.__ctor((_243_v1).IsProperSubsetOf(_dafny.Set.fromElements(_296_v43)));
          _340_v74 = _nw43;
          let _341_v75;
          _341_v75 = _module.D6.create_DC20(_module.__default.fm20(_296_v43, _246_globalState), _242_v0);
          let _342_v76;
          let _343_v77;
          let _344_v78;
          let _out6;
          let _out7;
          let _out8;
          let _outcollector3 = (_340_v74).m3(_325_v62, (_341_v75).dtor_cf38, _246_globalState);
          _out6 = _outcollector3[0];
          _out7 = _outcollector3[1];
          _out8 = _outcollector3[2];
          _342_v76 = _out6;
          _343_v77 = _out7;
          _344_v78 = _out8;
        }
        _247_v3 = (_326_v63).fm2(_246_globalState);
      } else if (_source8.is_DC17) {
        let _345___mcc_h2 = (_source8).cf35;
        let _346_cf35 = _345___mcc_h2;
        let _347_v79;
        _347_v79 = new _dafny.CodePoint('d'.codePointAt(0));
        let _348_v80;
        _348_v80 = _dafny.Seq.of(_347_v79, _347_v79, _347_v79, _347_v79);
        let _source9 = _module.D2.create_DC5(_348_v80);
        if (_source9.is_DC6) {
          let _349___mcc_h5 = (_source9).cf14;
          let _350___mcc_h6 = (_source9).cf15;
          let _351_cf15 = _350___mcc_h6;
          let _352_cf14 = _349___mcc_h5;
          let _353_v81;
          let _nw44 = Array((new BigNumber(4)).toNumber()).fill([]);
          _353_v81 = _nw44;
          _353_v81 = ((_296_v43) ? (_353_v81) : (_353_v81));
          let _354_v82;
          let _nw45 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _354_v82 = _nw45;
          let _index30 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_354_v82).length));
          (_354_v82)[_index30] = _247_v3;
          let _index31 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_354_v82).length));
          (_354_v82)[_index31] = _247_v3;
          let _355_v83;
          let _356_v84;
          let _out9;
          let _out10;
          let _outcollector4 = _module.__default.m6(_246_globalState);
          _out9 = _outcollector4[0];
          _out10 = _outcollector4[1];
          _355_v83 = _out9;
          _356_v84 = _out10;
          let _357_v85;
          let _nw46 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
          _357_v85 = _nw46;
          let _358_v86;
          let _nw47 = Array((new BigNumber(27)).toNumber()).fill(_module.D12.Default());
          _358_v86 = _nw47;
          let _359_v87;
          _359_v87 = _dafny.Map.Empty.slice().updateUnsafe(_358_v86,_346_cf35);
          let _index32 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_357_v85).length));
          (_357_v85)[_index32] = _359_v87;
          let _360_v88;
          _360_v88 = _dafny.Map.Empty.slice().updateUnsafe((_354_v82)[_module.__default.safeIndex(new BigNumber(109), new BigNumber((_354_v82).length))],_242_v0);
          let _index33 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_357_v85).length));
          let _rhs30 = _dafny.Map.Empty.slice().updateUnsafe(_358_v86,_346_cf35);
          let _rhs31 = _module.__default.safeDivisionInt(_247_v3, new BigNumber(((_360_v88).Merge(_360_v88)).length));
          let _rhs32 = _242_v0;
          let _rhs33 = new BigNumber(982);
          let _lhs16 = _357_v85;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_357_v85).length));
          _lhs16[_lhs17] = _rhs30;
          _247_v3 = _rhs31;
          _242_v0 = _rhs32;
          _247_v3 = _rhs33;
        } else if (_source9.is_DC7) {
          let _361___mcc_h7 = (_source9).cf16;
          let _362___mcc_h8 = (_source9).cf17;
          let _363___mcc_h9 = (_source9).cf18;
          let _364___mcc_h10 = (_source9).cf19;
          let _365___mcc_h11 = (_source9).cf20;
          let _366_cf20 = _365___mcc_h11;
          let _367_cf19 = _364___mcc_h10;
          let _368_cf18 = _363___mcc_h9;
          let _369_cf17 = _362___mcc_h8;
          let _370_cf16 = _361___mcc_h7;
          let _371_v89;
          _371_v89 = _dafny.Map.Empty.slice().updateUnsafe(_367_cf19,_242_v0);
          _366_cf20 = ((new BigNumber((_348_v80).length)).plus(_346_cf35)).plus((_346_cf35).minus(new BigNumber((_371_v89).length)));
          let _372_v90;
          _372_v90 = _dafny.Seq.of(_366_cf20);
          let _373_v91;
          _373_v91 = _module.D8.create_DC27(_242_v0, _242_v0, new BigNumber((_dafny.Set.fromElements(_372_v90)).length), _247_v3);
          let _374_v92;
          let _init10 = ((_375_cf17) => function (_376_i9) {
            return _dafny.Map.Empty.slice().updateUnsafe(false,_375_cf17);
          })(_369_cf17);
          let _nw48 = Array((new BigNumber(22)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw48.length); _i0_10++) {
            _nw48[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _374_v92 = _nw48;
          let _377_v93;
          let _nw49 = new _module.C2();
          _nw49.__ctor((_373_v91).dtor_cf50, new BigNumber((_dafny.Seq.of(new BigNumber(272))).length), _374_v92, _242_v0, _244_v2);
          _377_v93 = _nw49;
          let _378_v94;
          _378_v94 = _dafny.Seq.of(_377_v93);
          let _379_v96;
          _379_v96 = _dafny.Map.Empty.slice().updateUnsafe(false,_367_cf19);
          let _380_v97;
          _380_v97 = _module.D12.create_DC38(function () {
  let _coll28 = new _dafny.Set();
  for (const _compr_28 of _dafny.IntegerRange(new BigNumber(180), new BigNumber(360))) {
    let _381_v95 = _compr_28;
    if (((new BigNumber(180)).isLessThanOrEqualTo(_381_v95)) && ((_381_v95).isLessThan(new BigNumber(360)))) {
      _coll28.add((_381_v95).minus(_367_cf19));
    }
  }
  return _coll28;
}(), _366_cf20, _377_v93, new BigNumber((_379_v96).length));
          let _382_v98;
          let _nw50 = Array((new BigNumber(17)).toNumber());
          _nw50[(_dafny.ZERO).toNumber()] = _378_v94;
          _nw50[(_dafny.ONE).toNumber()] = _dafny.Seq.update(_378_v94, _module.__default.safeIndex((_377_v93).f8, new BigNumber((_378_v94).length)), (_380_v97).dtor_cf72);
          _nw50[(new BigNumber(2)).toNumber()] = _378_v94;
          _nw50[(new BigNumber(3)).toNumber()] = _378_v94;
          _nw50[(new BigNumber(4)).toNumber()] = _378_v94;
          _nw50[(new BigNumber(5)).toNumber()] = _378_v94;
          _nw50[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(_377_v93);
          _nw50[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_378_v94, _378_v94);
          _nw50[(new BigNumber(8)).toNumber()] = _378_v94;
          _nw50[(new BigNumber(9)).toNumber()] = _dafny.Seq.update(_378_v94, _module.__default.safeIndex(_369_cf17, new BigNumber((_378_v94).length)), _377_v93);
          _nw50[(new BigNumber(10)).toNumber()] = _378_v94;
          _nw50[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_378_v94, _378_v94);
          _nw50[(new BigNumber(12)).toNumber()] = _378_v94;
          _nw50[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_378_v94, _378_v94);
          _nw50[(new BigNumber(14)).toNumber()] = _378_v94;
          _nw50[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_378_v94, _dafny.Seq.of(_377_v93)), _module.__default.safeIndex(new BigNumber((_348_v80).length), new BigNumber((_dafny.Seq.Concat(_378_v94, _dafny.Seq.of(_377_v93))).length)), _377_v93);
          _nw50[(new BigNumber(16)).toNumber()] = _378_v94;
          _382_v98 = _nw50;
          let _index34 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_382_v98).length));
          (_382_v98)[_index34] = _dafny.Seq.Concat(_378_v94, _dafny.Seq.of(_377_v93));
          let _index35 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_382_v98).length));
          let _rhs34 = _dafny.Seq.of(_377_v93, _377_v93, _377_v93, _377_v93, ((_377_v93.f7) ? (_377_v93) : (_377_v93)));
          let _rhs35 = _366_cf20;
          let _lhs18 = _382_v98;
          let _lhs19 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_382_v98).length));
          _lhs18[_lhs19] = _rhs34;
          _247_v3 = _rhs35;
          let _383_v99;
          _383_v99 = _dafny.MultiSet.fromElements(new BigNumber(132));
          let _384_v100;
          _384_v100 = _dafny.Set.fromElements(_370_cf16);
          let _385_v101;
          _385_v101 = _243_v1;
          let _386_v102;
          _386_v102 = _dafny.Set.fromElements(_dafny.Seq.update(_372_v90, _module.__default.safeIndex(_367_cf19, new BigNumber((_372_v90).length)), _346_cf35), _372_v90, _372_v90, _372_v90);
          let _387_v103;
          let _init11 = ((_388_v43, _389_v80, _390_v93) => function (_391_i10) {
            return _module.D4.create_DC13(_388_v43, _dafny.MultiSet.FromArray(_389_v80), _390_v93.f7);
          })(_296_v43, _348_v80, _377_v93);
          let _nw51 = Array((new BigNumber(25)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw51.length); _i0_11++) {
            _nw51[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _387_v103 = _nw51;
          let _392_v104;
          _392_v104 = _dafny.Set.fromElements(_387_v103);
          let _393_v105;
          let _nw52 = Array((new BigNumber(26)).toNumber());
          _nw52[(_dafny.ZERO).toNumber()] = _367_cf19;
          _nw52[(_dafny.ONE).toNumber()] = new BigNumber((_348_v80).length);
          _nw52[(new BigNumber(2)).toNumber()] = new BigNumber((_372_v90).length);
          _nw52[(new BigNumber(3)).toNumber()] = new BigNumber(565);
          _nw52[(new BigNumber(4)).toNumber()] = (_377_v93).f8;
          _nw52[(new BigNumber(5)).toNumber()] = _370_cf16;
          _nw52[(new BigNumber(6)).toNumber()] = new BigNumber((_372_v90).length);
          _nw52[(new BigNumber(7)).toNumber()] = (_377_v93).f8;
          _nw52[(new BigNumber(8)).toNumber()] = new BigNumber(((_379_v96).update(_296_v43, _247_v3)).length);
          _nw52[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.of(_383_v99)).length);
          _nw52[(new BigNumber(10)).toNumber()] = _366_cf20;
          _nw52[(new BigNumber(11)).toNumber()] = _367_cf19;
          _nw52[(new BigNumber(12)).toNumber()] = new BigNumber((_384_v100).length);
          _nw52[(new BigNumber(13)).toNumber()] = _346_cf35;
          _nw52[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(_370_cf16);
          _nw52[(new BigNumber(15)).toNumber()] = (_377_v93).f8;
          _nw52[(new BigNumber(16)).toNumber()] = new BigNumber(((_385_v101)).length);
          _nw52[(new BigNumber(17)).toNumber()] = new BigNumber((_348_v80).length);
          _nw52[(new BigNumber(18)).toNumber()] = _370_cf16;
          _nw52[(new BigNumber(19)).toNumber()] = _366_cf20;
          _nw52[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_386_v102).length));
          _nw52[(new BigNumber(21)).toNumber()] = new BigNumber((_392_v104).length);
          _nw52[(new BigNumber(22)).toNumber()] = _247_v3;
          _nw52[(new BigNumber(23)).toNumber()] = (_377_v93).f8;
          _nw52[(new BigNumber(24)).toNumber()] = _247_v3;
          _nw52[(new BigNumber(25)).toNumber()] = new BigNumber((_372_v90).length);
          _393_v105 = _nw52;
          let _394_v106;
          _394_v106 = _dafny.MultiSet.fromElements(_393_v105);
          _394_v106 = _394_v106;
          let _395_v107;
          let _nw53 = new _module.C5();
          _nw53.__ctor(_296_v43, _244_v2, _374_v92, _296_v43);
          _395_v107 = _nw53;
        } else if (_source9.is_DC5) {
          let _396___mcc_h12 = (_source9).cf13;
          let _397_cf13 = _396___mcc_h12;
          let _398_v108;
          let _399_v109;
          let _out11;
          let _out12;
          let _outcollector5 = _module.__default.m6(_246_globalState);
          _out11 = _outcollector5[0];
          _out12 = _outcollector5[1];
          _398_v108 = _out11;
          _399_v109 = _out12;
          _242_v0 = (((_module.D0.create_DC2(_247_v3, _398_v108, _296_v43, _247_v3)).dtor_cf6) ? (_module.__default.fm20(true, _246_globalState)) : (_242_v0));
          let _400_v110;
          _400_v110 = _dafny.Seq.of(_398_v108, _296_v43, !(_242_v0), _242_v0);
          let _401_v111;
          _401_v111 = _dafny.Map.Empty.slice().updateUnsafe(_400_v110,(_247_v3).isLessThanOrEqualTo(new BigNumber((_348_v80).length)));
          _401_v111 = (_401_v111).update(_dafny.Seq.of(_296_v43, _242_v0), !(!(_346_cf35).isEqualTo(_346_cf35)));
          let _402_v112;
          _402_v112 = _dafny.Seq.of(new BigNumber((_243_v1).length));
          let _403_v113;
          _403_v113 = _module.D8.create_DC28(_296_v43, new BigNumber((_397_cf13).length));
          let _404_v114;
          _404_v114 = _dafny.Map.Empty.slice().updateUnsafe((_403_v113).dtor_cf53,new BigNumber(24));
          let _405_v115;
          let _nw54 = Array((new BigNumber(29)).toNumber());
          _nw54[(_dafny.ZERO).toNumber()] = _404_v114;
          _nw54[(_dafny.ONE).toNumber()] = _404_v114;
          _nw54[(new BigNumber(2)).toNumber()] = _404_v114;
          _nw54[(new BigNumber(3)).toNumber()] = _404_v114;
          _nw54[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(716));
          _nw54[(new BigNumber(5)).toNumber()] = _404_v114;
          _nw54[(new BigNumber(6)).toNumber()] = _404_v114;
          _nw54[(new BigNumber(7)).toNumber()] = _404_v114;
          _nw54[(new BigNumber(8)).toNumber()] = _404_v114;
          _nw54[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(true,_346_cf35);
          _nw54[(new BigNumber(10)).toNumber()] = _404_v114;
          _nw54[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_398_v108,_346_cf35);
          _nw54[(new BigNumber(12)).toNumber()] = _404_v114;
          _nw54[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_296_v43,new BigNumber((_400_v110).length));
          _nw54[(new BigNumber(14)).toNumber()] = _404_v114;
          _nw54[(new BigNumber(15)).toNumber()] = _module.__default.fm39(_module.__default.fm0(_346_cf35, _246_globalState), _398_v108, _246_globalState);
          _nw54[(new BigNumber(16)).toNumber()] = _404_v114;
          _nw54[(new BigNumber(17)).toNumber()] = (_404_v114).update(_296_v43, _346_cf35);
          _nw54[(new BigNumber(18)).toNumber()] = _404_v114;
          _nw54[(new BigNumber(19)).toNumber()] = _404_v114;
          _nw54[(new BigNumber(20)).toNumber()] = _404_v114;
          _nw54[(new BigNumber(21)).toNumber()] = (_module.D11.create_DC34(_404_v114)).dtor_cf62;
          _nw54[(new BigNumber(22)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(!(_296_v43),new BigNumber(107));
          _nw54[(new BigNumber(23)).toNumber()] = _404_v114;
          _nw54[(new BigNumber(24)).toNumber()] = _404_v114;
          _nw54[(new BigNumber(25)).toNumber()] = _404_v114;
          _nw54[(new BigNumber(26)).toNumber()] = _module.__default.fm39(new BigNumber((_397_cf13).length), _242_v0, _246_globalState);
          _nw54[(new BigNumber(27)).toNumber()] = _404_v114;
          _nw54[(new BigNumber(28)).toNumber()] = _404_v114;
          _405_v115 = _nw54;
          let _406_v116;
          _406_v116 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_402_v112).length),_405_v115);
          let _407_v117;
          let _nw55 = new _module.C5();
          _nw55.__ctor(((_dafny.ZERO).minus(_247_v3)).isEqualTo(_247_v3), _244_v2, (((_406_v116).contains(_247_v3)) ? ((_406_v116).get(_247_v3)) : (_405_v115)), !(_398_v108));
          _407_v117 = _nw55;
        } else {
          let _408___mcc_h13 = (_source9).cf21;
          let _409_cf21 = _408___mcc_h13;
          let _410_v118;
          _410_v118 = _dafny.Set.fromElements(_247_v3, new BigNumber(486));
          _296_v43 = ((!(_410_v118).equals(_410_v118)) ? (_242_v0) : (_242_v0));
          let _411_v119;
          _411_v119 = _dafny.Seq.of(new BigNumber(857), _247_v3);
          let _412_v120;
          _412_v120 = _dafny.MultiSet.fromElements(_244_v2);
          let _413_v121;
          _413_v121 = _module.D10.create_DC32(_412_v120, _346_cf35, _247_v3);
          let _414_v122;
          _414_v122 = _dafny.Map.Empty.slice().updateUnsafe(((_411_v119)[_module.__default.safeIndex(_247_v3, new BigNumber((_411_v119).length))]).isLessThanOrEqualTo(_247_v3),_413_v121);
          _414_v122 = (_414_v122).update(_296_v43, _413_v121);
          let _415_v123;
          let _nw56 = Array((new BigNumber(26)).toNumber());
          _415_v123 = _nw56;
          let _416_v124;
          _416_v124 = _dafny.Seq.of(_415_v123);
          let _417_v125;
          _417_v125 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.of(_415_v123), _416_v124),_346_cf35);
          _417_v125 = (_417_v125).update(_416_v124, _247_v3);
          let _418_v126;
          let _init12 = ((_419_v43, _420_v3) => function (_421_i11) {
            return _dafny.Map.Empty.slice().updateUnsafe(_419_v43,_420_v3);
          })(_296_v43, _247_v3);
          let _nw57 = Array((new BigNumber(13)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw57.length); _i0_12++) {
            _nw57[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _418_v126 = _nw57;
          let _422_v127;
          _422_v127 = _dafny.Map.Empty.slice().updateUnsafe(!(_242_v0),_418_v126);
          let _423_v128;
          let _nw58 = new _module.C2();
          _nw58.__ctor(!(_242_v0), _247_v3, (((_422_v127).contains(_296_v43)) ? ((_422_v127).get(_296_v43)) : (_418_v126)), _242_v0, _244_v2);
          _423_v128 = _nw58;
        }
        _346_cf35 = _346_cf35;
        let _424_v129;
        let _nw59 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _424_v129 = _nw59;
        let _index36 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_424_v129).length));
        (_424_v129)[_index36] = _348_v80;
        let _index37 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_424_v129).length));
        (_424_v129)[_index37] = _348_v80;
        if (_242_v0) {
          let _425_v131;
          let _init13 = ((_426_cf35, _427_v3, _428_v79) => function (_429_i12) {
            return (_dafny.Map.Empty.slice().updateUnsafe(_426_cf35,_427_v3)).Merge(function () {
              let _coll29 = new _dafny.Map();
              for (const _compr_29 of _dafny.IntegerRange(new BigNumber(-584), new BigNumber(501))) {
                let _430_v130 = _compr_29;
                if (((new BigNumber(-584)).isLessThanOrEqualTo(_430_v130)) && ((_430_v130).isLessThan(new BigNumber(501)))) {
                  _coll29.push([(_430_v130).multipliedBy(new BigNumber(-372)),new BigNumber((_dafny.Seq.of(_428_v79)).length)]);
                }
              }
              return _coll29;
            }());
          })(_346_cf35, _247_v3, _347_v79);
          let _nw60 = Array((new BigNumber(16)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw60.length); _i0_13++) {
            _nw60[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _425_v131 = _nw60;
          _425_v131 = _425_v131;
          let _index38 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_244_v2).length));
          (_244_v2)[_index38] = _242_v0;
          let _index39 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_244_v2).length));
          (_244_v2)[_index39] = _296_v43;
          let _431_v132;
          _431_v132 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(463),_247_v3);
          let _index40 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_425_v131).length));
          (_425_v131)[_index40] = ((_242_v0) ? (_dafny.Map.Empty.slice().updateUnsafe(_247_v3,_346_cf35)) : (_431_v132));
          let _index41 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_425_v131).length));
          (_425_v131)[_index41] = _431_v132;
          let _432_v133;
          let _433_v134;
          let _out13;
          let _out14;
          let _outcollector6 = _module.__default.m6(_246_globalState);
          _out13 = _outcollector6[0];
          _out14 = _outcollector6[1];
          _432_v133 = _out13;
          _433_v134 = _out14;
          let _434_v135;
          let _init14 = ((_435_v1) => function (_436_i13) {
            return (_435_v1).IsDisjointFrom(_435_v1);
          })(_243_v1);
          let _nw61 = Array((new BigNumber(29)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw61.length); _i0_14++) {
            _nw61[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _434_v135 = _nw61;
          (_246_globalState).f1 = _434_v135;
        } else {
          let _437_v136;
          _437_v136 = _dafny.Map.Empty.slice().updateUnsafe(_247_v3,_247_v3);
          let _438_v137;
          _438_v137 = _dafny.Seq.of(_437_v136, _437_v136);
          _438_v137 = _dafny.Seq.Concat(_438_v137, _dafny.Seq.Concat(_438_v137, _module.__default.fm40(_246_globalState)));
          let _439_v138;
          let _nw62 = new _module.C0();
          _nw62.__ctor();
          _439_v138 = _nw62;
          let _440_v139;
          _440_v139 = _dafny.Seq.of(new BigNumber((_module.__default.fm25(true, _246_globalState)).length));
          let _441_v140;
          _441_v140 = _dafny.Seq.of(_247_v3, (_440_v139)[_module.__default.safeIndex(new BigNumber(((_424_v129)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_424_v129).length))]).length), new BigNumber((_440_v139).length))]);
          let _442_v141;
          _442_v141 = _dafny.MultiSet.fromElements(_244_v2, _244_v2);
          let _443_v142;
          _443_v142 = _module.D10.create_DC32(_442_v141, new BigNumber((_module.__default.fm11(_346_cf35, _242_v0, _346_cf35, _246_globalState)).length), _247_v3);
          let _444_v143;
          _444_v143 = _dafny.Set.fromElements(_247_v3, (_443_v142).dtor_cf60, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_247_v3,new BigNumber(-861))).length));
          let _445_v144;
          _445_v144 = _dafny.Map.Empty.slice().updateUnsafe(_346_cf35,_244_v2);
          let _446_v145;
          let _nw63 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Map.Empty);
          _446_v145 = _nw63;
          let _447_v146;
          _447_v146 = _module.D10.create_DC33(_296_v43);
          let _448_v147;
          let _nw64 = new _module.C5();
          _nw64.__ctor(_296_v43, (((_445_v144).contains(_247_v3)) ? ((_445_v144).get(_247_v3)) : (_244_v2)), _446_v145, (_447_v146).dtor_cf61);
          _448_v147 = _nw64;
          let _449_v148;
          _449_v148 = _dafny.Map.Empty.slice().updateUnsafe(_448_v147,(_dafny.ZERO).minus((_448_v147).fm3(_242_v0, true, _346_cf35, _247_v3, _246_globalState)));
          let _450_v149;
          _450_v149 = _dafny.Map.Empty.slice().updateUnsafe(_449_v148,_347_v79);
          let _451_v150;
          _451_v150 = _dafny.Seq.of(!(false), _242_v0, (_448_v147).f4, _296_v43, _296_v43);
          let _452_v151;
          _452_v151 = _dafny.MultiSet.fromElements((_448_v147).f4);
          let _453_v152;
          _453_v152 = _dafny.Seq.of((_424_v129)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_424_v129).length))], _348_v80, _dafny.Seq.UnicodeFromString("rn"), _348_v80);
          let _454_v153;
          let _nw65 = new _module.C5();
          _nw65.__ctor(false, _244_v2, (_448_v147).f3, _296_v43);
          _454_v153 = _nw65;
          let _455_v154;
          _455_v154 = _module.D13.create_DC42(_dafny.Seq.update(_451_v150, _module.__default.safeIndex(_346_cf35, new BigNumber((_451_v150).length)), false), (_452_v151).update(_242_v0, _module.__default.abs(_247_v3)), _dafny.Seq.update(_453_v152, _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_454_v153,(_448_v147).f4)).length), new BigNumber((_453_v152).length)), _348_v80), _296_v43, _dafny.Seq.of(_346_cf35));
          let _rhs36 = _module.__default.fm4(_module.__default.safeDivisionInt(new BigNumber(551), _346_cf35), _444_v143, _246_globalState);
          let _rhs37 = ((_dafny.ZERO).minus(_346_cf35)).isEqualTo(new BigNumber((_450_v149).length));
          let _rhs38 = (_455_v154).dtor_cf81;
          let _rhs39 = _347_v79;
          let _rhs40 = _439_v138;
          _242_v0 = _rhs36;
          _242_v0 = _rhs37;
          _441_v140 = _rhs38;
          _347_v79 = _rhs39;
          _439_v138 = _rhs40;
          let _456_v155;
          _456_v155 = _dafny.MultiSet.fromElements(_346_cf35);
          _346_cf35 = (_440_v139)[_module.__default.safeIndex(((_448_v147).fm3(!((_448_v147).f4), false, _247_v3, new BigNumber((_456_v155).cardinality()), _246_globalState)).plus((_454_v153).fm3((_448_v147).f4, (_454_v153).f5, _346_cf35, _346_cf35, _246_globalState)), new BigNumber((_440_v139).length))];
          let _457_v156;
          _457_v156 = _dafny.Seq.of(_439_v138, _439_v138);
          let _458_v157;
          _458_v157 = _dafny.Seq.of(_457_v156);
          let _459_v158;
          _459_v158 = _dafny.Map.Empty.slice().updateUnsafe((_454_v153).f5,_dafny.Seq.Create(_module.__default.abs(new BigNumber(326)), ((_460_v79) => function (_461_i14) {
            return _460_v79;
          })(_347_v79)));
          let _462_v159;
          _462_v159 = _module.D18.create_DC53(_459_v158);
          let _463_v160;
          _463_v160 = _dafny.Seq.of((_462_v159).dtor_cf99, _459_v158);
          let _464_v161;
          let _init15 = ((_465_v129) => function (_466_i15) {
            return (_466_i15).minus(new BigNumber(((_465_v129)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_465_v129).length))]).length));
          })(_424_v129);
          let _nw66 = Array((new BigNumber(10)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw66.length); _i0_15++) {
            _nw66[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _464_v161 = _nw66;
          let _467_v162;
          _467_v162 = _dafny.Set.fromElements(_464_v161);
          let _468_v163;
          let _nw67 = Array((new BigNumber(16)).toNumber());
          _nw67[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_458_v157)).cardinality());
          _nw67[(_dafny.ONE).toNumber()] = new BigNumber(371);
          _nw67[(new BigNumber(2)).toNumber()] = _346_cf35;
          _nw67[(new BigNumber(3)).toNumber()] = _346_cf35;
          _nw67[(new BigNumber(4)).toNumber()] = new BigNumber(((_463_v160)[_module.__default.safeIndex(_346_cf35, new BigNumber((_463_v160).length))]).length);
          _nw67[(new BigNumber(5)).toNumber()] = (_247_v3).minus(_247_v3);
          _nw67[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(_247_v3), _346_cf35));
          _nw67[(new BigNumber(7)).toNumber()] = _247_v3;
          _nw67[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(((_296_v43) ? (new BigNumber((_467_v162).length)) : (_346_cf35))));
          _nw67[(new BigNumber(9)).toNumber()] = _346_cf35;
          _nw67[(new BigNumber(10)).toNumber()] = new BigNumber(747);
          _nw67[(new BigNumber(11)).toNumber()] = _346_cf35;
          _nw67[(new BigNumber(12)).toNumber()] = _346_cf35;
          _nw67[(new BigNumber(13)).toNumber()] = _247_v3;
          _nw67[(new BigNumber(14)).toNumber()] = _247_v3;
          _nw67[(new BigNumber(15)).toNumber()] = _module.__default.safeModuloInt(_346_cf35, (_441_v140)[_module.__default.safeIndex(_247_v3, new BigNumber((_441_v140).length))]);
          _468_v163 = _nw67;
          let _index42 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_464_v161).length));
          (_464_v161)[_index42] = _346_cf35;
          let _index43 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_464_v161).length));
          (_464_v161)[_index43] = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(new BigNumber((_451_v150).length), _346_cf35), new BigNumber((_437_v136).length));
        }
      } else if (_source8.is_DC14) {
        let _469___mcc_h3 = (_source8).cf32;
        let _470_cf32 = _469___mcc_h3;
        let _index44 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_244_v2).length));
        (_244_v2)[_index44] = _296_v43;
        let _index45 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_244_v2).length));
        (_244_v2)[_index45] = _242_v0;
        let _471_v164;
        _471_v164 = _dafny.Seq.UnicodeFromString("fryvwxa");
        _471_v164 = (((_244_v2)[_module.__default.safeIndex(new BigNumber(818), new BigNumber((_244_v2).length))]) ? (_dafny.Seq.UnicodeFromString("h")) : (_471_v164));
        let _472_v165;
        let _init16 = ((_473_v3) => function (_474_i16) {
          return _module.__default.safeDivisionInt(_474_i16, _473_v3);
        })(_247_v3);
        let _nw68 = Array((new BigNumber(18)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw68.length); _i0_16++) {
          _nw68[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _472_v165 = _nw68;
        let _475_v166;
        _475_v166 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(214),_472_v165);
        _475_v166 = (_475_v166).update(_247_v3, _472_v165);
        let _476_v167;
        _476_v167 = new _dafny.CodePoint('v'.codePointAt(0));
        let _477_v168;
        _477_v168 = _dafny.Seq.of(_471_v164);
        let _478_v169;
        let _nw69 = Array((new BigNumber(9)).toNumber());
        _nw69[(_dafny.ZERO).toNumber()] = _471_v164;
        _nw69[(_dafny.ONE).toNumber()] = _471_v164;
        _nw69[(new BigNumber(2)).toNumber()] = _471_v164;
        _nw69[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_471_v164, _471_v164);
        _nw69[(new BigNumber(4)).toNumber()] = _471_v164;
        _nw69[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(851)), function (_479_i17) {
          return new _dafny.CodePoint('m'.codePointAt(0));
        });
        _nw69[(new BigNumber(6)).toNumber()] = _471_v164;
        _nw69[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_471_v164, _dafny.Seq.of(_476_v167));
        _nw69[(new BigNumber(8)).toNumber()] = (_477_v168)[_module.__default.safeIndex(_247_v3, new BigNumber((_477_v168).length))];
        _478_v169 = _nw69;
        let _init17 = ((_480_v164, _481_v3, _482_v167) => function (_483_i18) {
          return _dafny.Seq.Concat(_480_v164, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(686)), function (_484_i19) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          }), _module.__default.safeIndex(_481_v3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(686)), function (_485_i19) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          })).length)), _482_v167));
        })(_471_v164, _247_v3, _476_v167);
        let _nw70 = Array((new BigNumber(27)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw70.length); _i0_17++) {
          _nw70[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _478_v169 = _nw70;
      } else {
        let _486___mcc_h4 = (_source8).cf36;
        let _487_cf36 = _486___mcc_h4;
        _247_v3 = _247_v3;
        let _index46 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_244_v2).length));
        (_244_v2)[_index46] = false;
        let _488_v170;
        _488_v170 = _dafny.Seq.UnicodeFromString("rh");
        let _index47 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_244_v2).length));
        let _rhs41 = new BigNumber((_488_v170).length);
        let _rhs42 = _296_v43;
        let _lhs20 = _244_v2;
        let _lhs21 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_244_v2).length));
        _247_v3 = _rhs41;
        _lhs20[_lhs21] = _rhs42;
        let _489_v171;
        _489_v171 = _dafny.Map.Empty.slice().updateUnsafe((_244_v2)[_module.__default.safeIndex(new BigNumber(496), new BigNumber((_244_v2).length))],_247_v3);
        let _490_v172;
        _490_v172 = _dafny.Set.fromElements(_247_v3);
        let _491_v173;
        _491_v173 = _dafny.Set.fromElements(_490_v172, _490_v172);
        let _492_v174;
        _492_v174 = _dafny.MultiSet.fromElements((_244_v2)[_module.__default.safeIndex(new BigNumber(496), new BigNumber((_244_v2).length))]);
        let _493_v175;
        _493_v175 = _module.D11.create_DC34(_489_v171);
        let _494_v176;
        let _nw71 = Array((new BigNumber(22)).toNumber());
        _nw71[(_dafny.ZERO).toNumber()] = _489_v171;
        _nw71[(_dafny.ONE).toNumber()] = _489_v171;
        _nw71[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_242_v0,_247_v3);
        _nw71[(new BigNumber(3)).toNumber()] = (_489_v171).update((_244_v2)[_module.__default.safeIndex(new BigNumber(496), new BigNumber((_244_v2).length))], new BigNumber((_491_v173).length));
        _nw71[(new BigNumber(4)).toNumber()] = (_489_v171).update(_296_v43, (_dafny.ZERO).minus(new BigNumber((_492_v174).cardinality())));
        _nw71[(new BigNumber(5)).toNumber()] = _489_v171;
        _nw71[(new BigNumber(6)).toNumber()] = (_493_v175).dtor_cf62;
        _nw71[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_244_v2)[_module.__default.safeIndex(new BigNumber(496), new BigNumber((_244_v2).length))],_247_v3);
        _nw71[(new BigNumber(8)).toNumber()] = _489_v171;
        _nw71[(new BigNumber(9)).toNumber()] = _489_v171;
        _nw71[(new BigNumber(10)).toNumber()] = _489_v171;
        _nw71[(new BigNumber(11)).toNumber()] = _489_v171;
        _nw71[(new BigNumber(12)).toNumber()] = _489_v171;
        _nw71[(new BigNumber(13)).toNumber()] = (_489_v171).update((_244_v2)[_module.__default.safeIndex(new BigNumber(496), new BigNumber((_244_v2).length))], _247_v3);
        _nw71[(new BigNumber(14)).toNumber()] = _489_v171;
        _nw71[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_242_v0,_247_v3);
        _nw71[(new BigNumber(16)).toNumber()] = _489_v171;
        _nw71[(new BigNumber(17)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_242_v0,_247_v3);
        _nw71[(new BigNumber(18)).toNumber()] = _module.__default.fm39(_247_v3, !((_244_v2)[_module.__default.safeIndex(new BigNumber(496), new BigNumber((_244_v2).length))]), _246_globalState);
        _nw71[(new BigNumber(19)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_296_v43,_247_v3);
        _nw71[(new BigNumber(20)).toNumber()] = _489_v171;
        _nw71[(new BigNumber(21)).toNumber()] = _489_v171;
        _494_v176 = _nw71;
        let _495_v177;
        let _nw72 = Array((new BigNumber(4)).toNumber());
        _nw72[(_dafny.ZERO).toNumber()] = _296_v43;
        _nw72[(_dafny.ONE).toNumber()] = _296_v43;
        _nw72[(new BigNumber(2)).toNumber()] = _242_v0;
        _nw72[(new BigNumber(3)).toNumber()] = _296_v43;
        _495_v177 = _nw72;
        let _496_v178;
        let _nw73 = new _module.C1();
        _nw73.__ctor(_494_v176, !(!(_dafny.Seq.IsProperPrefixOf(_488_v170, _488_v170))), _495_v177);
        _496_v178 = _nw73;
        let _index48 = _module.__default.safeIndex(new BigNumber(83), new BigNumber((_298_v45).length));
        (_298_v45)[_index48] = _dafny.Seq.Concat(_dafny.Seq.of(_247_v3, (_dafny.ZERO).minus(_247_v3)), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("gaxsxcrjr")).length)));
        let _497_v179;
        _497_v179 = _dafny.Seq.of(_247_v3);
        let _index49 = _module.__default.safeIndex(new BigNumber(83), new BigNumber((_298_v45).length));
        (_298_v45)[_index49] = _dafny.Seq.Concat(_497_v179, _dafny.Seq.Concat(_497_v179, _497_v179));
      }
      let _498_v180;
      _498_v180 = new _dafny.CodePoint('b'.codePointAt(0));
      _498_v180 = _module.__default.fm8(_498_v180, _246_globalState);
      let _499_v181;
      let _nw74 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
      _499_v181 = _nw74;
      let _500_v183;
      let _nw75 = new _module.C1();
      _nw75.__ctor(_499_v181, _module.__default.fm4(_247_v3, function () {
        let _coll30 = new _dafny.Set();
        for (const _compr_30 of _dafny.IntegerRange(new BigNumber(604), new BigNumber(782))) {
          let _501_v182 = _compr_30;
          if (((new BigNumber(604)).isLessThanOrEqualTo(_501_v182)) && ((_501_v182).isLessThan(new BigNumber(782)))) {
            _coll30.add((_501_v182).minus(_247_v3));
          }
        }
        return _coll30;
      }(), _246_globalState), ((_242_v0) ? (_244_v2) : (_244_v2)));
      _500_v183 = _nw75;
      let _hi3 = _247_v3;
      for (let _502_i20 = new BigNumber(((_243_v1).Intersect(_243_v1)).length); _502_i20.isLessThan(_hi3); _502_i20 = _502_i20.plus(_dafny.ONE)) {
        let _503_v184;
        _503_v184 = _dafny.Map.Empty.slice().updateUnsafe(_247_v3,_296_v43);
        let _504_v185;
        _504_v185 = _dafny.MultiSet.fromElements(_503_v184);
        let _505_v186;
        _505_v186 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(_502_i20)),(((_504_v185).contains(_503_v184)) ? ((_504_v185).get(_503_v184)) : (_502_i20)));
        let _506_v187;
        _506_v187 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_505_v186).length),(_dafny.ZERO).minus((_502_i20).multipliedBy(_502_i20)));
        let _507_v188;
        let _nw76 = new _module.C5();
        _nw76.__ctor(_296_v43, _244_v2, _499_v181, _242_v0);
        _507_v188 = _nw76;
        let _508_v189;
        _508_v189 = _dafny.Seq.of(_507_v188);
        _506_v187 = (_506_v187).update(new BigNumber((_508_v189).length), _247_v3);
        let _509_v190;
        _509_v190 = _dafny.MultiSet.fromElements(_244_v2, _244_v2);
        let _510_v191;
        _510_v191 = _dafny.Map.Empty.slice().updateUnsafe(_module.D13.create_DC43(_module.D13.create_DC40(_dafny.MultiSet.fromElements((_507_v188).f5, true))),new BigNumber((_module.__default.fm18(!(_296_v43), _503_v184, _246_globalState)).length));
        _247_v3 = ((true) ? ((((_509_v190).contains(_244_v2)) ? ((_509_v190).get(_244_v2)) : (new BigNumber((_510_v191).length)))) : ((_502_i20).minus(_247_v3)));
        let _index50 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_244_v2).length));
        (_244_v2)[_index50] = true;
        let _511_v192;
        _511_v192 = _243_v1;
        let _index51 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_244_v2).length));
        (_244_v2)[_index51] = (_243_v1).IsProperSubsetOf((_511_v192));
        let _512_v193;
        let _init18 = ((_513_i20, _514_v187, _515_v3) => function (_516_i21) {
          return (_516_i21).plus((((_514_v187).contains(_513_i20)) ? ((_514_v187).get(_513_i20)) : (_515_v3)));
        })(_502_i20, _506_v187, _247_v3);
        let _nw77 = Array((new BigNumber(12)).toNumber());
        for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw77.length); _i0_18++) {
          _nw77[_i0_18] = _init18(new BigNumber(_i0_18));
        }
        _512_v193 = _nw77;
        let _index52 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_512_v193).length));
        (_512_v193)[_index52] = _502_i20;
        let _index53 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_512_v193).length));
        (_512_v193)[_index53] = _247_v3;
      }
      let _517_v194;
      _517_v194 = _dafny.Seq.UnicodeFromString("epaistma");
      let _518_v195;
      _518_v195 = _module.D15.create_DC48(_296_v43, _244_v2, _247_v3, _296_v43, _517_v194);
      let _519_v196;
      let _nw78 = Array((new BigNumber(12)).toNumber());
      _nw78[(_dafny.ZERO).toNumber()] = _244_v2;
      _nw78[(_dafny.ONE).toNumber()] = _244_v2;
      _nw78[(new BigNumber(2)).toNumber()] = _244_v2;
      _nw78[(new BigNumber(3)).toNumber()] = _244_v2;
      _nw78[(new BigNumber(4)).toNumber()] = _244_v2;
      _nw78[(new BigNumber(5)).toNumber()] = _244_v2;
      _nw78[(new BigNumber(6)).toNumber()] = _244_v2;
      _nw78[(new BigNumber(7)).toNumber()] = _244_v2;
      _nw78[(new BigNumber(8)).toNumber()] = _244_v2;
      _nw78[(new BigNumber(9)).toNumber()] = (_518_v195).dtor_cf88;
      _nw78[(new BigNumber(10)).toNumber()] = _244_v2;
      _nw78[(new BigNumber(11)).toNumber()] = _244_v2;
      _519_v196 = _nw78;
      let _520_v197;
      _520_v197 = _module.D15.create_DC47(_519_v196);
      let _source10 = _520_v197;
      if (_source10.is_DC48) {
        let _521___mcc_h14 = (_source10).cf87;
        let _522___mcc_h15 = (_source10).cf88;
        let _523___mcc_h16 = (_source10).cf89;
        let _524___mcc_h17 = (_source10).cf90;
        let _525___mcc_h18 = (_source10).cf91;
        let _526_cf91 = _525___mcc_h18;
        let _527_cf90 = _524___mcc_h17;
        let _528_cf89 = _523___mcc_h16;
        let _529_cf88 = _522___mcc_h15;
        let _530_cf87 = _521___mcc_h14;
        let _531_v198;
        _531_v198 = _dafny.Map.Empty.slice().updateUnsafe(_528_cf89,_529_cf88);
        _531_v198 = (_531_v198).update(new BigNumber(262), _244_v2);
        let _532_v199;
        let _nw79 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.of());
        _532_v199 = _nw79;
        let _533_v200;
        _533_v200 = _dafny.MultiSet.fromElements(_528_cf89);
        let _534_v201;
        let _nw80 = new _module.C2();
        _nw80.__ctor(_242_v0, _528_cf89, _499_v181, _242_v0, _529_cf88);
        _534_v201 = _nw80;
        let _535_v202;
        _535_v202 = _dafny.MultiSet.fromElements(_534_v201, _534_v201);
        let _536_v203;
        _536_v203 = _module.D12.create_DC38(_dafny.Set.fromElements(_528_cf89, new BigNumber((_533_v200).cardinality()), new BigNumber(358), new BigNumber((_526_cf91).length)), new BigNumber((_535_v202).cardinality()), _534_v201, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(943)), ((_537_v180) => function (_538_i22) {
  return _537_v180;
})(_498_v180))).length));
        let _539_v204;
        _539_v204 = _module.D12.create_DC39(_536_v203);
        let _540_v205;
        _540_v205 = _dafny.Seq.of(_539_v204);
        let _index54 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_532_v199).length));
        (_532_v199)[_index54] = _540_v205;
        let _541_v206;
        _541_v206 = _dafny.Seq.of(_247_v3, _528_cf89, _247_v3, (_534_v201).f8);
        let _542_v207;
        _542_v207 = _module.D14.create_DC45(_dafny.Seq.IsProperPrefixOf(_541_v206, _module.__default.fm31(_247_v3, !(_527_cf90), (_534_v201).f8, _296_v43, _246_globalState)));
        let _543_v208;
        _543_v208 = _dafny.Map.Empty.slice().updateUnsafe(_247_v3,_530_cf87);
        let _544_v209;
        _544_v209 = _module.D0.create_DC1(_242_v0, _534_v201.f7, _247_v3);
        let _pat_let_tv3 = _536_v203;
        let _pat_let_tv4 = _536_v203;
        let _pat_let_tv5 = _242_v0;
        let _index55 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_532_v199).length));
        let _rhs43 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(function (_pat_let5_0) {
          return function (_545_dt__update__tmp_h0) {
            return function (_pat_let6_0) {
              return function (_546_dt__update_hcf74_h0) {
                return _module.D12.create_DC39(_546_dt__update_hcf74_h0);
              }(_pat_let6_0);
            }(_pat_let_tv3);
          }(_pat_let5_0);
        }(_539_v204)), _dafny.Seq.of(function (_pat_let7_0) {
          return function (_547_dt__update__tmp_h1) {
            return function (_pat_let8_0) {
              return function (_548_dt__update_hcf74_h1) {
                return _module.D12.create_DC39(_548_dt__update_hcf74_h1);
              }(_pat_let8_0);
            }(_pat_let_tv4);
          }(_pat_let7_0);
        }(_539_v204), _539_v204)), _dafny.Seq.Concat(_540_v205, _dafny.Seq.update(_540_v205, _module.__default.safeIndex(new BigNumber((_543_v208).length), new BigNumber((_540_v205).length)), _module.D12.create_DC39(_536_v203))));
        let _rhs44 = function (_pat_let9_0) {
          return function (_549_dt__update__tmp_h2) {
            return function (_pat_let10_0) {
              return function (_550_dt__update_hcf84_h0) {
                return _module.D14.create_DC45(_550_dt__update_hcf84_h0);
              }(_pat_let10_0);
            }(_pat_let_tv5);
          }(_pat_let9_0);
        }(_542_v207);
        let _rhs45 = _244_v2;
        let _rhs46 = _530_cf87;
        let _rhs47 = _module.__default.fm4(((_541_v206)[_module.__default.safeIndex((_544_v209).dtor_cf3, new BigNumber((_541_v206).length))]).multipliedBy((_534_v201).f8), _module.__default.fm29(_246_globalState), _246_globalState);
        let _lhs22 = _532_v199;
        let _lhs23 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_532_v199).length));
        _lhs22[_lhs23] = _rhs43;
        _542_v207 = _rhs44;
        _244_v2 = _rhs45;
        _527_cf90 = _rhs46;
        _296_v43 = _rhs47;
        let _index56 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_244_v2).length));
        (_244_v2)[_index56] = ((_530_cf87) ? (!(true)) : (_242_v0));
        let _index57 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_244_v2).length));
        (_244_v2)[_index57] = false;
        _498_v180 = _498_v180;
      } else if (_source10.is_DC47) {
        let _551___mcc_h19 = (_source10).cf86;
        let _552_cf86 = _551___mcc_h19;
        let _553_v210;
        let _nw81 = new _module.C4();
        _nw81.__ctor(true);
        _553_v210 = _nw81;
        let _554_v211;
        _554_v211 = _dafny.Seq.of(_553_v210, _553_v210);
        let _rhs48 = (_553_v210).f6;
        let _rhs49 = ((_296_v43) ? ((_554_v211)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(true), _module.__default.safeIndex(_247_v3, new BigNumber((_dafny.Seq.of(true)).length)), (_553_v210).f6)).length), new BigNumber((_554_v211).length))]) : (_553_v210));
        let _rhs50 = ((_242_v0) ? (_498_v180) : (_498_v180));
        let _rhs51 = (_500_v183).fm3(_242_v0, false, _247_v3, _247_v3, _246_globalState);
        _296_v43 = _rhs48;
        _553_v210 = _rhs49;
        _498_v180 = _rhs50;
        _247_v3 = _rhs51;
        let _555_v212;
        _555_v212 = _dafny.Seq.of(new BigNumber((_517_v194).length), _247_v3, _247_v3, _247_v3);
        let _index58 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_298_v45).length));
        (_298_v45)[_index58] = _555_v212;
        let _index59 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_298_v45).length));
        (_298_v45)[_index59] = _555_v212;
        let _556_v213;
        _556_v213 = _dafny.Seq.of(_296_v43);
        let _557_v214;
        _557_v214 = _dafny.Map.Empty.slice().updateUnsafe((_556_v213)[_module.__default.safeIndex(_247_v3, new BigNumber((_556_v213).length))],_296_v43);
        let _558_v215;
        _558_v215 = _dafny.MultiSet.fromElements(_557_v214, _module.__default.fm36(_242_v0, _246_globalState));
        let _559_v216;
        _559_v216 = _dafny.Seq.of((_558_v215).IsDisjointFrom(_558_v215), _module.__default.fm4(_247_v3, _dafny.Set.fromElements(_247_v3), _246_globalState), (_553_v210).f6, _242_v0);
        _242_v0 = (_559_v216)[_module.__default.safeIndex(_247_v3, new BigNumber((_559_v216).length))];
        let _560_v217;
        _560_v217 = _dafny.Map.Empty.slice().updateUnsafe(_242_v0,_517_v194);
        let _561_v218;
        let _nw82 = new _module.C4();
        _nw82.__ctor((_560_v217).contains((_553_v210).f6));
        _561_v218 = _nw82;
      } else {
        let _562___mcc_h20 = (_source10).cf92;
        let _563_cf92 = _562___mcc_h20;
        let _564_v219;
        let _nw83 = Array((new BigNumber(11)).toNumber()).fill([]);
        _564_v219 = _nw83;
        _564_v219 = _564_v219;
        _247_v3 = new BigNumber(294);
        _242_v0 = _296_v43;
        let _565_v220;
        _565_v220 = _dafny.Seq.of(_296_v43, _242_v0, _296_v43);
        if (!((_565_v220)[_module.__default.safeIndex(_247_v3, new BigNumber((_565_v220).length))])) {
          let _566_v221;
          let _init19 = ((_567_v3) => function (_568_i23) {
            return (_568_i23).plus(_567_v3);
          })(_247_v3);
          let _nw84 = Array((new BigNumber(7)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw84.length); _i0_19++) {
            _nw84[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _566_v221 = _nw84;
          let _index60 = _module.__default.safeIndex(new BigNumber(855), new BigNumber((_566_v221).length));
          (_566_v221)[_index60] = _247_v3;
          let _569_v222;
          _569_v222 = _dafny.Map.Empty.slice().updateUnsafe(_566_v221,_296_v43);
          let _570_v223;
          _570_v223 = _module.D18.create_DC54(true, _296_v43, _247_v3, new BigNumber((_517_v194).length));
          let _571_v224;
          _571_v224 = _dafny.Map.Empty.slice().updateUnsafe(_247_v3,_296_v43);
          let _index61 = _module.__default.safeIndex(new BigNumber(855), new BigNumber((_566_v221).length));
          let _rhs52 = _module.__default.safeModuloInt(new BigNumber((_569_v222).length), (_247_v3).plus(_247_v3));
          let _rhs53 = (((_570_v223).dtor_cf100) ? ((_247_v3).multipliedBy(new BigNumber((_571_v224).length))) : (_247_v3));
          let _lhs24 = _566_v221;
          let _lhs25 = _module.__default.safeIndex(new BigNumber(855), new BigNumber((_566_v221).length));
          _lhs24[_lhs25] = _rhs52;
          _247_v3 = _rhs53;
          let _572_v225;
          let _nw85 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _572_v225 = _nw85;
          let _573_v227;
          let _init20 = ((_574_v221, _575_v3, _576_v194) => function (_577_i24) {
            return function () {
              let _coll31 = new _dafny.Map();
              for (const _compr_31 of (_dafny.Seq.of((_574_v221)[_module.__default.safeIndex(new BigNumber(855), new BigNumber((_574_v221).length))], new BigNumber(461))).Elements) {
                let _578_v226 = _compr_31;
                if (_dafny.Seq.contains(_dafny.Seq.of((_574_v221)[_module.__default.safeIndex(new BigNumber(855), new BigNumber((_574_v221).length))], new BigNumber(461)), _578_v226)) {
                  _coll31.push([(_578_v226).plus(_575_v3),new BigNumber((_576_v194).length)]);
                }
              }
              return _coll31;
            }();
          })(_566_v221, _247_v3, _517_v194);
          let _nw86 = Array((new BigNumber(5)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw86.length); _i0_20++) {
            _nw86[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _573_v227 = _nw86;
          let _579_v228;
          let _580_v229;
          let _out15;
          let _out16;
          let _outcollector7 = (_500_v183).m1((_566_v221)[_module.__default.safeIndex(new BigNumber(855), new BigNumber((_566_v221).length))], _498_v180, _572_v225, _573_v227, _246_globalState);
          _out15 = _outcollector7[0];
          _out16 = _outcollector7[1];
          _579_v228 = _out15;
          _580_v229 = _out16;
          let _581_v230;
          _581_v230 = _dafny.MultiSet.fromElements(_580_v229);
          let _582_v231;
          _582_v231 = _dafny.Map.Empty.slice().updateUnsafe(_572_v225,(_247_v3).minus((_dafny.ZERO).minus(new BigNumber((_581_v230).cardinality()))));
          _582_v231 = (_582_v231).update(_572_v225, new BigNumber(379));
          let _583_v232;
          _583_v232 = _dafny.MultiSet.fromElements(new BigNumber(929), (_566_v221)[_module.__default.safeIndex(new BigNumber(855), new BigNumber((_566_v221).length))], (_566_v221)[_module.__default.safeIndex(new BigNumber(855), new BigNumber((_566_v221).length))], new BigNumber((((_296_v43) ? (_dafny.Seq.UnicodeFromString("yuc")) : (_517_v194))).length));
          _583_v232 = (_dafny.MultiSet.fromElements((_566_v221)[_module.__default.safeIndex(new BigNumber(855), new BigNumber((_566_v221).length))])).Intersect(_583_v232);
          let _584_v233;
          _584_v233 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm8(_498_v180, _246_globalState),_498_v180);
          _247_v3 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_584_v233).length)), (_566_v221)[_module.__default.safeIndex(new BigNumber(855), new BigNumber((_566_v221).length))]), (_dafny.ZERO).minus((_566_v221)[_module.__default.safeIndex(new BigNumber(855), new BigNumber((_566_v221).length))]));
        } else {
          let _585_v234;
          let _nw87 = new _module.C3();
          _nw87.__ctor(_244_v2);
          _585_v234 = _nw87;
          let _586_v235;
          _586_v235 = _dafny.Seq.of(_244_v2, _244_v2, _244_v2);
          let _587_v236;
          _587_v236 = _dafny.Map.Empty.slice().updateUnsafe(_242_v0,_244_v2);
          let _588_v237;
          let _out17;
          _out17 = (_500_v183).m0(_247_v3, _586_v235, _587_v236, _246_globalState);
          _588_v237 = _out17;
          let _index62 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_244_v2).length));
          (_244_v2)[_index62] = _296_v43;
          let _589_v238;
          _589_v238 = _dafny.Map.Empty.slice().updateUnsafe(_247_v3,_588_v237);
          let _590_v239;
          _590_v239 = _dafny.Map.Empty.slice().updateUnsafe(((_242_v0) ? (new BigNumber((_517_v194).length)) : (_247_v3)),(new BigNumber((_589_v238).length)).multipliedBy(_247_v3));
          let _index63 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_244_v2).length));
          let _rhs54 = !_dafny.areEqual(_517_v194, _517_v194);
          let _rhs55 = _module.__default.fm25(_296_v43, _246_globalState);
          let _rhs56 = _module.__default.fm20(_296_v43, _246_globalState);
          let _lhs26 = _244_v2;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_244_v2).length));
          _lhs26[_lhs27] = _rhs54;
          _590_v239 = _rhs55;
          _242_v0 = _rhs56;
          _498_v180 = _498_v180;
          let _591_v240;
          _591_v240 = _dafny.Seq.of(_500_v183);
          let _592_v241;
          _592_v241 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_591_v240),true);
          let _593_v242;
          _593_v242 = _dafny.MultiSet.fromElements(_500_v183);
          let _594_v243;
          _594_v243 = _dafny.Map.Empty.slice().updateUnsafe(_588_v237,true);
          let _595_v244;
          let _nw88 = Array((new BigNumber(11)).toNumber());
          _nw88[(_dafny.ZERO).toNumber()] = _247_v3;
          _nw88[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_247_v3);
          _nw88[(new BigNumber(2)).toNumber()] = (_247_v3).plus(new BigNumber((_module.__default.fm11(new BigNumber(615), (((_592_v241).contains(_593_v242)) ? ((_592_v241).get(_593_v242)) : ((_244_v2)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_244_v2).length))])), _247_v3, _246_globalState)).length));
          _nw88[(new BigNumber(3)).toNumber()] = _247_v3;
          _nw88[(new BigNumber(4)).toNumber()] = new BigNumber((_517_v194).length);
          _nw88[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_594_v243).length)), _247_v3);
          _nw88[(new BigNumber(6)).toNumber()] = (_500_v183).fm16((_244_v2)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_244_v2).length))], _246_globalState);
          _nw88[(new BigNumber(7)).toNumber()] = _247_v3;
          _nw88[(new BigNumber(8)).toNumber()] = _247_v3;
          _nw88[(new BigNumber(9)).toNumber()] = _247_v3;
          _nw88[(new BigNumber(10)).toNumber()] = _247_v3;
          _595_v244 = _nw88;
          let _index64 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_595_v244).length));
          (_595_v244)[_index64] = _247_v3;
          let _index65 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_595_v244).length));
          (_595_v244)[_index65] = _247_v3;
        }
      }
      let _hi4 = _247_v3;
      for (let _596_i25 = _247_v3; _596_i25.isLessThan(_hi4); _596_i25 = _596_i25.plus(_dafny.ONE)) {
        let _597_v245;
        let _598_v246;
        let _out18;
        let _out19;
        let _outcollector8 = _module.__default.m6(_246_globalState);
        _out18 = _outcollector8[0];
        _out19 = _outcollector8[1];
        _597_v245 = _out18;
        _598_v246 = _out19;
        _498_v180 = _498_v180;
        let _599_v247;
        _599_v247 = _dafny.Map.Empty.slice().updateUnsafe(_596_i25,_242_v0);
        let _600_v248;
        _600_v248 = _dafny.Seq.of(_242_v0);
        let _601_v249;
        _601_v249 = _dafny.Seq.of(new BigNumber((_599_v247).length), new BigNumber((_600_v248).length), new BigNumber((_243_v1).length));
        let _602_v250;
        let _nw89 = new _module.C5();
        _nw89.__ctor(_597_v245, _244_v2, _499_v181, _296_v43);
        _602_v250 = _nw89;
        let _603_v251;
        _603_v251 = _dafny.MultiSet.fromElements(_602_v250, _602_v250, _602_v250);
        let _604_v252;
        let _nw90 = new _module.C2();
        _nw90.__ctor(_597_v245, _module.__default.safeModuloInt((_601_v249)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_601_v249).length))], _596_i25), _499_v181, (_603_v251).IsSubsetOf(_603_v251), _244_v2);
        _604_v252 = _nw90;
        _604_v252 = _604_v252;
        (_604_v252).f7 = (_600_v248)[_module.__default.safeIndex(_596_i25, new BigNumber((_600_v248).length))];
      }
      let _605_v253;
      let _init21 = ((_606_v0) => function (_607_i27) {
        return _module.D0.create_DC0(_606_v0);
      })(_242_v0);
      let _nw91 = Array((new BigNumber(23)).toNumber());
      for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw91.length); _i0_21++) {
        _nw91[_i0_21] = _init21(new BigNumber(_i0_21));
      }
      _605_v253 = _nw91;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_605_v253).length))) {
        let _608_i26 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_608_i26)) && ((_608_i26).isLessThan(new BigNumber((_605_v253).length))))) {
          (_605_v253)[(_608_i26)] = _module.D0.create_DC0((_dafny.MultiSet.fromElements(_247_v3)).IsDisjointFrom(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(666)), ((_609_v0, _610_v3) => function (_611_i28) {
  return _dafny.Map.Empty.slice().updateUnsafe(_609_v0,_610_v3);
})(_242_v0, _247_v3))).length))));
        }
      }
      let _612_v254;
      _612_v254 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_242_v0,_247_v3)).length),_247_v3);
      let _613_v255;
      _613_v255 = _dafny.Seq.of(_247_v3);
      let _614_v256;
      _614_v256 = _dafny.Map.Empty.slice().updateUnsafe(_296_v43,new BigNumber((_613_v255).length));
      let _615_v257;
      _615_v257 = _dafny.Map.Empty.slice().updateUnsafe(_614_v256,_612_v254);
      let _616_v259;
      _616_v259 = _module.D0.create_DC2(_247_v3, _296_v43, _296_v43, _247_v3);
      let _617_v260;
      _617_v260 = _dafny.Set.fromElements((_616_v259).dtor_cf4, _247_v3);
      let _618_v261;
      _618_v261 = _dafny.Seq.of(_242_v0, _296_v43);
      let _619_v262;
      let _nw92 = Array((new BigNumber(17)).toNumber());
      _nw92[(_dafny.ZERO).toNumber()] = _module.__default.fm25(_296_v43, _246_globalState);
      _nw92[(_dafny.ONE).toNumber()] = (_612_v254).Merge(((_612_v254).update(_247_v3, _247_v3)).update(_247_v3, _247_v3));
      _nw92[(new BigNumber(2)).toNumber()] = _612_v254;
      _nw92[(new BigNumber(3)).toNumber()] = (_612_v254).Merge(_612_v254);
      _nw92[(new BigNumber(4)).toNumber()] = _612_v254;
      _nw92[(new BigNumber(5)).toNumber()] = (((_615_v257).contains(_614_v256)) ? ((_615_v257).get(_614_v256)) : (_dafny.Map.Empty.slice().updateUnsafe(_247_v3,_247_v3)));
      _nw92[(new BigNumber(6)).toNumber()] = _612_v254;
      _nw92[(new BigNumber(7)).toNumber()] = function () {
        let _coll32 = new _dafny.Map();
        for (const _compr_32 of (_617_v260).Elements) {
          let _620_v258 = _compr_32;
          if ((_617_v260).contains(_620_v258)) {
            _coll32.push([(_620_v258).multipliedBy(_247_v3),_247_v3]);
          }
        }
        return _coll32;
      }();
      _nw92[(new BigNumber(8)).toNumber()] = (_612_v254).Merge(_612_v254);
      _nw92[(new BigNumber(9)).toNumber()] = (_612_v254).Merge(_612_v254);
      _nw92[(new BigNumber(10)).toNumber()] = (_612_v254).update(_247_v3, (_dafny.ZERO).minus((_500_v183).fm3((_618_v261)[_module.__default.safeIndex(_247_v3, new BigNumber((_618_v261).length))], _242_v0, new BigNumber((_517_v194).length), _247_v3, _246_globalState)));
      _nw92[(new BigNumber(11)).toNumber()] = _612_v254;
      _nw92[(new BigNumber(12)).toNumber()] = _612_v254;
      _nw92[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(191)), ((_621_v180) => function (_622_i30) {
        return _621_v180;
      })(_498_v180))).length),new BigNumber((_517_v194).length));
      _nw92[(new BigNumber(14)).toNumber()] = _612_v254;
      _nw92[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_247_v3,_247_v3);
      _nw92[(new BigNumber(16)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_247_v3,new BigNumber((_dafny.Seq.UnicodeFromString("q")).length));
      _619_v262 = _nw92;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_619_v262).length))) {
        let _623_i29 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_623_i29)) && ((_623_i29).isLessThan(new BigNumber((_619_v262).length))))) {
          (_619_v262)[(_623_i29)] = ((_612_v254).Merge(_612_v254)).Merge(_612_v254);
        }
      }
      let _624_v263;
      _624_v263 = _dafny.Seq.of(_244_v2, _244_v2);
      let _625_v264;
      let _out20;
      _out20 = (_500_v183).m0(_247_v3, _dafny.Seq.Concat(_624_v263, _624_v263), _dafny.Map.Empty.slice().updateUnsafe(_242_v0,(_518_v195).dtor_cf88), _246_globalState);
      _625_v264 = _out20;
      let _index66 = _module.__default.safeIndex(new BigNumber(732), new BigNumber((_244_v2).length));
      (_244_v2)[_index66] = _296_v43;
      let _index67 = _module.__default.safeIndex(new BigNumber(732), new BigNumber((_244_v2).length));
      (_244_v2)[_index67] = !(false) || (((_625_v264) ? (_242_v0) : (_296_v43)));
      let _626_i31;
      _626_i31 = _dafny.ZERO;
      L0: {
        while ((_module.__default.safeDivisionInt((_dafny.ZERO).minus(_247_v3), new BigNumber(428))).isLessThanOrEqualTo(_247_v3)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_626_i31)) {
              break L0;
            }
            _626_i31 = (_626_i31).plus(_dafny.ONE);
            let _627_v265;
            let _nw93 = Array((new BigNumber(25)).toNumber());
            _nw93[(_dafny.ZERO).toNumber()] = _247_v3;
            _nw93[(_dafny.ONE).toNumber()] = (((_612_v254).contains(new BigNumber((_517_v194).length))) ? ((_612_v254).get(new BigNumber((_517_v194).length))) : (_247_v3));
            _nw93[(new BigNumber(2)).toNumber()] = (_247_v3).plus(_247_v3);
            _nw93[(new BigNumber(3)).toNumber()] = _247_v3;
            _nw93[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(-363));
            _nw93[(new BigNumber(5)).toNumber()] = (_247_v3).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("hkxt")).length));
            _nw93[(new BigNumber(6)).toNumber()] = _247_v3;
            _nw93[(new BigNumber(7)).toNumber()] = (_247_v3).multipliedBy(_247_v3);
            _nw93[(new BigNumber(8)).toNumber()] = _247_v3;
            _nw93[(new BigNumber(9)).toNumber()] = _247_v3;
            _nw93[(new BigNumber(10)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(228), _247_v3);
            _nw93[(new BigNumber(11)).toNumber()] = _247_v3;
            _nw93[(new BigNumber(12)).toNumber()] = _module.__default.safeDivisionInt(_247_v3, _247_v3);
            _nw93[(new BigNumber(13)).toNumber()] = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(234)), ((_628_v180) => function (_629_i32) {
              return _628_v180;
            })(_498_v180))).length)).minus(_247_v3);
            _nw93[(new BigNumber(14)).toNumber()] = _247_v3;
            _nw93[(new BigNumber(15)).toNumber()] = (_247_v3).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("pmdl")).length)));
            _nw93[(new BigNumber(16)).toNumber()] = _247_v3;
            _nw93[(new BigNumber(17)).toNumber()] = _247_v3;
            _nw93[(new BigNumber(18)).toNumber()] = _module.__default.safeModuloInt(_247_v3, _247_v3);
            _nw93[(new BigNumber(19)).toNumber()] = (_247_v3).multipliedBy(_module.__default.fm0(new BigNumber(-471), _246_globalState));
            _nw93[(new BigNumber(20)).toNumber()] = new BigNumber(-39);
            _nw93[(new BigNumber(21)).toNumber()] = _247_v3;
            _nw93[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus(_247_v3);
            _nw93[(new BigNumber(23)).toNumber()] = _247_v3;
            _nw93[(new BigNumber(24)).toNumber()] = (_247_v3).multipliedBy(_247_v3);
            _627_v265 = _nw93;
            let _index68 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_627_v265).length));
            (_627_v265)[_index68] = _247_v3;
            let _index69 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_627_v265).length));
            (_627_v265)[_index69] = _247_v3;
            let _630_v266;
            _630_v266 = _dafny.Map.Empty.slice().updateUnsafe(!(_625_v264),_244_v2);
            let _631_v267;
            let _out21;
            _out21 = (_500_v183).m0(new BigNumber(290), _dafny.Seq.of(_244_v2, _244_v2, _244_v2), _630_v266, _246_globalState);
            _631_v267 = _out21;
            let _632_v268;
            _632_v268 = _dafny.MultiSet.fromElements(_498_v180);
            let _633_v269;
            let _nw94 = new _module.C2();
            _nw94.__ctor(_242_v0, new BigNumber((_618_v261).length), _499_v181, _625_v264, _244_v2);
            _633_v269 = _nw94;
            let _634_v270;
            _634_v270 = _module.D4.create_DC13(_242_v0, _632_v268, (_618_v261)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_633_v269)).length), new BigNumber((_618_v261).length))]);
            _242_v0 = (_634_v270).dtor_cf29;
            let _635_v271;
            let _init22 = ((_636_v194) => function (_637_i33) {
              return _636_v194;
            })(_517_v194);
            let _nw95 = Array((new BigNumber(22)).toNumber());
            for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw95.length); _i0_22++) {
              _nw95[_i0_22] = _init22(new BigNumber(_i0_22));
            }
            _635_v271 = _nw95;
            let _638_v272;
            let _639_v273;
            let _out22;
            let _out23;
            let _outcollector9 = (_500_v183).m1((_627_v265)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_627_v265).length))], new _dafny.CodePoint('t'.codePointAt(0)), _635_v271, _619_v262, _246_globalState);
            _out22 = _outcollector9[0];
            _out23 = _outcollector9[1];
            _638_v272 = _out22;
            _639_v273 = _out23;
          }
        }
      }
      let _640_v274;
      let _init23 = ((_641_v194, _642_v3, _643_v180) => function (_644_i34) {
        return _dafny.Seq.Concat(_dafny.Seq.update(_641_v194, _module.__default.safeIndex(_642_v3, new BigNumber((_641_v194).length)), _643_v180), _641_v194);
      })(_517_v194, _247_v3, _498_v180);
      let _nw96 = Array((new BigNumber(14)).toNumber());
      for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw96.length); _i0_23++) {
        _nw96[_i0_23] = _init23(new BigNumber(_i0_23));
      }
      _640_v274 = _nw96;
      let _index70 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_640_v274).length));
      (_640_v274)[_index70] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(737)), function (_645_i35) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      });
      let _646_v275;
      let _nw97 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
      _646_v275 = _nw97;
      let _647_v276;
      _647_v276 = _dafny.Seq.of(_646_v275);
      let _648_v277;
      _648_v277 = _dafny.Map.Empty.slice().updateUnsafe(_647_v276,_517_v194);
      let _index71 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_640_v274).length));
      (_640_v274)[_index71] = _dafny.Seq.update((((_648_v277).contains(_dafny.Seq.of(_646_v275))) ? ((_648_v277).get(_dafny.Seq.of(_646_v275))) : ((_518_v195).dtor_cf91)), _module.__default.safeIndex(((_500_v183).fm2(_246_globalState)).plus(_247_v3), new BigNumber(((((_648_v277).contains(_dafny.Seq.of(_646_v275))) ? ((_648_v277).get(_dafny.Seq.of(_646_v275))) : ((_518_v195).dtor_cf91))).length)), _498_v180);
      let _index72 = _module.__default.safeIndex(new BigNumber(732), new BigNumber((_244_v2).length));
      (_244_v2)[_index72] = (_dafny.Set.fromElements(_247_v3, _247_v3)).IsProperSubsetOf(_dafny.Set.fromElements(_247_v3));
      let _index73 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_640_v274).length));
      (_640_v274)[_index73] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(313)), ((_649_v180) => function (_650_i36) {
        return _649_v180;
      })(_498_v180)), ((_242_v0) ? (_dafny.Seq.UnicodeFromString("ndi")) : ((_640_v274)[_module.__default.safeIndex(new BigNumber(37), new BigNumber((_640_v274).length))])));
      process.stdout.write(_dafny.toString(_242_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_243_v1).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_244_v2)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_244_v2)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_244_v2)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_244_v2)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_244_v2)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_244_v2)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_244_v2)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_244_v2)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_244_v2)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_244_v2)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_244_v2)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_244_v2)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_244_v2)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_244_v2)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_246_globalState).f0).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f1)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f1)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f1)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f1)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f1)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f1)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f1)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f1)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f1)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f1)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f1)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_globalState.f1)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_247_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_296_v43));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_297_v44).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_298_v45)[new BigNumber(6)], _dafny.Seq.of(new BigNumber(8), _dafny.ZERO, _dafny.ZERO, _dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_299_v46).dtor_cf32)[new BigNumber(6)], _dafny.Seq.of(new BigNumber(8), _dafny.ZERO, _dafny.ZERO, _dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_498_v180));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_517_v194).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_518_v195).dtor_cf87));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_518_v195).dtor_cf88)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_518_v195).dtor_cf88)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_518_v195).dtor_cf88)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_518_v195).dtor_cf88)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_518_v195).dtor_cf88)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_518_v195).dtor_cf88)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_518_v195).dtor_cf88)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_518_v195).dtor_cf88)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_518_v195).dtor_cf88)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_518_v195).dtor_cf88)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_518_v195).dtor_cf88)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_518_v195).dtor_cf88)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_518_v195).dtor_cf88)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_518_v195).dtor_cf88)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_518_v195).dtor_cf89));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_518_v195).dtor_cf90));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_518_v195).dtor_cf91).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ZERO])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ZERO])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ZERO])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ZERO])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ZERO])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ZERO])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ZERO])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ZERO])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ZERO])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ZERO])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ZERO])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ZERO])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ZERO])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ZERO])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ONE])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ONE])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ONE])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ONE])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ONE])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ONE])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ONE])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ONE])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ONE])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ONE])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ONE])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ONE])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ONE])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[_dafny.ONE])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(2)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(2)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(2)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(2)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(2)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(2)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(2)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(2)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(2)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(2)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(2)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(2)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(2)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(2)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(3)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(3)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(3)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(3)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(3)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(3)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(3)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(3)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(3)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(3)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(3)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(3)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(3)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(3)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(4)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(4)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(4)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(4)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(4)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(4)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(4)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(4)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(4)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(4)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(4)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(4)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(4)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(4)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(5)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(5)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(5)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(5)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(5)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(5)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(5)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(5)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(5)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(5)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(5)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(5)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(5)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(5)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(6)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(6)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(6)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(6)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(6)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(6)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(6)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(6)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(6)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(6)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(6)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(6)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(6)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(6)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(7)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(7)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(7)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(7)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(7)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(7)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(7)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(7)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(7)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(7)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(7)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(7)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(7)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(7)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(8)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(8)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(8)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(8)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(8)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(8)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(8)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(8)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(8)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(8)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(8)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(8)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(8)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(8)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(9)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(9)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(9)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(9)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(9)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(9)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(9)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(9)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(9)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(9)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(9)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(9)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(9)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(9)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(10)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(10)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(10)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(10)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(10)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(10)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(10)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(10)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(10)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(10)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(10)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(10)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(10)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(10)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(11)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(11)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(11)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(11)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(11)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(11)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(11)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(11)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(11)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(11)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(11)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(11)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(11)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_519_v196)[new BigNumber(11)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ZERO])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ZERO])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ZERO])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ZERO])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ZERO])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ZERO])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ZERO])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ZERO])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ZERO])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ZERO])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ZERO])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ZERO])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ZERO])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ZERO])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ONE])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ONE])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ONE])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ONE])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ONE])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ONE])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ONE])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ONE])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ONE])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ONE])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ONE])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ONE])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ONE])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[_dafny.ONE])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(2)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(2)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(2)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(2)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(2)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(2)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(2)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(2)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(2)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(2)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(2)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(2)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(2)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(2)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(3)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(3)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(3)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(3)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(3)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(3)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(3)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(3)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(3)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(3)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(3)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(3)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(3)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(3)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(4)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(4)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(4)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(4)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(4)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(4)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(4)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(4)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(4)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(4)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(4)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(4)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(4)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(4)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(5)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(5)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(5)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(5)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(5)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(5)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(5)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(5)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(5)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(5)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(5)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(5)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(5)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(5)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(6)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(6)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(6)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(6)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(6)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(6)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(6)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(6)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(6)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(6)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(6)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(6)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(6)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(6)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(7)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(7)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(7)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(7)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(7)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(7)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(7)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(7)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(7)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(7)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(7)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(7)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(7)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(7)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(8)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(8)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(8)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(8)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(8)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(8)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(8)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(8)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(8)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(8)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(8)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(8)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(8)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(8)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(9)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(9)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(9)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(9)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(9)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(9)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(9)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(9)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(9)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(9)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(9)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(9)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(9)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(9)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(10)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(10)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(10)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(10)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(10)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(10)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(10)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(10)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(10)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(10)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(10)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(10)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(10)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(10)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(11)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(11)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(11)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(11)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(11)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(11)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(11)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(11)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(11)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(11)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(11)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(11)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(11)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_520_v197).dtor_cf86)[new BigNumber(11)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[_dafny.ZERO]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[_dafny.ONE]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(2)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(3)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(4)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(5)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(6)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(7)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(8)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(9)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(10)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(11)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(12)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(13)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(14)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(15)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(16)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(17)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(18)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(19)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(20)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(21)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_605_v253)[new BigNumber(22)]).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_612_v254).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_613_v255, _dafny.Seq.of(_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_614_v256).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_615_v257).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ONE),_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_616_v259).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_616_v259).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_616_v259).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_616_v259).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_617_v260).equals(_dafny.Set.fromElements(_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_618_v261, _dafny.Seq.of(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_619_v262)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_619_v262)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_619_v262)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_619_v262)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_619_v262)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_619_v262)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_619_v262)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_619_v262)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_619_v262)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_619_v262)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_619_v262)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_619_v262)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_619_v262)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_619_v262)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_619_v262)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_619_v262)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_619_v262)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_624_v263).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_625_v264));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_626_i31));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_640_v274)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_640_v274)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_640_v274)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_640_v274)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_640_v274)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_640_v274)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_640_v274)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_640_v274)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_640_v274)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_640_v274)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_640_v274)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_640_v274)[new BigNumber(11)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_640_v274)[new BigNumber(12)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_640_v274)[new BigNumber(13)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_647_v276).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_648_v277).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2, cf3) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC2(cf4, cf5, cf6, cf7) {
      let $dt = new D0(1);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(2);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC0() { return this.$tag === 2; }
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
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1 && this.cf2 === other.cf2 && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf4, other.cf4) && this.cf5 === other.cf5 && this.cf6 === other.cf6 && _dafny.areEqual(this.cf7, other.cf7);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false, false, _dafny.ZERO);
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
    static create_DC4(cf9, cf10, cf11, cf12) {
      let $dt = new D1(0);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC3(cf8) {
      let $dt = new D1(1);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10 && _dafny.areEqual(this.cf11, other.cf11) && this.cf12 === other.cf12;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf8, other.cf8);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(_dafny.Map.Empty, false, _dafny.ZERO, false);
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
    static create_DC6(cf14, cf15) {
      let $dt = new D2(0);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC7(cf16, cf17, cf18, cf19, cf20) {
      let $dt = new D2(1);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC5(cf13) {
      let $dt = new D2(2);
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC8(cf21) {
      let $dt = new D2(3);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get is_DC8() { return this.$tag === 3; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC5" + "(" + this.cf13.toVerbatimString(true) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf14, other.cf14) && _dafny.areEqual(this.cf15, other.cf15);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17) && _dafny.areEqual(this.cf18, other.cf18) && _dafny.areEqual(this.cf19, other.cf19) && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf21, other.cf21);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6(new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO);
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
    static create_DC10(cf23, cf24, cf25, cf26, cf27) {
      let $dt = new D3(0);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC9(cf22) {
      let $dt = new D3(1);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf23 === other.cf23 && _dafny.areEqual(this.cf24, other.cf24) && this.cf25 === other.cf25 && this.cf26 === other.cf26 && this.cf27 === other.cf27;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf22, other.cf22);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(false, new _dafny.CodePoint('D'.codePointAt(0)), false, false, false);
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
    static create_DC12() {
      let $dt = new D4(0);
      return $dt;
    }
    static create_DC13(cf29, cf30, cf31) {
      let $dt = new D4(1);
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC11(cf28) {
      let $dt = new D4(2);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC12";
      } else if (this.$tag === 1) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf28) + ")";
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
        return other.$tag === 1 && this.cf29 === other.cf29 && _dafny.areEqual(this.cf30, other.cf30) && this.cf31 === other.cf31;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC12();
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
    static create_DC15(cf33, cf34) {
      let $dt = new D5(0);
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC16() {
      let $dt = new D5(1);
      return $dt;
    }
    static create_DC17(cf35) {
      let $dt = new D5(2);
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC14(cf32) {
      let $dt = new D5(3);
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC18(cf36) {
      let $dt = new D5(4);
      $dt.cf36 = cf36;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get is_DC14() { return this.$tag === 3; }
    get is_DC18() { return this.$tag === 4; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf36() { return this.cf36; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC16";
      } else if (this.$tag === 2) {
        return "D5.DC17" + "(" + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 3) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 4) {
        return "D5.DC18" + "(" + _dafny.toString(this.cf36) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf33, other.cf33) && this.cf34 === other.cf34;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf32 === other.cf32;
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf36, other.cf36);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC15(_module.D4.Default(), false);
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
    static create_DC20(cf38, cf39) {
      let $dt = new D6(0);
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC19(cf37) {
      let $dt = new D6(1);
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC21(cf40) {
      let $dt = new D6(2);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC20" + "(" + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC21" + "(" + _dafny.toString(this.cf40) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf38 === other.cf38 && this.cf39 === other.cf39;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf37 === other.cf37;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf40, other.cf40);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC20(false, false);
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
    static create_DC23(cf42) {
      let $dt = new D7(0);
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC24(cf43) {
      let $dt = new D7(1);
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC22(cf41) {
      let $dt = new D7(2);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC23" + "(" + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC24" + "(" + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC22" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf42, other.cf42);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf43 === other.cf43;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf41, other.cf41);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC23(_dafny.ZERO);
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
    static create_DC26(cf45, cf46, cf47, cf48) {
      let $dt = new D8(0);
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC27(cf49, cf50, cf51, cf52) {
      let $dt = new D8(1);
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC28(cf53, cf54) {
      let $dt = new D8(2);
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC25(cf44) {
      let $dt = new D8(3);
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC29(cf55) {
      let $dt = new D8(4);
      $dt.cf55 = cf55;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get is_DC28() { return this.$tag === 2; }
    get is_DC25() { return this.$tag === 3; }
    get is_DC29() { return this.$tag === 4; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf55() { return this.cf55; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC26" + "(" + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC27" + "(" + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC28" + "(" + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ")";
      } else if (this.$tag === 3) {
        return "D8.DC25" + "(" + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 4) {
        return "D8.DC29" + "(" + _dafny.toString(this.cf55) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf45 === other.cf45 && this.cf46 === other.cf46 && this.cf47 === other.cf47 && _dafny.areEqual(this.cf48, other.cf48);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf49 === other.cf49 && this.cf50 === other.cf50 && _dafny.areEqual(this.cf51, other.cf51) && _dafny.areEqual(this.cf52, other.cf52);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf53 === other.cf53 && _dafny.areEqual(this.cf54, other.cf54);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf44, other.cf44);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf55, other.cf55);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC26(false, false, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC30(cf56) {
      let $dt = new D9(0);
      $dt.cf56 = cf56;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get dtor_cf56() { return this.cf56; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC30" + "(" + _dafny.toString(this.cf56) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf56 === other.cf56;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return [];
    }
    static Rtd() {
      return class {
        static get Default() {
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC32(cf58, cf59, cf60) {
      let $dt = new D10(0);
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      return $dt;
    }
    static create_DC33(cf61) {
      let $dt = new D10(1);
      $dt.cf61 = cf61;
      return $dt;
    }
    static create_DC31(cf57) {
      let $dt = new D10(2);
      $dt.cf57 = cf57;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get is_DC33() { return this.$tag === 1; }
    get is_DC31() { return this.$tag === 2; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf57() { return this.cf57; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC32" + "(" + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC33" + "(" + _dafny.toString(this.cf61) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC31" + "(" + _dafny.toString(this.cf57) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf58, other.cf58) && _dafny.areEqual(this.cf59, other.cf59) && _dafny.areEqual(this.cf60, other.cf60);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf61 === other.cf61;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf57 === other.cf57;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC32(_dafny.MultiSet.Empty, _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC35(cf63, cf64, cf65) {
      let $dt = new D11(0);
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      $dt.cf65 = cf65;
      return $dt;
    }
    static create_DC34(cf62) {
      let $dt = new D11(1);
      $dt.cf62 = cf62;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get is_DC34() { return this.$tag === 1; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf62() { return this.cf62; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC35" + "(" + _dafny.toString(this.cf63) + ", " + _dafny.toString(this.cf64) + ", " + _dafny.toString(this.cf65) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC34" + "(" + _dafny.toString(this.cf62) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf63, other.cf63) && this.cf64 === other.cf64 && this.cf65 === other.cf65;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf62, other.cf62);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC35(_dafny.ZERO, false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D11.Default();
        }
      };
    }
  }

  $module.D12 = class D12 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC37(cf67, cf68, cf69) {
      let $dt = new D12(0);
      $dt.cf67 = cf67;
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      return $dt;
    }
    static create_DC38(cf70, cf71, cf72, cf73) {
      let $dt = new D12(1);
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      $dt.cf72 = cf72;
      $dt.cf73 = cf73;
      return $dt;
    }
    static create_DC36(cf66) {
      let $dt = new D12(2);
      $dt.cf66 = cf66;
      return $dt;
    }
    static create_DC39(cf74) {
      let $dt = new D12(3);
      $dt.cf74 = cf74;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get is_DC38() { return this.$tag === 1; }
    get is_DC36() { return this.$tag === 2; }
    get is_DC39() { return this.$tag === 3; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf74() { return this.cf74; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC37" + "(" + _dafny.toString(this.cf67) + ", " + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC38" + "(" + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ", " + _dafny.toString(this.cf72) + ", " + _dafny.toString(this.cf73) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC36" + "(" + _dafny.toString(this.cf66) + ")";
      } else if (this.$tag === 3) {
        return "D12.DC39" + "(" + _dafny.toString(this.cf74) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf67 === other.cf67 && _dafny.areEqual(this.cf68, other.cf68) && _dafny.areEqual(this.cf69, other.cf69);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf70, other.cf70) && _dafny.areEqual(this.cf71, other.cf71) && this.cf72 === other.cf72 && _dafny.areEqual(this.cf73, other.cf73);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf66, other.cf66);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf74, other.cf74);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC37(false, _dafny.Seq.of(), _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC41(cf76) {
      let $dt = new D13(0);
      $dt.cf76 = cf76;
      return $dt;
    }
    static create_DC42(cf77, cf78, cf79, cf80, cf81) {
      let $dt = new D13(1);
      $dt.cf77 = cf77;
      $dt.cf78 = cf78;
      $dt.cf79 = cf79;
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      return $dt;
    }
    static create_DC40(cf75) {
      let $dt = new D13(2);
      $dt.cf75 = cf75;
      return $dt;
    }
    static create_DC43(cf82) {
      let $dt = new D13(3);
      $dt.cf82 = cf82;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get is_DC42() { return this.$tag === 1; }
    get is_DC40() { return this.$tag === 2; }
    get is_DC43() { return this.$tag === 3; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf82() { return this.cf82; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC41" + "(" + _dafny.toString(this.cf76) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC42" + "(" + _dafny.toString(this.cf77) + ", " + _dafny.toString(this.cf78) + ", " + _dafny.toString(this.cf79) + ", " + _dafny.toString(this.cf80) + ", " + _dafny.toString(this.cf81) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC40" + "(" + _dafny.toString(this.cf75) + ")";
      } else if (this.$tag === 3) {
        return "D13.DC43" + "(" + _dafny.toString(this.cf82) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf76, other.cf76);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf77, other.cf77) && _dafny.areEqual(this.cf78, other.cf78) && _dafny.areEqual(this.cf79, other.cf79) && this.cf80 === other.cf80 && _dafny.areEqual(this.cf81, other.cf81);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf75, other.cf75);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf82, other.cf82);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC41(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC45(cf84) {
      let $dt = new D14(0);
      $dt.cf84 = cf84;
      return $dt;
    }
    static create_DC44(cf83) {
      let $dt = new D14(1);
      $dt.cf83 = cf83;
      return $dt;
    }
    static create_DC46(cf85) {
      let $dt = new D14(2);
      $dt.cf85 = cf85;
      return $dt;
    }
    get is_DC45() { return this.$tag === 0; }
    get is_DC44() { return this.$tag === 1; }
    get is_DC46() { return this.$tag === 2; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf85() { return this.cf85; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC45" + "(" + _dafny.toString(this.cf84) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC44" + "(" + _dafny.toString(this.cf83) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC46" + "(" + _dafny.toString(this.cf85) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf84 === other.cf84;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf83, other.cf83);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf85, other.cf85);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC45(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D14.Default();
        }
      };
    }
  }

  $module.D15 = class D15 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC48(cf87, cf88, cf89, cf90, cf91) {
      let $dt = new D15(0);
      $dt.cf87 = cf87;
      $dt.cf88 = cf88;
      $dt.cf89 = cf89;
      $dt.cf90 = cf90;
      $dt.cf91 = cf91;
      return $dt;
    }
    static create_DC47(cf86) {
      let $dt = new D15(1);
      $dt.cf86 = cf86;
      return $dt;
    }
    static create_DC49(cf92) {
      let $dt = new D15(2);
      $dt.cf92 = cf92;
      return $dt;
    }
    get is_DC48() { return this.$tag === 0; }
    get is_DC47() { return this.$tag === 1; }
    get is_DC49() { return this.$tag === 2; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf92() { return this.cf92; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC48" + "(" + _dafny.toString(this.cf87) + ", " + _dafny.toString(this.cf88) + ", " + _dafny.toString(this.cf89) + ", " + _dafny.toString(this.cf90) + ", " + this.cf91.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC47" + "(" + _dafny.toString(this.cf86) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC49" + "(" + _dafny.toString(this.cf92) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf87 === other.cf87 && this.cf88 === other.cf88 && _dafny.areEqual(this.cf89, other.cf89) && this.cf90 === other.cf90 && _dafny.areEqual(this.cf91, other.cf91);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf86 === other.cf86;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf92, other.cf92);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC48(false, [], _dafny.ZERO, false, _dafny.Seq.UnicodeFromString(""));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D15.Default();
        }
      };
    }
  }

  $module.D16 = class D16 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC51(cf94, cf95, cf96, cf97) {
      let $dt = new D16(0);
      $dt.cf94 = cf94;
      $dt.cf95 = cf95;
      $dt.cf96 = cf96;
      $dt.cf97 = cf97;
      return $dt;
    }
    static create_DC50(cf93) {
      let $dt = new D16(1);
      $dt.cf93 = cf93;
      return $dt;
    }
    get is_DC51() { return this.$tag === 0; }
    get is_DC50() { return this.$tag === 1; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf97() { return this.cf97; }
    get dtor_cf93() { return this.cf93; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC51" + "(" + _dafny.toString(this.cf94) + ", " + _dafny.toString(this.cf95) + ", " + _dafny.toString(this.cf96) + ", " + _dafny.toString(this.cf97) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC50" + "(" + _dafny.toString(this.cf93) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf94 === other.cf94 && _dafny.areEqual(this.cf95, other.cf95) && this.cf96 === other.cf96 && _dafny.areEqual(this.cf97, other.cf97);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf93, other.cf93);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC51(false, _dafny.Set.Empty, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D16.Default();
        }
      };
    }
  }

  $module.D17 = class D17 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC52(cf98) {
      let $dt = new D17(0);
      $dt.cf98 = cf98;
      return $dt;
    }
    get is_DC52() { return this.$tag === 0; }
    get dtor_cf98() { return this.cf98; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC52" + "(" + _dafny.toString(this.cf98) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf98, other.cf98);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Set.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D17.Default();
        }
      };
    }
  }

  $module.D18 = class D18 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC54(cf100, cf101, cf102, cf103) {
      let $dt = new D18(0);
      $dt.cf100 = cf100;
      $dt.cf101 = cf101;
      $dt.cf102 = cf102;
      $dt.cf103 = cf103;
      return $dt;
    }
    static create_DC55(cf104, cf105) {
      let $dt = new D18(1);
      $dt.cf104 = cf104;
      $dt.cf105 = cf105;
      return $dt;
    }
    static create_DC53(cf99) {
      let $dt = new D18(2);
      $dt.cf99 = cf99;
      return $dt;
    }
    get is_DC54() { return this.$tag === 0; }
    get is_DC55() { return this.$tag === 1; }
    get is_DC53() { return this.$tag === 2; }
    get dtor_cf100() { return this.cf100; }
    get dtor_cf101() { return this.cf101; }
    get dtor_cf102() { return this.cf102; }
    get dtor_cf103() { return this.cf103; }
    get dtor_cf104() { return this.cf104; }
    get dtor_cf105() { return this.cf105; }
    get dtor_cf99() { return this.cf99; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC54" + "(" + _dafny.toString(this.cf100) + ", " + _dafny.toString(this.cf101) + ", " + _dafny.toString(this.cf102) + ", " + _dafny.toString(this.cf103) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC55" + "(" + this.cf104.toVerbatimString(true) + ", " + _dafny.toString(this.cf105) + ")";
      } else if (this.$tag === 2) {
        return "D18.DC53" + "(" + _dafny.toString(this.cf99) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf100 === other.cf100 && this.cf101 === other.cf101 && _dafny.areEqual(this.cf102, other.cf102) && _dafny.areEqual(this.cf103, other.cf103);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf104, other.cf104) && this.cf105 === other.cf105;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf99, other.cf99);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC54(false, false, _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D18.Default();
        }
      };
    }
  }

  $module.D19 = class D19 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC57() {
      let $dt = new D19(0);
      return $dt;
    }
    static create_DC58(cf107) {
      let $dt = new D19(1);
      $dt.cf107 = cf107;
      return $dt;
    }
    static create_DC56(cf106) {
      let $dt = new D19(2);
      $dt.cf106 = cf106;
      return $dt;
    }
    static create_DC59(cf108) {
      let $dt = new D19(3);
      $dt.cf108 = cf108;
      return $dt;
    }
    get is_DC57() { return this.$tag === 0; }
    get is_DC58() { return this.$tag === 1; }
    get is_DC56() { return this.$tag === 2; }
    get is_DC59() { return this.$tag === 3; }
    get dtor_cf107() { return this.cf107; }
    get dtor_cf106() { return this.cf106; }
    get dtor_cf108() { return this.cf108; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC57";
      } else if (this.$tag === 1) {
        return "D19.DC58" + "(" + _dafny.toString(this.cf107) + ")";
      } else if (this.$tag === 2) {
        return "D19.DC56" + "(" + _dafny.toString(this.cf106) + ")";
      } else if (this.$tag === 3) {
        return "D19.DC59" + "(" + _dafny.toString(this.cf108) + ")";
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
        return other.$tag === 1 && this.cf107 === other.cf107;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf106 === other.cf106;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf108, other.cf108);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC57();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D19.Default();
        }
      };
    }
  }

  $module.D20 = class D20 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC61(cf110, cf111) {
      let $dt = new D20(0);
      $dt.cf110 = cf110;
      $dt.cf111 = cf111;
      return $dt;
    }
    static create_DC62(cf112) {
      let $dt = new D20(1);
      $dt.cf112 = cf112;
      return $dt;
    }
    static create_DC60(cf109) {
      let $dt = new D20(2);
      $dt.cf109 = cf109;
      return $dt;
    }
    get is_DC61() { return this.$tag === 0; }
    get is_DC62() { return this.$tag === 1; }
    get is_DC60() { return this.$tag === 2; }
    get dtor_cf110() { return this.cf110; }
    get dtor_cf111() { return this.cf111; }
    get dtor_cf112() { return this.cf112; }
    get dtor_cf109() { return this.cf109; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC61" + "(" + _dafny.toString(this.cf110) + ", " + _dafny.toString(this.cf111) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC62" + "(" + _dafny.toString(this.cf112) + ")";
      } else if (this.$tag === 2) {
        return "D20.DC60" + "(" + _dafny.toString(this.cf109) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf110 === other.cf110 && this.cf111 === other.cf111;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf112, other.cf112);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf109, other.cf109);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC61(false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D20.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f1 = [];
      this._f0 = _dafny.Set.Empty;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm7(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(202),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,false),new _dafny.CodePoint('k'.codePointAt(0)))).length))).length),false)).length))).length)))).IsProperSubsetOf(_dafny.MultiSet.fromElements(function () {
        let _coll33 = new _dafny.Map();
        for (const _compr_33 of (_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(!(!(true)))).cardinality())), new BigNumber(107))).Elements) {
          let _651_v0 = _compr_33;
          if (_dafny.Seq.contains(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(!(!(true)))).cardinality())), new BigNumber(107)), _651_v0)) {
            _coll33.push([(_651_v0).minus(new BigNumber(429)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))).length)]);
          }
        }
        return _coll33;
      }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(130),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-100)), function (_652_i0) {
        return new BigNumber(602);
      }))).cardinality())))));
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f2 = [];
      this._f3 = [];
      this._f4 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f2() {
      let _this = this;
      return _this._f2;
    };
    set f2(value) {
      let _this = this;
      _this._f2 = value;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    __ctor(f3, f4, f2) {
      let _this = this;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f2 = f2;
      return;
    }
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(new BigNumber(163), new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f4,new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(true)).length))).length))).length)))).length));
    };
    fm1(p0, p1, globalState) {
      let _this = this;
      return ((((_this).f4) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(373),(_this).f4)) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(928),(_this).f4)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(956),(_this).f4)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(379),(_this).f4)));
    };
    fm2(globalState) {
      let _this = this;
      return ((new BigNumber((_dafny.Set.fromElements(false, (_this).f4)).length)).multipliedBy(new BigNumber(-718))).multipliedBy(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(234)), function (_653_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("vfbddobcp"))).length));
    };
    fm16(p0, globalState) {
      let _this = this;
      let _source11 = (((_this).f4) ? (_module.D2.create_DC7(new BigNumber(921), new BigNumber(443), _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(32),false)).length),new BigNumber(429)),false)).length), new BigNumber((function () {
  let _coll34 = new _dafny.Map();
  for (const _compr_34 of (_dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)))).Elements) {
    let _654_v0 = _compr_34;
    if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0))), _654_v0)) {
      _coll34.push([_654_v0,(_this).f4]);
    }
  }
  return _coll34;
}()).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f4)).length), new BigNumber(363))).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(586), new BigNumber(972), new BigNumber(-986))).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-871))).length))), new BigNumber(450), new BigNumber(181))) : (_module.D2.create_DC7(new BigNumber(268), new BigNumber((_dafny.Seq.of(function () {
  let _coll35 = new _dafny.Map();
  for (const _compr_35 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f4,new _dafny.CodePoint('q'.codePointAt(0)))).length))).Elements) {
    let _655_v1 = _compr_35;
    if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f4,new _dafny.CodePoint('q'.codePointAt(0)))).length))).contains(_655_v1)) {
      _coll35.push([_module.__default.safeDivisionInt(_655_v1, new BigNumber((_dafny.Seq.of(new BigNumber(397))).length)),(_this).f4]);
    }
  }
  return _coll35;
}(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()),(_this).f4))).length), _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f4,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-452)), function (_656_i0) {
  return new BigNumber(309);
})).length))).length), (_module.D0.create_DC1((_this).f4, true, new BigNumber(247))).dtor_cf3, new BigNumber((_dafny.Seq.of((_this).f4, (_this).f4)).length))), new BigNumber(417), new BigNumber((_dafny.Seq.of((_this).f4, (_this).f4, (_this).f4, (_this).f4, (_this).f4)).length))));
      if (_source11.is_DC6) {
        let _657___mcc_h0 = (_source11).cf14;
        let _658___mcc_h1 = (_source11).cf15;
        let _659_cf15 = _658___mcc_h1;
        let _660_cf14 = _657___mcc_h0;
        return _659_cf15;
      } else if (_source11.is_DC7) {
        let _661___mcc_h2 = (_source11).cf16;
        let _662___mcc_h3 = (_source11).cf17;
        let _663___mcc_h4 = (_source11).cf18;
        let _664___mcc_h5 = (_source11).cf19;
        let _665___mcc_h6 = (_source11).cf20;
        let _666_cf20 = _665___mcc_h6;
        let _667_cf19 = _664___mcc_h5;
        let _668_cf18 = _663___mcc_h4;
        let _669_cf17 = _662___mcc_h3;
        let _670_cf16 = _661___mcc_h2;
        return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_this).f4, (_this).f4), _dafny.Seq.of((_this).f4, (_this).f4))).length);
      } else if (_source11.is_DC5) {
        let _671___mcc_h7 = (_source11).cf13;
        let _672_cf13 = _671___mcc_h7;
        return new BigNumber((_dafny.Seq.Concat(_672_cf13, _dafny.Seq.Create(_module.__default.abs(new BigNumber(903)), function (_673_i1) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }))).length);
      } else {
        let _674___mcc_h8 = (_source11).cf21;
        let _675_cf21 = _674___mcc_h8;
        return (new BigNumber(465)).minus(new BigNumber(536));
      }
    };
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let _676_v0;
      let _nw98 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Map.Empty);
      _676_v0 = _nw98;
      let _677_v1;
      _677_v1 = _dafny.Map.Empty.slice().updateUnsafe(_676_v0,(_this).f4);
      let _678_v2;
      _678_v2 = _dafny.Seq.of(_676_v0, _676_v0);
      _677_v1 = (_677_v1).update((_678_v2)[_module.__default.safeIndex(p0, new BigNumber((_678_v2).length))], (_this).f4);
      let _679_v3;
      let _nw99 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
      _679_v3 = _nw99;
      _679_v3 = _679_v3;
      let _680_v4;
      let _nw100 = new _module.C0();
      _nw100.__ctor();
      _680_v4 = _nw100;
      _680_v4 = _680_v4;
      let _681_v5;
      _681_v5 = new BigNumber(146);
      _681_v5 = _module.__default.fm0(_681_v5, globalState);
      let _hi5 = p0;
      for (let _682_i0 = _module.__default.fm0(new BigNumber(581), globalState); _682_i0.isLessThan(_hi5); _682_i0 = _682_i0.plus(_dafny.ONE)) {
        let _index74 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_679_v3).length));
        (_679_v3)[_index74] = p0;
        let _index75 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_679_v3).length));
        (_679_v3)[_index75] = (_dafny.ZERO).minus(_681_v5);
        let _683_v6;
        _683_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f4);
        let _684_v7;
        _684_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_683_v6);
        let _index76 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_679_v3).length));
        (_679_v3)[_index76] = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((((_684_v7).contains((_this).f4)) ? ((_684_v7).get((_this).f4)) : (_683_v6)),(_this).f4)).length)).multipliedBy(_module.__default.fm0(new BigNumber(891), globalState));
        let _685_v8;
        _685_v8 = _dafny.Map.Empty.slice().updateUnsafe(_682_i0,(_this).f4);
        let _686_v9;
        _686_v9 = _dafny.Seq.of(_module.D1.create_DC4(_685_v8, (_680_v4).fm7(true, _682_i0, (_this).f4, new BigNumber(843), globalState), _682_i0, (_this).f4));
        _686_v9 = _686_v9;
        let _687_v10;
        _687_v10 = _module.D1.create_DC4((_this).fm1((_this).f4, (_this).f4, globalState), !((_this).f4), p0, (_this).f4);
        let _688_v11;
        _688_v11 = _dafny.Set.fromElements(p0, _681_v5);
        let _689_v12;
        _689_v12 = _dafny.Set.fromElements(_688_v11);
        let _index77 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_679_v3).length));
        let _index78 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_679_v3).length));
        let _rhs57 = _681_v5;
        let _rhs58 = (_this).f4;
        let _rhs59 = (_687_v10).dtor_cf11;
        let _rhs60 = (p0).plus((_dafny.ZERO).minus((_dafny.ZERO).minus((_679_v3)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((_679_v3).length))])));
        let _rhs61 = new BigNumber((_689_v12).length);
        let _lhs28 = _679_v3;
        let _lhs29 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_679_v3).length));
        let _lhs30 = _679_v3;
        let _lhs31 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_679_v3).length));
        _681_v5 = _rhs57;
        r0 = _rhs58;
        _681_v5 = _rhs59;
        _lhs28[_lhs29] = _rhs60;
        _lhs30[_lhs31] = _rhs61;
      }
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_this.f2).length))) {
        let _690_i1 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_690_i1)) && ((_690_i1).isLessThan(new BigNumber((_this.f2).length))))) {
          let _arr0 = _this.f2;
          _arr0[(_690_i1)] = (_this).f4;
        }
      }
      r0 = (_this).f4;
      return r0;
    }
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.MultiSet.Empty;
      let r1 = false;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_this.f2).length))) {
        let _691_i0 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_691_i0)) && ((_691_i0).isLessThan(new BigNumber((_this.f2).length))))) {
          let _arr1 = _this.f2;
          _arr1[(_691_i0)] = !(!(_dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("meicjlcv"), _dafny.Seq.UnicodeFromString("yewbcbs")), _dafny.Seq.UnicodeFromString("guyoe"))));
        }
      }
      let _692_v0;
      _692_v0 = _dafny.Seq.of((_this).f4, false, (_this).f4);
      let _693_v1;
      _693_v1 = _dafny.Seq.of(_dafny.Seq.of((_this).f4), _692_v0);
      let _694_v2;
      _694_v2 = _module.D0.create_DC2(new BigNumber(154), (_this).f4, !_dafny.Seq.contains(_693_v1, _692_v0), (_this).fm2(globalState));
      let _695_v3;
      _695_v3 = _dafny.Set.fromElements(p0);
      let _696_v4;
      _696_v4 = _dafny.Seq.UnicodeFromString("b");
      let _pat_let_tv6 = p0;
      let _pat_let_tv7 = _694_v2;
      let _pat_let_tv8 = _694_v2;
      _694_v2 = function (_source12) {
        if (_source12.is_DC1) {
          let _697___mcc_h0 = (_source12).cf1;
          let _698___mcc_h1 = (_source12).cf2;
          let _699___mcc_h2 = (_source12).cf3;
          let _700_cf3 = _699___mcc_h2;
          let _701_cf2 = _698___mcc_h1;
          let _702_cf1 = _697___mcc_h0;
          return _module.D0.create_DC2(_pat_let_tv6, !((_this).f4), _702_cf1, _700_cf3);
        } else if (_source12.is_DC2) {
          let _703___mcc_h3 = (_source12).cf4;
          let _704___mcc_h4 = (_source12).cf5;
          let _705___mcc_h5 = (_source12).cf6;
          let _706___mcc_h6 = (_source12).cf7;
          let _707_cf7 = _706___mcc_h6;
          let _708_cf6 = _705___mcc_h5;
          let _709_cf5 = _704___mcc_h4;
          let _710_cf4 = _703___mcc_h3;
          return _pat_let_tv7;
        } else {
          let _711___mcc_h7 = (_source12).cf0;
          let _712_cf0 = _711___mcc_h7;
          return _pat_let_tv8;
        }
      }(_module.__default.fm17(_695_v3, _696_v4, globalState));
      let _713_v5;
      let _nw101 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
      _713_v5 = _nw101;
      let _index79 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
      (_713_v5)[_index79] = p0;
      let _index80 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
      (_713_v5)[_index80] = p0;
      if (((_this).f4) || ((_this).f4)) {
        let _714_v6;
        let _nw102 = new _module.C0();
        _nw102.__ctor();
        _714_v6 = _nw102;
        let _index81 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
        (_713_v5)[_index81] = p0;
        let _715_v7;
        _715_v7 = _module.D0.create_DC1((_this).f4, (_692_v0)[_module.__default.safeIndex((_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))], new BigNumber((_692_v0).length))], (_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))]);
        if ((p0).isEqualTo((_715_v7).dtor_cf3)) {
          r1 = !(p0).isEqualTo(new BigNumber((_692_v0).length));
          let _index82 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
          (_713_v5)[_index82] = p0;
          let _716_v8;
          _716_v8 = _dafny.Set.fromElements((_this).f4, (_this).f4, (_this).f4);
          _716_v8 = _716_v8;
          let _arr2 = _this.f2;
          let _index83 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_this.f2).length));
          _arr2[_index83] = (true) === ((_this).f4);
          let _717_v9;
          let _nw103 = Array((new BigNumber(3)).toNumber()).fill([]);
          _717_v9 = _nw103;
          let _arr3 = _this.f2;
          let _index84 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_this.f2).length));
          let _rhs62 = (((p0).isLessThanOrEqualTo(new BigNumber(-366))) ? ((_this).f4) : ((_this).f4));
          let _rhs63 = _dafny.Seq.Concat(_696_v4, _dafny.Seq.update(_dafny.Seq.Concat(_696_v4, _696_v4), _module.__default.safeIndex((_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))], new BigNumber((_dafny.Seq.Concat(_696_v4, _696_v4)).length)), p1));
          let _rhs64 = _717_v9;
          let _rhs65 = (_this).f4;
          let _lhs32 = _this.f2;
          let _lhs33 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_this.f2).length));
          _lhs32[_lhs33] = _rhs62;
          _696_v4 = _rhs63;
          _717_v9 = _rhs64;
          r1 = _rhs65;
          let _718_v10;
          _718_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))]);
          let _719_v11;
          _719_v11 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f4);
          let _720_v12;
          let _nw104 = Array((new BigNumber(24)).toNumber());
          _nw104[(_dafny.ZERO).toNumber()] = (_this.f2)[_module.__default.safeIndex(new BigNumber(455), new BigNumber((_this.f2).length))];
          _nw104[(_dafny.ONE).toNumber()] = !((_718_v10).equals(_718_v10));
          _nw104[(new BigNumber(2)).toNumber()] = (_this).f4;
          _nw104[(new BigNumber(3)).toNumber()] = _dafny.areEqual(_module.__default.fm18((((_719_v11).contains(new BigNumber((_716_v8).length))) ? ((_719_v11).get(new BigNumber((_716_v8).length))) : ((_this).f4)), _719_v11, globalState), _dafny.Seq.UnicodeFromString("wsxhj"));
          _nw104[(new BigNumber(4)).toNumber()] = (_this).f4;
          _nw104[(new BigNumber(5)).toNumber()] = (_this).f4;
          _nw104[(new BigNumber(6)).toNumber()] = (_this).f4;
          _nw104[(new BigNumber(7)).toNumber()] = (_this).f4;
          _nw104[(new BigNumber(8)).toNumber()] = (_this.f2)[_module.__default.safeIndex(new BigNumber(455), new BigNumber((_this.f2).length))];
          _nw104[(new BigNumber(9)).toNumber()] = (_this).f4;
          _nw104[(new BigNumber(10)).toNumber()] = (_this).f4;
          _nw104[(new BigNumber(11)).toNumber()] = (_714_v6).fm7(false, new BigNumber(224), (_this.f2)[_module.__default.safeIndex(new BigNumber(455), new BigNumber((_this.f2).length))], new BigNumber(-510), globalState);
          _nw104[(new BigNumber(12)).toNumber()] = !((_this).f4);
          _nw104[(new BigNumber(13)).toNumber()] = ((_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))]).isLessThanOrEqualTo((_this).fm3((_this.f2)[_module.__default.safeIndex(new BigNumber(455), new BigNumber((_this.f2).length))], !((_this).f4), (_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))], p0, globalState));
          _nw104[(new BigNumber(14)).toNumber()] = true;
          _nw104[(new BigNumber(15)).toNumber()] = (_this).f4;
          _nw104[(new BigNumber(16)).toNumber()] = (_this).f4;
          _nw104[(new BigNumber(17)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_696_v4, _dafny.Seq.of(p1, p1, p1));
          _nw104[(new BigNumber(18)).toNumber()] = (_this).f4;
          _nw104[(new BigNumber(19)).toNumber()] = false;
          _nw104[(new BigNumber(20)).toNumber()] = !(true);
          _nw104[(new BigNumber(21)).toNumber()] = (_this.f2)[_module.__default.safeIndex(new BigNumber(455), new BigNumber((_this.f2).length))];
          _nw104[(new BigNumber(22)).toNumber()] = (_this).f4;
          _nw104[(new BigNumber(23)).toNumber()] = (_this).f4;
          _720_v12 = _nw104;
          (globalState).f1 = _720_v12;
        } else {
          let _index85 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
          (_713_v5)[_index85] = (_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))];
          let _index86 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
          (_713_v5)[_index86] = (_dafny.ZERO).minus((_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))]);
          let _index87 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
          (_713_v5)[_index87] = p0;
          let _index88 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
          (_713_v5)[_index88] = (_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))];
          let _721_v13;
          let _nw105 = new _module.C0();
          _nw105.__ctor();
          _721_v13 = _nw105;
        }
        let _arr4 = _this.f2;
        let _index89 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_this.f2).length));
        _arr4[_index89] = ((_dafny.ZERO).minus(p0)).isLessThan((_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))]);
        let _722_v14;
        _722_v14 = new _dafny.CodePoint('k'.codePointAt(0));
        let _723_v15;
        _723_v15 = _dafny.Seq.of((_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))]);
        let _724_v16;
        _724_v16 = _dafny.MultiSet.fromElements(p0, new BigNumber((_692_v0).length), new BigNumber(899));
        let _725_v17;
        _725_v17 = _dafny.Map.Empty.slice().updateUnsafe((_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))],_dafny.Seq.of((_this).f4));
        let _726_v19;
        _726_v19 = _dafny.MultiSet.fromElements(_713_v5);
        let _727_v20;
        _727_v20 = _dafny.Map.Empty.slice().updateUnsafe((_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))],_dafny.Seq.update(_723_v15, _module.__default.safeIndex(new BigNumber((_726_v19).cardinality()), new BigNumber((_723_v15).length)), (_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))]));
        let _arr5 = _this.f2;
        let _index90 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_this.f2).length));
        let _index91 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
        let _index92 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
        let _index93 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
        let _rhs66 = !(_dafny.MultiSet.FromArray(_723_v15)).equals(_724_v16);
        let _rhs67 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_696_v4, _696_v4), _696_v4)).length);
        let _rhs68 = (new BigNumber(((_725_v17).Merge(function () {
          let _coll36 = new _dafny.Map();
          for (const _compr_36 of ((((_727_v20).contains(p0)) ? ((_727_v20).get(p0)) : (_dafny.Seq.of(new BigNumber(713))))).Elements) {
            let _728_v18 = _compr_36;
            if (_dafny.Seq.contains((((_727_v20).contains(p0)) ? ((_727_v20).get(p0)) : (_dafny.Seq.of(new BigNumber(713)))), _728_v18)) {
              _coll36.push([_module.__default.safeDivisionInt(_728_v18, new BigNumber(72)),_692_v0]);
            }
          }
          return _coll36;
        }())).length)).minus(_module.__default.safeModuloInt((_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))], (_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))]));
        let _rhs69 = _module.__default.safeDivisionInt((_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))], (new BigNumber(532)).multipliedBy((_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))]));
        let _rhs70 = p1;
        let _lhs34 = _this.f2;
        let _lhs35 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_this.f2).length));
        let _lhs36 = _713_v5;
        let _lhs37 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
        let _lhs38 = _713_v5;
        let _lhs39 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
        let _lhs40 = _713_v5;
        let _lhs41 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
        _lhs34[_lhs35] = _rhs66;
        _lhs36[_lhs37] = _rhs67;
        _lhs38[_lhs39] = _rhs68;
        _lhs40[_lhs41] = _rhs69;
        _722_v14 = _rhs70;
        _692_v0 = (((_this).f4) ? (_692_v0) : (_dafny.Seq.Concat((_module.__default.fm19((_this).f4, p0, globalState)).dtor_cf28, _692_v0)));
      } else {
        let _729_v21;
        _729_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_this.f2);
        (_this).f2 = (((_729_v21).contains((_this).f4)) ? ((_729_v21).get((_this).f4)) : (_this.f2));
        let _730_v22;
        _730_v22 = _dafny.Seq.of(_696_v4, _696_v4, _696_v4);
        let _731_v23;
        _731_v23 = _dafny.MultiSet.fromElements(p0, new BigNumber((_696_v4).length), new BigNumber(896), p0);
        let _732_v24;
        _732_v24 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),false);
        let _733_v25;
        _733_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_732_v24).length),true);
        let _734_v26;
        _734_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_730_v22).length),_module.D2.create_DC7(p0, (_dafny.ZERO).minus((_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))]), _dafny.MultiSet.fromElements(_731_v23, _731_v23), new BigNumber(((_733_v25).update(new BigNumber((_696_v4).length), false)).length), p0));
        let _735_v27;
        _735_v27 = _module.D2.create_DC8((((_734_v26).contains((_dafny.ZERO).minus(p0))) ? ((_734_v26).get((_dafny.ZERO).minus(p0))) : (_module.D2.create_DC5(_dafny.Seq.of(p1)))));
        let _source13 = _735_v27;
        if (_source13.is_DC6) {
          let _736___mcc_h8 = (_source13).cf14;
          let _737___mcc_h9 = (_source13).cf15;
          let _738_cf15 = _737___mcc_h9;
          let _739_cf14 = _736___mcc_h8;
          let _740_v28;
          _740_v28 = _dafny.Map.Empty.slice().updateUnsafe((_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))],_738_cf15);
          let _741_v29;
          _741_v29 = _dafny.Set.fromElements(new _dafny.CodePoint('t'.codePointAt(0)), p1);
          let _742_v32;
          _742_v32 = _module.D1.create_DC4(function () {
  let _coll37 = new _dafny.Map();
  for (const _compr_37 of (function () {
    let _coll38 = new _dafny.Map();
    for (const _compr_38 of _dafny.IntegerRange(new BigNumber(-366), new BigNumber(537))) {
      let _743_v31 = _compr_38;
      if (((new BigNumber(-366)).isLessThanOrEqualTo(_743_v31)) && ((_743_v31).isLessThan(new BigNumber(537)))) {
        _coll38.push([(_743_v31).plus(new BigNumber(814)),_696_v4]);
      }
    }
    return _coll38;
  }()).Keys.Elements) {
    let _744_v30 = _compr_37;
    if ((function () {
      let _coll39 = new _dafny.Map();
      for (const _compr_39 of _dafny.IntegerRange(new BigNumber(-366), new BigNumber(537))) {
        let _743_v31 = _compr_39;
        if (((new BigNumber(-366)).isLessThanOrEqualTo(_743_v31)) && ((_743_v31).isLessThan(new BigNumber(537)))) {
          _coll39.push([(_743_v31).plus(new BigNumber(814)),_696_v4]);
        }
      }
      return _coll39;
    }()).contains(_744_v30)) {
      _coll37.push([_module.__default.safeModuloInt(_744_v30, new BigNumber((_dafny.Set.fromElements((_this).f4, (_this).f4)).length)),(_this).f4]);
    }
  }
  return _coll37;
}(), (_this).f4, p0, (_this).f4);
          let _index94 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
          let _index95 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
          let _rhs71 = (_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_740_v28).length)), p0)).minus((_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))]);
          let _rhs72 = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_741_v29)).length)).plus((_742_v32).dtor_cf11);
          let _rhs73 = (_this).f4;
          let _rhs74 = (_this).f4;
          let _lhs42 = _713_v5;
          let _lhs43 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
          let _lhs44 = _713_v5;
          let _lhs45 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
          _lhs42[_lhs43] = _rhs71;
          _lhs44[_lhs45] = _rhs72;
          r1 = _rhs73;
          r1 = _rhs74;
          _738_cf15 = (((_731_v23).contains(new BigNumber(71))) ? ((_731_v23).get(new BigNumber(71))) : ((_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))]));
          let _index96 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
          (_713_v5)[_index96] = _738_cf15;
          _739_cf14 = (_696_v4)[_module.__default.safeIndex((_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))], new BigNumber((_696_v4).length))];
        } else if (_source13.is_DC7) {
          let _745___mcc_h10 = (_source13).cf16;
          let _746___mcc_h11 = (_source13).cf17;
          let _747___mcc_h12 = (_source13).cf18;
          let _748___mcc_h13 = (_source13).cf19;
          let _749___mcc_h14 = (_source13).cf20;
          let _750_cf20 = _749___mcc_h14;
          let _751_cf19 = _748___mcc_h13;
          let _752_cf18 = _747___mcc_h12;
          let _753_cf17 = _746___mcc_h11;
          let _754_cf16 = _745___mcc_h10;
          let _arr6 = _this.f2;
          let _index97 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_this.f2).length));
          _arr6[_index97] = (_this).f4;
          let _755_v33;
          _755_v33 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm20(_module.__default.fm20((_this).f4, globalState), globalState),(_this).f4);
          let _756_v34;
          _756_v34 = _dafny.Seq.of(new BigNumber(407), _751_cf19);
          let _index98 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
          let _arr7 = _this.f2;
          let _index99 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_this.f2).length));
          let _rhs75 = (((_755_v33).contains((((_this).f4) ? ((_this).f4) : ((_this).f4)))) ? ((_755_v33).get((((_this).f4) ? ((_this).f4) : ((_this).f4)))) : ((_this).f4));
          let _rhs76 = (((_this).f4) ? ((_756_v34)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements((_this).f4)).length), new BigNumber((_756_v34).length))]) : (((_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))]).minus((_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))])));
          let _rhs77 = _module.__default.fm0(_module.__default.safeDivisionInt(_750_cf20, new BigNumber((_692_v0).length)), globalState);
          let _rhs78 = (_this).f4;
          let _lhs46 = _713_v5;
          let _lhs47 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
          let _lhs48 = _this.f2;
          let _lhs49 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_this.f2).length));
          r1 = _rhs75;
          _754_cf16 = _rhs76;
          _lhs46[_lhs47] = _rhs77;
          _lhs48[_lhs49] = _rhs78;
          let _arr8 = _this.f2;
          let _index100 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_this.f2).length));
          _arr8[_index100] = !((_731_v23).Difference(_731_v23)).contains((_750_cf20).multipliedBy(_754_cf16));
          _694_v2 = _694_v2;
          let _index101 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
          (_713_v5)[_index101] = (((_this.f2)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_this.f2).length))]) ? (_754_cf16) : (_750_cf20));
        } else if (_source13.is_DC5) {
          let _757___mcc_h15 = (_source13).cf13;
          let _758_cf13 = _757___mcc_h15;
          let _759_v35;
          let _nw106 = new _module.C0();
          _nw106.__ctor();
          _759_v35 = _nw106;
          _758_cf13 = _696_v4;
          _758_cf13 = _696_v4;
          let _760_v36;
          _760_v36 = _dafny.Map.Empty.slice().updateUnsafe(p1,_730_v22);
          let _761_v37;
          let _nw107 = Array((new BigNumber(13)).toNumber());
          _nw107[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_730_v22, _module.__default.safeIndex((_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))], new BigNumber((_730_v22).length)), _758_cf13);
          _nw107[(_dafny.ONE).toNumber()] = _730_v22;
          _nw107[(new BigNumber(2)).toNumber()] = _730_v22;
          _nw107[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_730_v22, _module.__default.fm21(globalState));
          _nw107[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_696_v4), _730_v22);
          _nw107[(new BigNumber(5)).toNumber()] = _730_v22;
          _nw107[(new BigNumber(6)).toNumber()] = _730_v22;
          _nw107[(new BigNumber(7)).toNumber()] = _dafny.Seq.of(_758_cf13, _758_cf13);
          _nw107[(new BigNumber(8)).toNumber()] = (((_760_v36).contains(_module.__default.fm22(_758_cf13, globalState))) ? ((_760_v36).get(_module.__default.fm22(_758_cf13, globalState))) : (_module.__default.fm21(globalState)));
          _nw107[(new BigNumber(9)).toNumber()] = _730_v22;
          _nw107[(new BigNumber(10)).toNumber()] = _730_v22;
          _nw107[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_730_v22, _730_v22);
          _nw107[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_730_v22, _730_v22);
          _761_v37 = _nw107;
          let _index102 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_761_v37).length));
          (_761_v37)[_index102] = _730_v22;
          let _index103 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_761_v37).length));
          (_761_v37)[_index103] = _module.__default.fm21(globalState);
        } else {
          let _762___mcc_h16 = (_source13).cf21;
          let _763_cf21 = _762___mcc_h16;
          let _764_v38;
          let _init24 = function (_765_i1) {
            return _module.D4.create_DC12();
          };
          let _nw108 = Array((new BigNumber(19)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw108.length); _i0_24++) {
            _nw108[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _764_v38 = _nw108;
          let _index104 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_764_v38).length));
          (_764_v38)[_index104] = _module.__default.fm23(globalState);
          let _766_v39;
          _766_v39 = _module.D4.create_DC12();
          let _index105 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_764_v38).length));
          (_764_v38)[_index105] = _766_v39;
          let _767_v40;
          _767_v40 = _dafny.Map.Empty.slice().updateUnsafe(_this.f2,(_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))]);
          let _768_v41;
          _768_v41 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
          _767_v40 = (_767_v40).update(_this.f2, (new BigNumber((_768_v41).length)).minus((_this).fm16(false, globalState)));
          let _nw109 = Array((new BigNumber(2)).toNumber()).fill(false);
          (_this).f2 = _nw109;
          let _arr9 = _this.f2;
          let _index106 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_this.f2).length));
          _arr9[_index106] = (_this).f4;
          let _769_v42;
          _769_v42 = _dafny.MultiSet.fromElements(p1);
          let _arr10 = _this.f2;
          let _index107 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_this.f2).length));
          _arr10[_index107] = (_769_v42).IsProperSubsetOf(_769_v42);
        }
        let _770_v43;
        _770_v43 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,true);
        let _771_v44;
        _771_v44 = _dafny.Seq.of(p0, new BigNumber((_770_v43).length), p0);
        _771_v44 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of((_this).f4)).length));
        let _772_v45;
        _772_v45 = _dafny.Seq.of(_731_v23);
        let _773_v46;
        _773_v46 = _dafny.MultiSet.fromElements((_772_v45)[_module.__default.safeIndex(p0, new BigNumber((_772_v45).length))]);
        let _774_v47;
        _774_v47 = _module.D2.create_DC7(p0, new BigNumber((_696_v4).length), _773_v46, (_dafny.ZERO).minus(p0), (_this).fm2(globalState));
        _774_v47 = _module.__default.fm24(_730_v22, globalState);
        let _775_v48;
        _775_v48 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))]);
        let _776_v49;
        _776_v49 = _dafny.Map.Empty.slice().updateUnsafe(_696_v4,_775_v48);
        let _777_v51;
        _777_v51 = _dafny.Set.fromElements(_696_v4);
        if (((_777_v51).Difference(_777_v51)).IsSubsetOf(function () {
          let _coll40 = new _dafny.Set();
          for (const _compr_40 of (_776_v49).Keys.Elements) {
            let _778_v50 = _compr_40;
            if ((_776_v49).contains(_778_v50)) {
              _coll40.add(_778_v50);
            }
          }
          return _coll40;
        }())) {
          let _779_v52;
          let _nw110 = new _module.C0();
          _nw110.__ctor();
          _779_v52 = _nw110;
          let _780_v53;
          let _init25 = ((_781_v44) => function (_782_i2) {
            return _781_v44;
          })(_771_v44);
          let _nw111 = Array((new BigNumber(28)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw111.length); _i0_25++) {
            _nw111[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _780_v53 = _nw111;
          let _783_v54;
          _783_v54 = _module.D5.create_DC14(_780_v53);
          let _784_v55;
          let _nw112 = Array((new BigNumber(24)).toNumber());
          _nw112[(_dafny.ZERO).toNumber()] = _780_v53;
          _nw112[(_dafny.ONE).toNumber()] = (_783_v54).dtor_cf32;
          _nw112[(new BigNumber(2)).toNumber()] = (((((_770_v43).contains((_this).f4)) ? ((_770_v43).get((_this).f4)) : (true))) ? (_780_v53) : (_780_v53));
          _nw112[(new BigNumber(3)).toNumber()] = _780_v53;
          _nw112[(new BigNumber(4)).toNumber()] = _780_v53;
          _nw112[(new BigNumber(5)).toNumber()] = _780_v53;
          _nw112[(new BigNumber(6)).toNumber()] = _780_v53;
          _nw112[(new BigNumber(7)).toNumber()] = _780_v53;
          _nw112[(new BigNumber(8)).toNumber()] = _780_v53;
          _nw112[(new BigNumber(9)).toNumber()] = _780_v53;
          _nw112[(new BigNumber(10)).toNumber()] = _780_v53;
          _nw112[(new BigNumber(11)).toNumber()] = _780_v53;
          _nw112[(new BigNumber(12)).toNumber()] = _780_v53;
          _nw112[(new BigNumber(13)).toNumber()] = _780_v53;
          _nw112[(new BigNumber(14)).toNumber()] = _780_v53;
          _nw112[(new BigNumber(15)).toNumber()] = _780_v53;
          _nw112[(new BigNumber(16)).toNumber()] = _780_v53;
          _nw112[(new BigNumber(17)).toNumber()] = _780_v53;
          _nw112[(new BigNumber(18)).toNumber()] = _780_v53;
          _nw112[(new BigNumber(19)).toNumber()] = _780_v53;
          _nw112[(new BigNumber(20)).toNumber()] = _780_v53;
          _nw112[(new BigNumber(21)).toNumber()] = _780_v53;
          _nw112[(new BigNumber(22)).toNumber()] = _780_v53;
          _nw112[(new BigNumber(23)).toNumber()] = _780_v53;
          _784_v55 = _nw112;
          let _index108 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_784_v55).length));
          (_784_v55)[_index108] = _780_v53;
          let _index109 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_784_v55).length));
          (_784_v55)[_index109] = _780_v53;
          let _785_v56;
          let _nw113 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Map.Empty);
          _785_v56 = _nw113;
          let _786_v58;
          _786_v58 = _dafny.Set.fromElements(_775_v48);
          let _index110 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_785_v56).length));
          (_785_v56)[_index110] = (_dafny.Map.Empty.slice().updateUnsafe(_775_v48,p1)).Merge((function () {
            let _coll41 = new _dafny.Map();
            for (const _compr_41 of (_786_v58).Elements) {
              let _787_v57 = _compr_41;
              if ((_786_v58).contains(_787_v57)) {
                _coll41.push([_787_v57,p1]);
              }
            }
            return _coll41;
          }()).update(_dafny.Map.Empty.slice().updateUnsafe((_this).f4,p0), new _dafny.CodePoint('q'.codePointAt(0))));
          let _index111 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_785_v56).length));
          (_785_v56)[_index111] = function () {
            let _coll42 = new _dafny.Map();
            for (const _compr_42 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(585)), ((_788_v5) => function (_789_i3) {
              return _dafny.Map.Empty.slice().updateUnsafe(true,(_788_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_788_v5).length))]);
            })(_713_v5))).Elements) {
              let _790_v59 = _compr_42;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(585)), ((_791_v5) => function (_789_i3) {
                return _dafny.Map.Empty.slice().updateUnsafe(true,(_791_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_791_v5).length))]);
              })(_713_v5)), _790_v59)) {
                _coll42.push([_790_v59,new _dafny.CodePoint('e'.codePointAt(0))]);
              }
            }
            return _coll42;
          }();
          let _792_v60;
          let _nw114 = new _module.C0();
          _nw114.__ctor();
          _792_v60 = _nw114;
          let _index112 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
          (_713_v5)[_index112] = (_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))];
        } else {
          let _arr11 = _this.f2;
          let _index113 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_this.f2).length));
          _arr11[_index113] = (_this).f4;
          let _arr12 = _this.f2;
          let _index114 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_this.f2).length));
          let _index115 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
          let _rhs79 = !(_dafny.Seq.IsPrefixOf(_692_v0, _692_v0));
          let _rhs80 = _module.__default.safeDivisionInt(p0, p0);
          let _lhs50 = _this.f2;
          let _lhs51 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_this.f2).length));
          let _lhs52 = _713_v5;
          let _lhs53 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
          _lhs50[_lhs51] = _rhs79;
          _lhs52[_lhs53] = _rhs80;
          _775_v48 = (_775_v48).update((_this.f2)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_this.f2).length))], p0);
          let _793_v61;
          let _nw115 = Array((new BigNumber(18)).toNumber()).fill([]);
          _793_v61 = _nw115;
          let _794_v62;
          let _nw116 = Array((new BigNumber(12)).toNumber());
          _nw116[(_dafny.ZERO).toNumber()] = _770_v43;
          _nw116[(_dafny.ONE).toNumber()] = _770_v43;
          _nw116[(new BigNumber(2)).toNumber()] = _770_v43;
          _nw116[(new BigNumber(3)).toNumber()] = _770_v43;
          _nw116[(new BigNumber(4)).toNumber()] = _770_v43;
          _nw116[(new BigNumber(5)).toNumber()] = _770_v43;
          _nw116[(new BigNumber(6)).toNumber()] = (_770_v43).update((_this.f2)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_this.f2).length))], (_this.f2)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_this.f2).length))]);
          _nw116[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(false,(_this.f2)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_this.f2).length))]);
          _nw116[(new BigNumber(8)).toNumber()] = _770_v43;
          _nw116[(new BigNumber(9)).toNumber()] = (_770_v43).update((_this.f2)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_this.f2).length))], !((_this.f2)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_this.f2).length))]));
          _nw116[(new BigNumber(10)).toNumber()] = _770_v43;
          _nw116[(new BigNumber(11)).toNumber()] = _770_v43;
          _794_v62 = _nw116;
          let _index116 = _module.__default.safeIndex(new BigNumber(814), new BigNumber((_793_v61).length));
          (_793_v61)[_index116] = _794_v62;
          let _795_v63;
          _795_v63 = _module.D6.create_DC19(_794_v62);
          let _index117 = _module.__default.safeIndex(new BigNumber(814), new BigNumber((_793_v61).length));
          (_793_v61)[_index117] = (_795_v63).dtor_cf37;
          let _796_v64;
          _796_v64 = _module.D2.create_DC6(p1, p0);
          let _index118 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
          (_713_v5)[_index118] = (_796_v64).dtor_cf15;
          let _797_v65;
          _797_v65 = _dafny.Set.fromElements(_713_v5, _713_v5, _713_v5);
          let _798_v66;
          _798_v66 = _dafny.Seq.of(_797_v65);
          r1 = (_797_v65).IsProperSubsetOf((_798_v66)[_module.__default.safeIndex(new BigNumber((_696_v4).length), new BigNumber((_798_v66).length))]);
        }
      }
      let _799_v67;
      _799_v67 = _dafny.Map.Empty.slice().updateUnsafe((_692_v0)[_module.__default.safeIndex(new BigNumber(254), new BigNumber((_692_v0).length))],(_this).f4);
      let _800_v68;
      _800_v68 = _dafny.Seq.of(_799_v67, _799_v67, (_799_v67).Merge(_799_v67));
      let _index119 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
      let _rhs81 = _800_v68;
      let _rhs82 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))]));
      let _rhs83 = !(!((_this).f4));
      let _rhs84 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nfnnwvpf"), _696_v4);
      let _rhs85 = (_this).f4;
      let _lhs54 = _713_v5;
      let _lhs55 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length));
      _800_v68 = _rhs81;
      _lhs54[_lhs55] = _rhs82;
      r1 = _rhs83;
      _696_v4 = _rhs84;
      r1 = _rhs85;
      (_this).f2 = _this.f2;
      let _801_v69;
      _801_v69 = _dafny.MultiSet.fromElements(_this.f2);
      let _802_v71;
      _802_v71 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_713_v5)[_module.__default.safeIndex(new BigNumber(400), new BigNumber((_713_v5).length))]);
      let _803_v72;
      _803_v72 = _dafny.Seq.of(new BigNumber((function () {
        let _coll43 = new _dafny.Map();
        for (const _compr_43 of (_802_v71).Keys.Elements) {
          let _804_v70 = _compr_43;
          if ((_802_v71).contains(_804_v70)) {
            _coll43.push([_804_v70,(_this).f4]);
          }
        }
        return _coll43;
      }()).length), p0);
      r0 = (_801_v69).update((((_this).f4) ? (_this.f2) : (_this.f2)), _module.__default.abs((_803_v72)[_module.__default.safeIndex(p0, new BigNumber((_803_v72).length))]));
      r1 = (_this).f4;
      return [r0, r1];
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f2 = [];
      this._f3 = [];
      this._f4 = false;
      this.f7 = false;
      this._f8 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f2() {
      let _this = this;
      return _this._f2;
    };
    set f2(value) {
      let _this = this;
      _this._f2 = value;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    __ctor(f7, f8, f3, f4, f2) {
      let _this = this;
      (_this).f7 = f7;
      (_this)._f8 = f8;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f2 = f2;
      return;
    }
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f8;
    };
    fm1(p0, p1, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe(((_this).f8).plus((_this).f8),(_this).f4);
    };
    fm2(globalState) {
      let _this = this;
      return ((_this).f8).minus(new BigNumber(396));
    };
    fm12(globalState) {
      let _this = this;
      return _this.f7;
    };
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      r0 = (_this).fm12(globalState);
      let _805_i0;
      _805_i0 = _dafny.ZERO;
      L1: {
        while (!((new BigNumber(578)).isLessThanOrEqualTo(new BigNumber(298))) || ((_this).f4)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_805_i0)) {
              break L1;
            }
            _805_i0 = (_805_i0).plus(_dafny.ONE);
            (_this).f7 = (_this).f4;
            let _806_v0;
            _806_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,p0);
            let _807_v1;
            _807_v1 = _dafny.Set.fromElements(_806_v0);
            let _808_v2;
            _808_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,new BigNumber((_807_v1).length));
            _808_v2 = (_808_v2).update(!((_this).fm12(globalState)), (p0).plus((_this).f8));
            let _809_v3;
            _809_v3 = new BigNumber(615);
            _809_v3 = p0;
            let _810_v4;
            _810_v4 = new _dafny.CodePoint('c'.codePointAt(0));
            let _811_v5;
            _811_v5 = _module.D2.create_DC6(_810_v4, (_this).f8);
            let _812_v6;
            _812_v6 = _module.D2.create_DC8(_811_v5);
            let _813_v7;
            _813_v7 = _dafny.Set.fromElements(((_this.f7) ? (_812_v6) : (_812_v6)), _812_v6, _812_v6, ((false) ? (_812_v6) : (_812_v6)), _812_v6);
            let _814_v8;
            _814_v8 = _dafny.Seq.of((_813_v7).Union(_dafny.Set.fromElements(_812_v6, _812_v6, _812_v6, _812_v6, _812_v6)), _813_v7, (_813_v7).Intersect(_813_v7), (_813_v7).Difference(_813_v7));
            _813_v7 = (_814_v8)[_module.__default.safeIndex(_809_v3, new BigNumber((_814_v8).length))];
          }
        }
      }
      let _815_v9;
      _815_v9 = _dafny.Seq.of((_dafny.ZERO).minus((_this).f8));
      let _816_i1;
      _816_i1 = _dafny.ZERO;
      L2: {
        while (!(((_815_v9)[_module.__default.safeIndex(p0, new BigNumber((_815_v9).length))]).isEqualTo((_this).f8))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_816_i1)) {
              break L2;
            }
            _816_i1 = (_816_i1).plus(_dafny.ONE);
            let _817_v10;
            _817_v10 = new BigNumber(-909);
            _817_v10 = new BigNumber(-880);
            let _818_v11;
            _818_v11 = _dafny.Seq.UnicodeFromString("dsko");
            (_this).f7 = !(!_dafny.areEqual(_818_v11, _818_v11));
            _817_v10 = p0;
            let _819_v12;
            _819_v12 = _dafny.Set.fromElements(_this.f7, _this.f7, !(true), (_this).f4);
            let _820_v13;
            _820_v13 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,(new BigNumber((_819_v12).length)).plus(_817_v10));
            _820_v13 = (_820_v13).update((new BigNumber(573)).isLessThanOrEqualTo((_this).f8), (_this).f8);
          }
        }
      }
      let _821_v14;
      _821_v14 = new _dafny.CodePoint('j'.codePointAt(0));
      let _822_v15;
      _822_v15 = _module.D3.create_DC10((_this).f4, _821_v14, _this.f7, (_this).f4, _this.f7);
      let _source14 = _822_v15;
      if (_source14.is_DC10) {
        let _823___mcc_h0 = (_source14).cf23;
        let _824___mcc_h1 = (_source14).cf24;
        let _825___mcc_h2 = (_source14).cf25;
        let _826___mcc_h3 = (_source14).cf26;
        let _827___mcc_h4 = (_source14).cf27;
        let _828_cf27 = _827___mcc_h4;
        let _829_cf26 = _826___mcc_h3;
        let _830_cf25 = _825___mcc_h2;
        let _831_cf24 = _824___mcc_h1;
        let _832_cf23 = _823___mcc_h0;
        let _833_v16;
        _833_v16 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-883)), function (_834_i2) {
          return (_dafny.ZERO).minus((_this).f8);
        }),_this.f2);
        _833_v16 = _833_v16;
        let _835_v17;
        _835_v17 = _dafny.Seq.UnicodeFromString("ngnnvi");
        let _836_v18;
        _836_v18 = _dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber((_835_v17).length)));
        let _837_v19;
        _837_v19 = _dafny.MultiSet.fromElements(_836_v18);
        let _838_v20;
        _838_v20 = _dafny.Map.Empty.slice().updateUnsafe(_830_cf25,p0);
        let _839_v21;
        _839_v21 = _dafny.Seq.of(_832_cf23, _828_cf27);
        let _840_v22;
        let _nw117 = Array((new BigNumber(13)).toNumber());
        _nw117[(_dafny.ZERO).toNumber()] = p0;
        _nw117[(_dafny.ONE).toNumber()] = p0;
        _nw117[(new BigNumber(2)).toNumber()] = (_this).fm2(globalState);
        _nw117[(new BigNumber(3)).toNumber()] = p0;
        _nw117[(new BigNumber(4)).toNumber()] = (_this).f8;
        _nw117[(new BigNumber(5)).toNumber()] = p0;
        _nw117[(new BigNumber(6)).toNumber()] = (_this).f8;
        _nw117[(new BigNumber(7)).toNumber()] = (((_837_v19).contains(_836_v18)) ? ((_837_v19).get(_836_v18)) : ((_this).f8));
        _nw117[(new BigNumber(8)).toNumber()] = (((_838_v20).contains(_832_cf23)) ? ((_838_v20).get(_832_cf23)) : (p0));
        _nw117[(new BigNumber(9)).toNumber()] = new BigNumber((_835_v17).length);
        _nw117[(new BigNumber(10)).toNumber()] = (_815_v9)[_module.__default.safeIndex(p0, new BigNumber((_815_v9).length))];
        _nw117[(new BigNumber(11)).toNumber()] = new BigNumber((_839_v21).length);
        _nw117[(new BigNumber(12)).toNumber()] = (_this).fm2(globalState);
        _840_v22 = _nw117;
        let _index120 = _module.__default.safeIndex(new BigNumber(946), new BigNumber((_840_v22).length));
        (_840_v22)[_index120] = (_this).f8;
        let _841_v23;
        _841_v23 = new BigNumber(371);
        let _index121 = _module.__default.safeIndex(new BigNumber(946), new BigNumber((_840_v22).length));
        let _rhs86 = _832_cf23;
        let _rhs87 = (_this).f8;
        let _rhs88 = (_this).f8;
        let _rhs89 = _828_cf27;
        let _rhs90 = (_this).fm2(globalState);
        let _lhs56 = _this;
        let _lhs57 = _840_v22;
        let _lhs58 = _module.__default.safeIndex(new BigNumber(946), new BigNumber((_840_v22).length));
        _lhs56.f7 = _rhs86;
        _lhs57[_lhs58] = _rhs87;
        _841_v23 = _rhs88;
        _832_cf23 = _rhs89;
        _841_v23 = _rhs90;
        let _842_v24;
        _842_v24 = _dafny.Map.Empty.slice().updateUnsafe(_821_v14,_829_cf26);
        let _843_v25;
        _843_v25 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC9(_842_v24),_832_cf23);
        let _844_v26;
        _844_v26 = _module.D3.create_DC9(_842_v24);
        _843_v25 = (_843_v25).update((((_this).f4) ? (_844_v26) : (_module.D3.create_DC9(_842_v24))), _dafny.Seq.IsPrefixOf(_module.__default.fm13(_830_cf25, _830_cf25, _832_cf23, globalState), _835_v17));
        _828_cf27 = (_this).fm12(globalState);
      } else {
        let _845___mcc_h5 = (_source14).cf22;
        let _846_cf22 = _845___mcc_h5;
        let _847_v27;
        _847_v27 = _dafny.Seq.of(_this.f7, _this.f7);
        (_this).f7 = _dafny.areEqual(_847_v27, _847_v27);
        let _848_v28;
        let _nw118 = new _module.C0();
        _nw118.__ctor();
        _848_v28 = _nw118;
        if (_this.f7) {
          (_this).f7 = (_847_v27)[_module.__default.safeIndex(new BigNumber(-704), new BigNumber((_847_v27).length))];
          let _849_v29;
          let _nw119 = Array((new BigNumber(2)).toNumber()).fill(_module.D0.Default());
          _849_v29 = _nw119;
          let _850_v30;
          _850_v30 = _module.D0.create_DC0((_this).f4);
          let _index122 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_849_v29).length));
          (_849_v29)[_index122] = _850_v30;
          let _851_v31;
          _851_v31 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f8),(_this).f4);
          let _852_v32;
          _852_v32 = _module.D1.create_DC4(_851_v31, _this.f7, (_this).f8, (_this).f4);
          let _853_v33;
          _853_v33 = _dafny.MultiSet.fromElements((_this).f4);
          let _854_v34;
          _854_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f4);
          let _index123 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_849_v29).length));
          (_849_v29)[_index123] = _module.__default.fm14((_852_v32).dtor_cf11, p0, (_853_v33).IsProperSubsetOf(_853_v33), (((_854_v34).contains(false)) ? ((_854_v34).get(false)) : ((_this).f4)), globalState);
          let _855_v35;
          _855_v35 = _dafny.Set.fromElements(_821_v14);
          let _856_v36;
          _856_v36 = _dafny.Map.Empty.slice().updateUnsafe(_855_v35,_852_v32);
          _856_v36 = (_856_v36).update(_855_v35, _852_v32);
          let _857_v37;
          _857_v37 = new BigNumber(697);
          let _858_v38;
          _858_v38 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(new BigNumber((_853_v33).cardinality())).multipliedBy(p0));
          _857_v37 = (((_858_v38).contains(_this.f7)) ? ((_858_v38).get(_this.f7)) : (_857_v37));
          let _859_v39;
          let _860_v40;
          let _out24;
          let _out25;
          let _outcollector10 = (_this).m4((_this).f4, globalState);
          _out24 = _outcollector10[0];
          _out25 = _outcollector10[1];
          _859_v39 = _out24;
          _860_v40 = _out25;
        } else {
          (_this).m5((_this).f4, (_this.f7) && (_this.f7), globalState);
          let _arr13 = _this.f2;
          let _index124 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_this.f2).length));
          _arr13[_index124] = !(_dafny.Seq.IsPrefixOf(_847_v27, _dafny.Seq.of(_this.f7, (_this).f4)));
          let _arr14 = _this.f2;
          let _index125 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_this.f2).length));
          _arr14[_index125] = (_this).fm12(globalState);
          let _861_v41;
          _861_v41 = _dafny.Set.fromElements(_821_v14);
          _861_v41 = _861_v41;
          let _862_v42;
          _862_v42 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-992),(_this).f4);
          let _863_v43;
          _863_v43 = _dafny.Map.Empty.slice().updateUnsafe(_862_v42,!((_this.f2)[_module.__default.safeIndex(new BigNumber(832), new BigNumber((_this.f2).length))]));
          _863_v43 = (_863_v43).update((_this).fm1((_this.f2)[_module.__default.safeIndex(new BigNumber(832), new BigNumber((_this.f2).length))], _this.f7, globalState), (_this).f4);
          let _864_v44;
          let _nw120 = Array((_dafny.ONE).toNumber());
          _nw120[(_dafny.ZERO).toNumber()] = (_this).f4;
          _864_v44 = _nw120;
          let _865_v45;
          _865_v45 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).f8,_864_v44),_847_v27);
          let _866_v46;
          _866_v46 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_864_v44);
          _865_v45 = (_865_v45).update((_866_v46).Merge(_866_v46), _dafny.Seq.Concat(_847_v27, _847_v27));
        }
        let _867_v47;
        _867_v47 = _dafny.Seq.UnicodeFromString("mseo");
        _821_v14 = (_867_v47)[_module.__default.safeIndex((_this).f8, new BigNumber((_867_v47).length))];
      }
      if (_this.f7) {
        let _868_v48;
        let _nw121 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
        _868_v48 = _nw121;
        let _869_v49;
        _869_v49 = _dafny.Seq.of(_868_v48);
        let _870_v50;
        _870_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_868_v48);
        let _871_v51;
        let _nw122 = Array((new BigNumber(28)).toNumber());
        _nw122[(_dafny.ZERO).toNumber()] = _868_v48;
        _nw122[(_dafny.ONE).toNumber()] = _868_v48;
        _nw122[(new BigNumber(2)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(3)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(4)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(5)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(6)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(7)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(8)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(9)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(10)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(11)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(12)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(13)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(14)).toNumber()] = (_869_v49)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f8), new BigNumber((_869_v49).length))];
        _nw122[(new BigNumber(15)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(16)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(17)).toNumber()] = (((_870_v50).contains(p0)) ? ((_870_v50).get(p0)) : (_868_v48));
        _nw122[(new BigNumber(18)).toNumber()] = (_869_v49)[_module.__default.safeIndex(p0, new BigNumber((_869_v49).length))];
        _nw122[(new BigNumber(19)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(20)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(21)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(22)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(23)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(24)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(25)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(26)).toNumber()] = _868_v48;
        _nw122[(new BigNumber(27)).toNumber()] = _868_v48;
        _871_v51 = _nw122;
        _871_v51 = _871_v51;
        (_this).f7 = _this.f7;
        let _872_v52;
        _872_v52 = new BigNumber(555);
        _872_v52 = p0;
        let _873_v53;
        _873_v53 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f4);
        let _874_v54;
        _874_v54 = _module.D0.create_DC0((((_873_v53).contains(_this.f7)) ? ((_873_v53).get(_this.f7)) : (!((_this).f4))));
        let _source15 = function (_pat_let11_0) {
          return function (_875_dt__update__tmp_h0) {
            return function (_pat_let12_0) {
              return function (_876_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_876_dt__update_hcf0_h0);
              }(_pat_let12_0);
            }(true);
          }(_pat_let11_0);
        }(_874_v54);
        if (_source15.is_DC1) {
          let _877___mcc_h6 = (_source15).cf1;
          let _878___mcc_h7 = (_source15).cf2;
          let _879___mcc_h8 = (_source15).cf3;
          let _880_cf3 = _879___mcc_h8;
          let _881_cf2 = _878___mcc_h7;
          let _882_cf1 = _877___mcc_h6;
          let _883_v55;
          _883_v55 = _dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0)), _821_v14);
          let _884_v56;
          _884_v56 = _dafny.Seq.of(_883_v55, _883_v55);
          let _885_v57;
          _885_v57 = _dafny.Map.Empty.slice().updateUnsafe(_884_v56,_872_v52);
          _885_v57 = (_885_v57).update(_884_v56, p0);
          let _arr15 = _this.f2;
          let _index126 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_this.f2).length));
          _arr15[_index126] = _882_cf1;
          let _arr16 = _this.f2;
          let _index127 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_this.f2).length));
          _arr16[_index127] = !(_this.f7) || ((_this).fm12(globalState));
          let _886_v58;
          let _init26 = ((_887_v55) => function (_888_i3) {
            return _887_v55;
          })(_883_v55);
          let _nw123 = Array((new BigNumber(23)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw123.length); _i0_26++) {
            _nw123[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _886_v58 = _nw123;
          let _889_v59;
          _889_v59 = _dafny.MultiSet.fromElements(new BigNumber(367), (_this).f8);
          let _890_v60;
          _890_v60 = _dafny.Map.Empty.slice().updateUnsafe(_886_v58,(_889_v59).IsDisjointFrom(_889_v59));
          _881_cf2 = (((_890_v60).contains(_886_v58)) ? ((_890_v60).get(_886_v58)) : (_881_cf2));
          _872_v52 = _module.__default.safeModuloInt((((_this).fm12(globalState)) ? (_872_v52) : (_880_cf3)), _872_v52);
        } else if (_source15.is_DC2) {
          let _891___mcc_h9 = (_source15).cf4;
          let _892___mcc_h10 = (_source15).cf5;
          let _893___mcc_h11 = (_source15).cf6;
          let _894___mcc_h12 = (_source15).cf7;
          let _895_cf7 = _894___mcc_h12;
          let _896_cf6 = _893___mcc_h11;
          let _897_cf5 = _892___mcc_h10;
          let _898_cf4 = _891___mcc_h9;
          let _899_v61;
          _899_v61 = _dafny.Set.fromElements(new BigNumber(892), new BigNumber(516));
          let _900_v62;
          _900_v62 = _dafny.Set.fromElements(_896_cf6);
          _897_cf5 = !((_899_v61).IsProperSubsetOf(_899_v61)) || ((_900_v62).IsSubsetOf(_dafny.Set.fromElements(_896_cf6, _897_cf5)));
          let _index128 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_868_v48).length));
          (_868_v48)[_index128] = (_895_cf7).plus(_895_cf7);
          let _index129 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_868_v48).length));
          (_868_v48)[_index129] = _872_v52;
          r0 = _897_cf5;
          let _901_v63;
          _901_v63 = _module.D0.create_DC1(false, !(_897_cf5) || ((_this).f4), new BigNumber(450));
          let _902_v64;
          _902_v64 = _dafny.Seq.of(false);
          let _903_v65;
          _903_v65 = _dafny.Map.Empty.slice().updateUnsafe(_872_v52,_896_cf6);
          let _904_v66;
          _904_v66 = _dafny.Map.Empty.slice().updateUnsafe(_898_cf4,_903_v65);
          let _rhs91 = ((_896_cf6) ? (new _dafny.CodePoint('k'.codePointAt(0))) : (new _dafny.CodePoint('o'.codePointAt(0))));
          let _rhs92 = _this.f2;
          let _rhs93 = (_module.__default.safeDivisionInt(new BigNumber((_902_v64).length), new BigNumber((_904_v66).length))).isEqualTo((new BigNumber((_902_v64).length)).minus(p0));
          let _rhs94 = _897_cf5;
          let _rhs95 = _901_v63;
          let _lhs59 = globalState;
          _821_v14 = _rhs91;
          _lhs59.f1 = _rhs92;
          _896_cf6 = _rhs93;
          r0 = _rhs94;
          _901_v63 = _rhs95;
        } else {
          let _905___mcc_h13 = (_source15).cf0;
          let _906_cf0 = _905___mcc_h13;
          let _907_v67;
          _907_v67 = _dafny.MultiSet.fromElements(_this.f7, _906_cf0, (_this).f4);
          let _908_v68;
          _908_v68 = _dafny.Map.Empty.slice().updateUnsafe((_module.__default.fm15(new BigNumber(-368), _872_v52, globalState)).dtor_cf1,_dafny.MultiSet.fromElements(false, _this.f7, _906_cf0, (_this).f4));
          _907_v67 = (((_908_v68).contains((_872_v52).isLessThan(new BigNumber(755)))) ? ((_908_v68).get((_872_v52).isLessThan(new BigNumber(755)))) : (_907_v67));
          let _nw124 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _868_v48 = _nw124;
          _872_v52 = (_this).f8;
          let _909_v69;
          let _init27 = ((_910_cf0, _911_v14) => function (_912_i4) {
            return ((_910_cf0) ? (_dafny.Seq.UnicodeFromString("pi")) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-39)), ((_913_v14) => function (_914_i5) {
              return _913_v14;
            })(_911_v14))));
          })(_906_cf0, _821_v14);
          let _nw125 = Array((new BigNumber(17)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw125.length); _i0_27++) {
            _nw125[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _909_v69 = _nw125;
          let _915_v70;
          _915_v70 = _dafny.Seq.UnicodeFromString("jemmqpy");
          let _index130 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_909_v69).length));
          (_909_v69)[_index130] = _915_v70;
          let _index131 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_909_v69).length));
          let _rhs96 = ((p0).plus(_872_v52)).plus((_this).fm2(globalState));
          let _rhs97 = (new BigNumber(981)).plus((((_822_v15).dtor_cf25) ? (_872_v52) : (new BigNumber((_dafny.Seq.of(_906_cf0, _this.f7)).length))));
          let _rhs98 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wrwcm"), _915_v70);
          let _lhs60 = _909_v69;
          let _lhs61 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_909_v69).length));
          _872_v52 = _rhs96;
          _872_v52 = _rhs97;
          _lhs60[_lhs61] = _rhs98;
        }
        (_this).f7 = (_dafny.Seq.contains(_815_v9, new BigNumber(819))) || ((_this).f4);
      } else {
        let _916_v71;
        _916_v71 = new BigNumber(-829);
        _916_v71 = p0;
        let _917_v72;
        let _nw126 = new _module.C0();
        _nw126.__ctor();
        _917_v72 = _nw126;
        let _918_v73;
        let _init28 = ((_919_v14) => function (_920_i6) {
          return _919_v14;
        })(_821_v14);
        let _nw127 = Array((new BigNumber(15)).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw127.length); _i0_28++) {
          _nw127[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _918_v73 = _nw127;
        let _index132 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_918_v73).length));
        (_918_v73)[_index132] = _821_v14;
        let _index133 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_918_v73).length));
        (_918_v73)[_index133] = _821_v14;
        _916_v71 = new BigNumber(423);
        let _921_v74;
        _921_v74 = _dafny.Set.fromElements(_917_v72);
        let _922_v75;
        _922_v75 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),(_dafny.Set.fromElements(_917_v72)).Difference(_921_v74));
        _922_v75 = (_922_v75).update(new _dafny.CodePoint('l'.codePointAt(0)), (_921_v74).Union(_921_v74));
      }
      let _923_v76;
      _923_v76 = new BigNumber(100);
      let _924_v77;
      let _nw128 = new _module.C1();
      _nw128.__ctor((_this).f3, (_this).f4, _this.f2);
      _924_v77 = _nw128;
      let _925_v78;
      _925_v78 = _dafny.Map.Empty.slice().updateUnsafe(_924_v77,_924_v77);
      _923_v76 = (_dafny.ZERO).minus((_923_v76).minus((new BigNumber((_dafny.Set.fromElements(p0, new BigNumber(((_925_v78).update(_924_v77, _924_v77)).length))).length)).minus(_923_v76)));
      r0 = ((_dafny.Seq.IsProperPrefixOf(_815_v9, _dafny.Seq.of(p0, _923_v76))) ? ((_this).f4) : ((_this).f4));
      return r0;
    }
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.MultiSet.Empty;
      let r1 = false;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_this.f2).length))) {
        let _926_i0 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_926_i0)) && ((_926_i0).isLessThan(new BigNumber((_this.f2).length))))) {
          let _arr17 = _this.f2;
          _arr17[(_926_i0)] = !_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(834)), function (_927_i1) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-224)), function (_928_i2) {
              return new _dafny.CodePoint('q'.codePointAt(0));
            });
          }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(388)), function (_929_i3) {
            return _dafny.Seq.UnicodeFromString("yv");
          }));
        }
      }
      let _930_v0;
      _930_v0 = _dafny.Seq.of(false);
      let _931_v1;
      _931_v1 = _module.D3.create_DC10((_this).f4, p1, false, (_930_v0)[_module.__default.safeIndex(new BigNumber((_930_v0).length), new BigNumber((_930_v0).length))], (_this).f4);
      let _932_v2;
      _932_v2 = _dafny.Seq.of(_module.D0.create_DC1((_this).f4, true, p0));
      let _933_v3;
      _933_v3 = _module.D1.create_DC3(_932_v2);
      let _934_v4;
      _934_v4 = _dafny.Seq.UnicodeFromString("fre");
      let _935_v5;
      _935_v5 = _dafny.Map.Empty.slice().updateUnsafe(_933_v3,new BigNumber((_934_v4).length));
      let _936_v6;
      _936_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(true,(_931_v1).dtor_cf24)).update(_this.f7, p1)).length),new BigNumber(((_935_v5).update(_module.D1.create_DC3(_932_v2), (_this).f8)).length));
      (_this).f7 = (_936_v6).equals(_dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f8));
      let _937_v7;
      _937_v7 = _dafny.Seq.of(p0, new BigNumber(773), (_this).f8, (_dafny.ZERO).minus(p0), p0);
      let _938_v8;
      _938_v8 = _module.D4.create_DC11(_dafny.Seq.update(_930_v0, _module.__default.safeIndex((_this).f8, new BigNumber((_930_v0).length)), _this.f7));
      let _939_v9;
      _939_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_938_v8);
      let _rhs99 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(845)), ((_940_p0) => function (_941_i4) {
        return _940_p0;
      })(p0));
      let _rhs100 = (_939_v9).Merge(_939_v9);
      _937_v7 = _rhs99;
      _939_v9 = _rhs100;
      let _hi6 = p0;
      for (let _942_i5 = ((_this).f8).multipliedBy(p0); _942_i5.isLessThan(_hi6); _942_i5 = _942_i5.plus(_dafny.ONE)) {
        let _943_v10;
        let _nw129 = Array((new BigNumber(6)).toNumber()).fill(_dafny.MultiSet.Empty);
        _943_v10 = _nw129;
        let _nw130 = Array((_dafny.ONE).toNumber()).fill(_dafny.MultiSet.Empty);
        _943_v10 = _nw130;
        if (false) {
          let _944_v11;
          _944_v11 = new BigNumber(124);
          _944_v11 = _942_i5;
          let _945_v12;
          _945_v12 = _dafny.Set.fromElements(_944_v11, p0);
          let _946_v13;
          let _nw131 = Array((new BigNumber(19)).toNumber());
          _nw131[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(_942_i5);
          _nw131[(_dafny.ONE).toNumber()] = _937_v7;
          _nw131[(new BigNumber(2)).toNumber()] = _dafny.Seq.of((_this).f8);
          _nw131[(new BigNumber(3)).toNumber()] = _937_v7;
          _nw131[(new BigNumber(4)).toNumber()] = _dafny.Seq.of((_this).f8, _942_i5);
          _nw131[(new BigNumber(5)).toNumber()] = _937_v7;
          _nw131[(new BigNumber(6)).toNumber()] = _937_v7;
          _nw131[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_937_v7, _module.__default.safeIndex(_942_i5, new BigNumber((_937_v7).length)), _944_v11);
          _nw131[(new BigNumber(8)).toNumber()] = _937_v7;
          _nw131[(new BigNumber(9)).toNumber()] = _937_v7;
          _nw131[(new BigNumber(10)).toNumber()] = _937_v7;
          _nw131[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_937_v7, _module.__default.safeIndex((_this).f8, new BigNumber((_937_v7).length)), (_this).f8);
          _nw131[(new BigNumber(12)).toNumber()] = _937_v7;
          _nw131[(new BigNumber(13)).toNumber()] = _937_v7;
          _nw131[(new BigNumber(14)).toNumber()] = _dafny.Seq.of(p0, _944_v11, new BigNumber((_945_v12).length));
          _nw131[(new BigNumber(15)).toNumber()] = _dafny.Seq.of((_this).f8);
          _nw131[(new BigNumber(16)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(740)), function (_947_i6) {
            return new BigNumber((_dafny.Seq.of((_this).f4, (_this).f4)).length);
          });
          _nw131[(new BigNumber(17)).toNumber()] = _937_v7;
          _nw131[(new BigNumber(18)).toNumber()] = _937_v7;
          _946_v13 = _nw131;
          let _948_v14;
          _948_v14 = _dafny.Map.Empty.slice().updateUnsafe(_module.D5.create_DC14(_946_v13),_944_v11);
          _948_v14 = _948_v14;
          r1 = _this.f7;
          r1 = !_dafny.Seq.contains(_dafny.Seq.update(_934_v4, _module.__default.safeIndex((_this).f8, new BigNumber((_934_v4).length)), p1), p1);
          r1 = (_this).f4;
        } else {
          let _arr18 = _this.f2;
          let _index134 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_this.f2).length));
          _arr18[_index134] = (_this).f4;
          let _949_v15;
          _949_v15 = _module.D0.create_DC2((_this).f8, (_this).f4, (_this).f4, _942_i5);
          let _950_v16;
          _950_v16 = _dafny.Seq.of(function (_pat_let13_0) {
            return function (_951_dt__update__tmp_h0) {
              return function (_pat_let14_0) {
                return function (_952_dt__update_hcf5_h0) {
                  return _module.D0.create_DC2((_951_dt__update__tmp_h0).dtor_cf4, _952_dt__update_hcf5_h0, (_951_dt__update__tmp_h0).dtor_cf6, (_951_dt__update__tmp_h0).dtor_cf7);
                }(_pat_let14_0);
              }(false);
            }(_pat_let13_0);
          }(_949_v15));
          let _953_v17;
          _953_v17 = _dafny.Seq.of(_950_v16);
          let _arr19 = _this.f2;
          let _index135 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_this.f2).length));
          _arr19[_index135] = (_module.__default.fm25((_this).f4, globalState)).contains(((_this.f7) ? (new BigNumber((_953_v17).length)) : (p0)));
          let _954_v18;
          let _955_v19;
          let _out26;
          let _out27;
          let _outcollector11 = (_this).m4(_this.f7, globalState);
          _out26 = _outcollector11[0];
          _out27 = _outcollector11[1];
          _954_v18 = _out26;
          _955_v19 = _out27;
          _936_v6 = (_936_v6).update((_942_i5).multipliedBy(p0), _942_i5);
          let _956_v20;
          let _init29 = ((_957_p0) => function (_958_i7) {
            return _module.__default.safeModuloInt(_958_i7, _957_p0);
          })(p0);
          let _nw132 = Array((_dafny.ONE).toNumber());
          for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw132.length); _i0_29++) {
            _nw132[_i0_29] = _init29(new BigNumber(_i0_29));
          }
          _956_v20 = _nw132;
          let _index136 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_956_v20).length));
          (_956_v20)[_index136] = (_this).f8;
          let _index137 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_956_v20).length));
          (_956_v20)[_index137] = (_dafny.ZERO).minus(_942_i5);
          _955_v19 = (_this.f2)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_this.f2).length))];
        }
        let _959_v21;
        _959_v21 = new BigNumber(337);
        _959_v21 = ((_this).f8).plus(new BigNumber((_dafny.Seq.update(_934_v4, _module.__default.safeIndex(_module.__default.fm0((_this).f8, globalState), new BigNumber((_934_v4).length)), p1)).length));
        r1 = _this.f7;
      }
      let _960_v22;
      _960_v22 = new BigNumber(722);
      let _961_v23;
      _961_v23 = _dafny.Map.Empty.slice().updateUnsafe((_934_v4)[_module.__default.safeIndex((_this).f8, new BigNumber((_934_v4).length))],_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0((_this).f4),_930_v0));
      let _962_v24;
      _962_v24 = _module.D0.create_DC0(false);
      let _963_v25;
      _963_v25 = _dafny.Map.Empty.slice().updateUnsafe(_934_v4,new BigNumber((_dafny.Seq.UnicodeFromString("rhdr")).length));
      let _964_v26;
      _964_v26 = _dafny.Map.Empty.slice().updateUnsafe((((_961_v23).contains(p1)) ? ((_961_v23).get(p1)) : (_dafny.Map.Empty.slice().updateUnsafe(_962_v24,_930_v0))),(new BigNumber((_963_v25).length)).plus((_this).f8));
      let _965_v27;
      _965_v27 = _dafny.Map.Empty.slice().updateUnsafe(_962_v24,_930_v0);
      _960_v22 = (((_964_v26).contains(_965_v27)) ? ((_964_v26).get(_965_v27)) : ((_960_v22).multipliedBy((_this).f8)));
      let _966_v28;
      _966_v28 = _dafny.Seq.of(_this.f2, _this.f2);
      let _967_i8;
      _967_i8 = _dafny.ZERO;
      L3: {
        let _pat_let_tv9 = _930_v0;
        while (((_this).fm3(_this.f7, _this.f7, new BigNumber((_966_v28).length), p0, globalState)).isEqualTo(((_dafny.ZERO).minus(new BigNumber((_module.__default.fm26(_this.f7, _this.f7, function (_pat_let15_0) {
          return function (_994_dt__update__tmp_h1) {
            return function (_pat_let16_0) {
              return function (_995_dt__update_hcf33_h0) {
                return _module.D5.create_DC15(_995_dt__update_hcf33_h0, (_994_dt__update__tmp_h1).dtor_cf34);
              }(_pat_let16_0);
            }(_module.D4.create_DC11(_pat_let_tv9));
          }(_pat_let15_0);
        }(_module.D5.create_DC15(_938_v8, (_this).f4)), new BigNumber(546), globalState)).cardinality()))).minus(_module.__default.fm0((_this).f8, globalState)))) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_967_i8)) {
              break L3;
            }
            _967_i8 = (_967_i8).plus(_dafny.ONE);
            let _968_v29;
            _968_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_this.f7);
            let _969_v30;
            let _nw133 = new _module.C1();
            _nw133.__ctor((_this).f3, (((_968_v29).contains(_this.f7)) ? ((_968_v29).get(_this.f7)) : (false)), _this.f2);
            _969_v30 = _nw133;
            if (_this.f7) {
              let _970_v31;
              _970_v31 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(737)), ((_971_p1) => function (_972_i9) {
                return _971_p1;
              })(p1))).length));
              let _973_v32;
              _973_v32 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-765),_970_v31);
              let _974_v33;
              _974_v33 = _dafny.Map.Empty.slice().updateUnsafe(_934_v4,_973_v32);
              let _975_v34;
              _975_v34 = _dafny.Map.Empty.slice().updateUnsafe(_937_v7,_this.f7);
              _974_v33 = (_974_v33).update(_934_v4, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((((_975_v34).contains(_dafny.Seq.of(new BigNumber(688)))) ? ((_975_v34).get(_dafny.Seq.of(new BigNumber(688)))) : (false)),_968_v29)).length),_970_v31));
              (_this).f7 = _this.f7;
              let _976_v35;
              _976_v35 = _dafny.Set.fromElements(_this.f7);
              let _977_v36;
              _977_v36 = _dafny.MultiSet.fromElements((_this).f8, (_this).f8);
              let _978_v37;
              let _nw134 = Array((new BigNumber(3)).toNumber());
              _nw134[(_dafny.ZERO).toNumber()] = new BigNumber((_977_v36).cardinality());
              _nw134[(_dafny.ONE).toNumber()] = new BigNumber((_930_v0).length);
              _nw134[(new BigNumber(2)).toNumber()] = (_this).f8;
              _978_v37 = _nw134;
              let _979_v38;
              _979_v38 = _dafny.Map.Empty.slice().updateUnsafe(_960_v22,_this.f7);
              let _980_v39;
              _980_v39 = _dafny.Map.Empty.slice().updateUnsafe(_978_v37,_979_v38);
              let _981_v40;
              _981_v40 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_980_v39);
              let _982_v41;
              _982_v41 = _dafny.Map.Empty.slice().updateUnsafe(_this.f2,p0);
              let _rhs101 = _dafny.Seq.Concat(_934_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(987)), ((_983_p1) => function (_984_i10) {
                return _983_p1;
              })(p1)));
              let _rhs102 = !((((_981_v40).contains(_this.f7)) ? ((_981_v40).get(_this.f7)) : ((_980_v39).update(_978_v37, _979_v38)))).equals((_980_v39).Merge(_980_v39));
              let _rhs103 = (((_982_v41).contains(_this.f2)) ? ((_982_v41).get(_this.f2)) : (new BigNumber(197)));
              let _rhs104 = _976_v35;
              _934_v4 = _rhs101;
              r1 = _rhs102;
              _960_v22 = _rhs103;
              _976_v35 = _rhs104;
              _979_v38 = ((_979_v38).Merge(_979_v38)).Merge(_979_v38);
              (_this).f2 = _this.f2;
            } else {
              (_this).f7 = _this.f7;
              _960_v22 = _960_v22;
              let _985_v42;
              let _nw135 = new _module.C1();
              _nw135.__ctor((_this).f3, (_this).f4, _this.f2);
              _985_v42 = _nw135;
              let _986_v43;
              _986_v43 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm13(_this.f7, _this.f7, _this.f7, globalState),_934_v4);
              r1 = !(_986_v43).contains(_934_v4);
              let _987_v44;
              _987_v44 = _dafny.Map.Empty.slice().updateUnsafe(_960_v22,false);
              let _988_v45;
              _988_v45 = _module.D1.create_DC4(_987_v44, false, new BigNumber(135), (_this).f4);
              let _989_v46;
              _989_v46 = _dafny.Set.fromElements(_988_v45);
              _989_v46 = _dafny.Set.fromElements(_988_v45);
            }
            let _990_v47;
            _990_v47 = _dafny.Seq.of(_969_v30, _969_v30, _969_v30);
            r1 = _dafny.Seq.IsPrefixOf(_990_v47, _990_v47);
            let _991_v48;
            _991_v48 = _dafny.Set.fromElements((_this).f4);
            let _992_v49;
            _992_v49 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(713),_991_v48);
            let _993_v50;
            _993_v50 = _dafny.Seq.of(_991_v48, (((_992_v49).contains(_960_v22)) ? ((_992_v49).get(_960_v22)) : (_991_v48)));
            r1 = (((_991_v48).IsProperSubsetOf((_993_v50)[_module.__default.safeIndex(_960_v22, new BigNumber((_993_v50).length))])) ? ((!((_this).f4)) && ((_this).f4)) : ((true) || (_this.f7)));
          }
        }
      }
      let _996_v51;
      _996_v51 = _dafny.MultiSet.fromElements((_966_v28)[_module.__default.safeIndex(_960_v22, new BigNumber((_966_v28).length))], _this.f2);
      let _997_v52;
      _997_v52 = _module.D7.create_DC22(_996_v51);
      r0 = ((_996_v51).Difference((_997_v52).dtor_cf41)).Difference(_996_v51);
      r1 = _this.f7;
      return [r0, r1];
    }
    m4(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = false;
      let _998_v0;
      let _init30 = function (_999_i1) {
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("eporvvi"), _dafny.Seq.UnicodeFromString("mrfgh"));
      };
      let _nw136 = Array((new BigNumber(29)).toNumber());
      for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw136.length); _i0_30++) {
        _nw136[_i0_30] = _init30(new BigNumber(_i0_30));
      }
      _998_v0 = _nw136;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_998_v0).length))) {
        let _1000_i0 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1000_i0)) && ((_1000_i0).isLessThan(new BigNumber((_998_v0).length))))) {
          (_998_v0)[(_1000_i0)] = _dafny.Seq.Concat((((_this).f4) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-422)), function (_1001_i2) {
            return new _dafny.CodePoint('x'.codePointAt(0));
          })) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(540)), function (_1002_i3) {
            return (_module.D2.create_DC6(new _dafny.CodePoint('n'.codePointAt(0)), (_this).f8)).dtor_cf14;
          }))), _dafny.Seq.UnicodeFromString("side"));
        }
      }
      let _1003_i4;
      _1003_i4 = _dafny.ZERO;
      L4: {
        while (((_this).f8).isEqualTo((_this).f8)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1003_i4)) {
              break L4;
            }
            _1003_i4 = (_1003_i4).plus(_dafny.ONE);
            let _1004_v1;
            let _init31 = function (_1005_i5) {
              return (_1005_i5).multipliedBy((_this).f8);
            };
            let _nw137 = Array((new BigNumber(8)).toNumber());
            for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw137.length); _i0_31++) {
              _nw137[_i0_31] = _init31(new BigNumber(_i0_31));
            }
            _1004_v1 = _nw137;
            let _1006_v2;
            _1006_v2 = _dafny.Seq.UnicodeFromString("niyp");
            let _index138 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_1004_v1).length));
            (_1004_v1)[_index138] = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f2,new BigNumber((_1006_v2).length))).length)).multipliedBy((_this).f8);
            let _index139 = _module.__default.safeIndex(new BigNumber(523), new BigNumber((_1004_v1).length));
            (_1004_v1)[_index139] = (_this).f8;
            let _1007_v3;
            _1007_v3 = new _dafny.CodePoint('t'.codePointAt(0));
            let _1008_v4;
            _1008_v4 = _module.D2.create_DC6(_1007_v3, (_this).f8);
            let _1009_v5;
            _1009_v5 = _dafny.Seq.of(_1008_v4, _1008_v4, _module.D2.create_DC6(_1007_v3, (_this).fm3(_this.f7, p0, (_this).f8, new BigNumber((_1006_v2).length), globalState)));
            let _1010_v6;
            _1010_v6 = _dafny.Seq.of(_1009_v5);
            let _index140 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_1004_v1).length));
            let _index141 = _module.__default.safeIndex(new BigNumber(523), new BigNumber((_1004_v1).length));
            let _rhs105 = ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1007_v3,_1004_v1)).length)).plus((_this).f8)).multipliedBy((_this).f8);
            let _rhs106 = (_this).f8;
            let _rhs107 = _1010_v6;
            let _lhs62 = _1004_v1;
            let _lhs63 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_1004_v1).length));
            let _lhs64 = _1004_v1;
            let _lhs65 = _module.__default.safeIndex(new BigNumber(523), new BigNumber((_1004_v1).length));
            _lhs62[_lhs63] = _rhs105;
            _lhs64[_lhs65] = _rhs106;
            _1010_v6 = _rhs107;
            let _1011_v7;
            _1011_v7 = _dafny.MultiSet.fromElements(new BigNumber(818));
            let _1012_v8;
            _1012_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1011_v7).cardinality()),!((_this).f4));
            let _1013_v9;
            _1013_v9 = _dafny.Set.fromElements(_this.f7);
            let _1014_v10;
            _1014_v10 = _dafny.Map.Empty.slice().updateUnsafe((_1004_v1)[_module.__default.safeIndex(new BigNumber(413), new BigNumber((_1004_v1).length))],_1013_v9);
            let _rhs108 = !((new BigNumber((_1006_v2).length)).isLessThanOrEqualTo(new BigNumber((_1012_v8).length))) || (!(_this.f7) || (p0));
            let _rhs109 = (_this.f7) && (((_1004_v1)[_module.__default.safeIndex(new BigNumber(413), new BigNumber((_1004_v1).length))]).isEqualTo((_dafny.ZERO).minus(new BigNumber((_1014_v10).length))));
            let _lhs66 = _this;
            let _lhs67 = _this;
            _lhs66.f7 = _rhs108;
            _lhs67.f7 = _rhs109;
            let _1015_v11;
            _1015_v11 = _dafny.Set.fromElements(_1006_v2);
            let _arr20 = _this.f2;
            let _index142 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_this.f2).length));
            _arr20[_index142] = !(_1015_v11).contains(_1006_v2);
            let _arr21 = _this.f2;
            let _index143 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_this.f2).length));
            _arr21[_index143] = (_this).f4;
            let _1016_v12;
            _1016_v12 = _module.D7.create_DC24(p0);
            let _rhs110 = _1016_v12;
            let _rhs111 = (((_1012_v8).contains(_module.__default.safeModuloInt((_dafny.ZERO).minus((_1004_v1)[_module.__default.safeIndex(new BigNumber(413), new BigNumber((_1004_v1).length))]), (_1004_v1)[_module.__default.safeIndex(new BigNumber(413), new BigNumber((_1004_v1).length))]))) ? ((_1012_v8).get(_module.__default.safeModuloInt((_dafny.ZERO).minus((_1004_v1)[_module.__default.safeIndex(new BigNumber(413), new BigNumber((_1004_v1).length))]), (_1004_v1)[_module.__default.safeIndex(new BigNumber(413), new BigNumber((_1004_v1).length))]))) : ((_this.f2)[_module.__default.safeIndex(new BigNumber(9), new BigNumber((_this.f2).length))]));
            let _rhs112 = _1006_v2;
            _1016_v12 = _rhs110;
            r1 = _rhs111;
            _1006_v2 = _rhs112;
          }
        }
      }
      let _1017_v13;
      let _nw138 = new _module.C1();
      _nw138.__ctor((_this).f3, (_this.f7) === (_this.f7), _this.f2);
      _1017_v13 = _nw138;
      for (const _guard_loop_6 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_998_v0).length))) {
        let _1018_i6 = _guard_loop_6;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1018_i6)) && ((_1018_i6).isLessThan(new BigNumber((_998_v0).length))))) {
          (_998_v0)[(_1018_i6)] = ((false) ? (_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("koiruaywh"), _dafny.Seq.UnicodeFromString("m"))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(533)), function (_1019_i7) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          })));
        }
      }
      let _hi7 = (_this).f8;
      for (let _1020_i8 = (new BigNumber(556)).plus((_this).f8); _1020_i8.isLessThan(_hi7); _1020_i8 = _1020_i8.plus(_dafny.ONE)) {
        let _1021_v14;
        _1021_v14 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
        _1021_v14 = (_1021_v14).update((_this).f4, _this.f7);
        let _1022_v15;
        _1022_v15 = new BigNumber(727);
        let _1023_v16;
        let _nw139 = Array((new BigNumber(13)).toNumber()).fill(_module.D0.Default());
        _1023_v16 = _nw139;
        let _1024_v17;
        _1024_v17 = _dafny.MultiSet.fromElements(_1023_v16, _1023_v16, _1023_v16);
        _1022_v15 = new BigNumber((_1024_v17).cardinality());
        let _1025_v18;
        let _init32 = function (_1026_i9) {
          return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("i"),new BigNumber(668))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(442)), function (_1027_i10) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          }),new BigNumber(906)));
        };
        let _nw140 = Array((new BigNumber(16)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw140.length); _i0_32++) {
          _nw140[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _1025_v18 = _nw140;
        let _1028_v19;
        _1028_v19 = _dafny.Seq.UnicodeFromString("acwigaa");
        let _1029_v20;
        _1029_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1028_v19,new BigNumber(48));
        let _1030_v21;
        _1030_v21 = _module.D8.create_DC25(_1029_v20);
        let _index144 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_1025_v18).length));
        (_1025_v18)[_index144] = ((_1030_v21).dtor_cf44).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1028_v19,_1022_v15));
        let _index145 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_1025_v18).length));
        let _rhs113 = (_1029_v20).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(252)), function (_1031_i11) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }), (_dafny.ZERO).minus((_this).fm2(globalState)));
        let _rhs114 = p0;
        let _lhs68 = _1025_v18;
        let _lhs69 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_1025_v18).length));
        _lhs68[_lhs69] = _rhs113;
        r1 = _rhs114;
        let _1032_v22;
        let _nw141 = new _module.C1();
        _nw141.__ctor((_this).f3, (_this).f4, _this.f2);
        _1032_v22 = _nw141;
      }
      let _1033_v23;
      _1033_v23 = _dafny.Seq.UnicodeFromString("apes");
      let _1034_v24;
      _1034_v24 = _module.D0.create_DC2((_this).f8, (_this).f4, !((_this).f4), (_1017_v13).fm3(p0, false, new BigNumber((_1033_v23).length), (_dafny.ZERO).minus((_this).f8), globalState));
      (_this).m5((_1034_v24).dtor_cf6, p0, globalState);
      let _1035_v25;
      _1035_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1033_v23,(_this).f8);
      let _1036_v27;
      _1036_v27 = new _dafny.CodePoint('t'.codePointAt(0));
      let _1037_v28;
      _1037_v28 = _dafny.Set.fromElements(_dafny.Seq.update(_dafny.Seq.Concat(_1033_v23, _1033_v23), _module.__default.safeIndex(new BigNumber((function () {
        let _coll44 = new _dafny.Set();
        for (const _compr_44 of (_1035_v25).Keys.Elements) {
          let _1038_v26 = _compr_44;
          if ((_1035_v25).contains(_1038_v26)) {
            _coll44.add(_1038_v26);
          }
        }
        return _coll44;
      }()).length), new BigNumber((_dafny.Seq.Concat(_1033_v23, _1033_v23)).length)), _1036_v27), _1033_v23, _1033_v23, _dafny.Seq.UnicodeFromString("uwwdmx"), _dafny.Seq.Concat(_1033_v23, _1033_v23));
      r0 = _1037_v28;
      r1 = p0;
      return [r0, r1];
    }
    m5(p0, p1, globalState) {
      let _this = this;
      if (p1) {
        let _1039_v0;
        _1039_v0 = _dafny.Seq.UnicodeFromString("almht");
        _1039_v0 = _1039_v0;
        let _1040_v1;
        _1040_v1 = _module.D5.create_DC17((_this).f8);
        let _1041_v2;
        _1041_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,(_1040_v1).dtor_cf35);
        let _1042_v3;
        _1042_v3 = new BigNumber(-992);
        let _rhs115 = ((_this).f8).isLessThanOrEqualTo(_1042_v3);
        let _rhs116 = _1041_v2;
        let _rhs117 = !((_this).f8).isEqualTo(_1042_v3);
        let _rhs118 = new BigNumber(984);
        let _lhs70 = _this;
        let _lhs71 = _this;
        _lhs70.f7 = _rhs115;
        _1041_v2 = _rhs116;
        _lhs71.f7 = _rhs117;
        _1042_v3 = _rhs118;
        let _arr22 = _this.f2;
        let _index146 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_this.f2).length));
        _arr22[_index146] = (_this).fm12(globalState);
        let _1043_v4;
        _1043_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,p1);
        let _1044_v5;
        _1044_v5 = _dafny.Set.fromElements(p0, p0, (_this).f4, (((_1043_v4).contains((_this).f4)) ? ((_1043_v4).get((_this).f4)) : ((_this).f4)));
        let _1045_v6;
        _1045_v6 = _dafny.MultiSet.fromElements((_this).f8, new BigNumber((_1044_v5).length));
        let _1046_v7;
        _1046_v7 = _dafny.MultiSet.fromElements(p0, _this.f7);
        let _1047_v8;
        _1047_v8 = _dafny.Seq.of(_1046_v7);
        let _arr23 = _this.f2;
        let _index147 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_this.f2).length));
        _arr23[_index147] = !(!(((((_1045_v6).contains(_1042_v3)) ? ((_1045_v6).get(_1042_v3)) : (new BigNumber((_1047_v8).length)))).isLessThan(_1042_v3))) || (!(_this.f7) || (p0));
        if (_this.f7) {
          let _1048_v9;
          let _init33 = function (_1049_i0) {
            return _dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('s'.codePointAt(0)));
          };
          let _nw142 = Array((new BigNumber(5)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw142.length); _i0_33++) {
            _nw142[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1048_v9 = _nw142;
          let _1050_v10;
          _1050_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this.f2)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_this.f2).length))],new _dafny.CodePoint('x'.codePointAt(0)));
          let _index148 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_1048_v9).length));
          (_1048_v9)[_index148] = _1050_v10;
          let _1051_v11;
          let _nw143 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _1051_v11 = _nw143;
          let _index149 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_1051_v11).length));
          (_1051_v11)[_index149] = (_this).f8;
          let _1052_v12;
          _1052_v12 = _dafny.Set.fromElements(_module.__default.fm27(new BigNumber(60), _1042_v3, globalState));
          let _1053_v13;
          _1053_v13 = new _dafny.CodePoint('o'.codePointAt(0));
          let _index150 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_1048_v9).length));
          let _index151 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_1051_v11).length));
          let _arr24 = _this.f2;
          let _index152 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_this.f2).length));
          let _rhs119 = ((_1050_v10).Merge(_1050_v10)).update((_1052_v12).IsDisjointFrom(_1052_v12), _1053_v13);
          let _rhs120 = _this.f7;
          let _rhs121 = _1042_v3;
          let _rhs122 = (_this.f2)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_this.f2).length))];
          let _lhs72 = _1048_v9;
          let _lhs73 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_1048_v9).length));
          let _lhs74 = _this;
          let _lhs75 = _1051_v11;
          let _lhs76 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_1051_v11).length));
          let _lhs77 = _this.f2;
          let _lhs78 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_this.f2).length));
          _lhs72[_lhs73] = _rhs119;
          _lhs74.f7 = _rhs120;
          _lhs75[_lhs76] = _rhs121;
          _lhs77[_lhs78] = _rhs122;
          let _1054_v14;
          let _init34 = ((_1055_v3) => function (_1056_i1) {
            return _dafny.Seq.Concat(_dafny.Seq.of(_1055_v3, (_this).f8), _dafny.Seq.Create(_module.__default.abs(new BigNumber(122)), ((_1057_v3) => function (_1058_i2) {
              return _1057_v3;
            })(_1055_v3)));
          })(_1042_v3);
          let _nw144 = Array((new BigNumber(27)).toNumber());
          for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw144.length); _i0_34++) {
            _nw144[_i0_34] = _init34(new BigNumber(_i0_34));
          }
          _1054_v14 = _nw144;
          let _index153 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_1054_v14).length));
          (_1054_v14)[_index153] = _dafny.Seq.of((_this).f8, (_1051_v11)[_module.__default.safeIndex(new BigNumber(326), new BigNumber((_1051_v11).length))]);
          let _1059_v15;
          let _init35 = function (_1060_i3) {
            return _module.D4.create_DC11(_dafny.Seq.of(!((_this).f4)));
          };
          let _nw145 = Array((new BigNumber(28)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw145.length); _i0_35++) {
            _nw145[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1059_v15 = _nw145;
          let _index154 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_1059_v15).length));
          (_1059_v15)[_index154] = _module.__default.fm19(p0, (_1051_v11)[_module.__default.safeIndex(new BigNumber(326), new BigNumber((_1051_v11).length))], globalState);
          let _1061_v16;
          _1061_v16 = _dafny.Seq.of(p1, !(!(p0)), _this.f7, (_this).f4, true);
          let _1062_v17;
          _1062_v17 = _module.D4.create_DC11(_1061_v16);
          let _arr25 = _this.f2;
          let _index155 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_this.f2).length));
          let _index156 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_1051_v11).length));
          let _index157 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_1054_v14).length));
          let _index158 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_1059_v15).length));
          let _rhs123 = _this.f7;
          let _rhs124 = _1042_v3;
          let _rhs125 = _dafny.Seq.of(new BigNumber((_1045_v6).cardinality()), new BigNumber((_1061_v16).length), _1042_v3, (_1051_v11)[_module.__default.safeIndex(new BigNumber(326), new BigNumber((_1051_v11).length))]);
          let _rhs126 = _1062_v17;
          let _lhs79 = _this.f2;
          let _lhs80 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_this.f2).length));
          let _lhs81 = _1051_v11;
          let _lhs82 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_1051_v11).length));
          let _lhs83 = _1054_v14;
          let _lhs84 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_1054_v14).length));
          let _lhs85 = _1059_v15;
          let _lhs86 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_1059_v15).length));
          _lhs79[_lhs80] = _rhs123;
          _lhs81[_lhs82] = _rhs124;
          _lhs83[_lhs84] = _rhs125;
          _lhs85[_lhs86] = _rhs126;
          let _index159 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_1051_v11).length));
          (_1051_v11)[_index159] = _1042_v3;
          let _1063_v18;
          _1063_v18 = _module.D2.create_DC5(_1039_v0);
          let _index160 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_1051_v11).length));
          (_1051_v11)[_index160] = new BigNumber(((_1063_v18).dtor_cf13).length);
          let _1064_v19;
          _1064_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_1051_v11)[_module.__default.safeIndex(new BigNumber(326), new BigNumber((_1051_v11).length))]);
          let _index161 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_1051_v11).length));
          (_1051_v11)[_index161] = (((_1064_v19).contains((_this).f4)) ? ((_1064_v19).get((_this).f4)) : ((_1051_v11)[_module.__default.safeIndex(new BigNumber(326), new BigNumber((_1051_v11).length))]));
        } else {
          (_this).f7 = p0;
          let _1065_v20;
          _1065_v20 = new _dafny.CodePoint('i'.codePointAt(0));
          (_this).f7 = !_dafny.Seq.contains(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-936)), ((_1066_v20) => function (_1067_i4) {
            return _1066_v20;
          })(_1065_v20)), _module.__default.safeIndex((_this).f8, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-936)), ((_1068_v20) => function (_1069_i4) {
            return _1068_v20;
          })(_1065_v20))).length)), _1065_v20), _1065_v20);
          let _1070_v21;
          _1070_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1065_v20,_1042_v3);
          _1070_v21 = _1070_v21;
          let _1071_v22;
          _1071_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1042_v3,(_this.f2)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_this.f2).length))]);
          let _1072_v23;
          _1072_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this.f2)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_this.f2).length))],_module.__default.fm18(p0, (_1071_v22).update((_this).f8, p1), globalState));
          _1039_v0 = (((_1072_v23).contains((_this).f4)) ? ((_1072_v23).get((_this).f4)) : (_1039_v0));
          let _1073_v24;
          let _nw146 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _1073_v24 = _nw146;
          let _index162 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_1073_v24).length));
          (_1073_v24)[_index162] = _1042_v3;
          let _index163 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_1073_v24).length));
          (_1073_v24)[_index163] = _module.__default.safeDivisionInt(new BigNumber(688), (_this).f8);
        }
        let _1074_v25;
        let _nw147 = Array((new BigNumber(19)).toNumber()).fill([]);
        _1074_v25 = _nw147;
        let _1075_v26;
        _1075_v26 = _dafny.Seq.of(_1042_v3);
        let _1076_v27;
        let _nw148 = Array((new BigNumber(24)).toNumber());
        _nw148[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_this).f8);
        _nw148[(_dafny.ONE).toNumber()] = (_this).f8;
        _nw148[(new BigNumber(2)).toNumber()] = new BigNumber(387);
        _nw148[(new BigNumber(3)).toNumber()] = new BigNumber((_1039_v0).length);
        _nw148[(new BigNumber(4)).toNumber()] = (_this).f8;
        _nw148[(new BigNumber(5)).toNumber()] = new BigNumber((_1039_v0).length);
        _nw148[(new BigNumber(6)).toNumber()] = (_this).f8;
        _nw148[(new BigNumber(7)).toNumber()] = (_this).f8;
        _nw148[(new BigNumber(8)).toNumber()] = new BigNumber(-572);
        _nw148[(new BigNumber(9)).toNumber()] = _1042_v3;
        _nw148[(new BigNumber(10)).toNumber()] = (_this).f8;
        _nw148[(new BigNumber(11)).toNumber()] = _1042_v3;
        _nw148[(new BigNumber(12)).toNumber()] = (_this).f8;
        _nw148[(new BigNumber(13)).toNumber()] = _1042_v3;
        _nw148[(new BigNumber(14)).toNumber()] = _1042_v3;
        _nw148[(new BigNumber(15)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("soi")).length);
        _nw148[(new BigNumber(16)).toNumber()] = _1042_v3;
        _nw148[(new BigNumber(17)).toNumber()] = new BigNumber((_1075_v26).length);
        _nw148[(new BigNumber(18)).toNumber()] = (_this).f8;
        _nw148[(new BigNumber(19)).toNumber()] = (_this).f8;
        _nw148[(new BigNumber(20)).toNumber()] = (_this).f8;
        _nw148[(new BigNumber(21)).toNumber()] = (_this).fm3(p1, p0, new BigNumber(486), (_this).f8, globalState);
        _nw148[(new BigNumber(22)).toNumber()] = _1042_v3;
        _nw148[(new BigNumber(23)).toNumber()] = _1042_v3;
        _1076_v27 = _nw148;
        let _1077_v28;
        _1077_v28 = _1076_v27;
        let _index164 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_1074_v25).length));
        (_1074_v25)[_index164] = (_1077_v28);
        let _1078_v29;
        _1078_v29 = _dafny.Seq.of(_1076_v27, _1076_v27);
        let _index165 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_1074_v25).length));
        (_1074_v25)[_index165] = ((!((_this).f8).isEqualTo((_dafny.ZERO).minus(new BigNumber((_1046_v7).cardinality())))) ? ((_1078_v29)[_module.__default.safeIndex(_1042_v3, new BigNumber((_1078_v29).length))]) : (_1076_v27));
      } else {
        let _1079_v30;
        _1079_v30 = _module.D8.create_DC27(!(p1), !(p1), (_dafny.ZERO).minus((_this).f8), (_this).f8);
        let _1080_v31;
        _1080_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1079_v30,(_this).f8);
        let _1081_v32;
        _1081_v32 = _dafny.Set.fromElements((_this).f8);
        _1080_v31 = (_1080_v31).update(_1079_v30, new BigNumber((_1081_v32).length));
        let _1082_v33;
        _1082_v33 = new BigNumber(544);
        let _1083_v34;
        _1083_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,!(_this.f7));
        let _1084_v35;
        _1084_v35 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,(_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f8)));
        let _arr26 = _this.f2;
        let _index166 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_this.f2).length));
        _arr26[_index166] = (_1084_v35).contains((((_1083_v34).contains(false)) ? ((_1083_v34).get(false)) : ((_this).f4)));
        let _1085_v36;
        _1085_v36 = _dafny.MultiSet.fromElements((_this).f8);
        let _arr27 = _this.f2;
        let _index167 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_this.f2).length));
        let _rhs127 = _1082_v33;
        let _rhs128 = (_this).f4;
        let _rhs129 = (_1085_v36).IsSubsetOf(_1085_v36);
        let _rhs130 = (_this).fm12(globalState);
        let _lhs87 = _this;
        let _lhs88 = _this;
        let _lhs89 = _this.f2;
        let _lhs90 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_this.f2).length));
        _1082_v33 = _rhs127;
        _lhs87.f7 = _rhs128;
        _lhs88.f7 = _rhs129;
        _lhs89[_lhs90] = _rhs130;
        let _1086_v37;
        _1086_v37 = _dafny.Seq.of(_1082_v33, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-202)), function (_1087_i5) {
          return new _dafny.CodePoint('m'.codePointAt(0));
        })).length), (_this).f8);
        let _1088_v38;
        _1088_v38 = _dafny.Seq.UnicodeFromString("cqil");
        let _1089_v39;
        let _nw149 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _1089_v39 = _nw149;
        let _1090_v40;
        _1090_v40 = new _dafny.CodePoint('j'.codePointAt(0));
        let _1091_v41;
        _1091_v41 = _dafny.Seq.of(_1089_v39);
        let _1092_v42;
        _1092_v42 = _dafny.MultiSet.fromElements((_this).f4, (_this).f4, (_this).f4);
        let _1093_v43;
        _1093_v43 = _dafny.Seq.of(p0, true);
        let _1094_v44;
        _1094_v44 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_1093_v43)).length),(_this).f8);
        let _rhs131 = _dafny.Seq.of((_this).f8, (_dafny.ZERO).minus(new BigNumber((_1088_v38).length)), _module.__default.safeDivisionInt((_this).f8, (_this).f8));
        let _rhs132 = _dafny.Seq.Concat(_dafny.Seq.update(_1088_v38, _module.__default.safeIndex(_1082_v33, new BigNumber((_1088_v38).length)), _1090_v40), _dafny.Seq.Concat(_1088_v38, _1088_v38));
        let _rhs133 = (_1091_v41)[_module.__default.safeIndex(((_this).f8).plus(_1082_v33), new BigNumber((_1091_v41).length))];
        let _rhs134 = (new BigNumber(((_1092_v42).Union(_1092_v42)).cardinality())).plus((((_1094_v44).contains(new BigNumber((_1084_v35).length))) ? ((_1094_v44).get(new BigNumber((_1084_v35).length))) : (_1082_v33)));
        let _rhs135 = new BigNumber((_dafny.Seq.Concat(_1088_v38, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("b"), _1088_v38))).length);
        _1086_v37 = _rhs131;
        _1088_v38 = _rhs132;
        _1089_v39 = _rhs133;
        _1082_v33 = _rhs134;
        _1082_v33 = _rhs135;
        let _arr28 = _this.f2;
        let _index168 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_this.f2).length));
        _arr28[_index168] = ((_1093_v43)[_module.__default.safeIndex(_1082_v33, new BigNumber((_1093_v43).length))]) || (p0);
        if (!(p1)) {
          let _index169 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_1089_v39).length));
          (_1089_v39)[_index169] = _1082_v33;
          let _index170 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_1089_v39).length));
          let _arr29 = _this.f2;
          let _index171 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_this.f2).length));
          let _rhs136 = _module.__default.fm0(_1082_v33, globalState);
          let _rhs137 = _1094_v44;
          let _rhs138 = (_this.f2)[_module.__default.safeIndex(new BigNumber(349), new BigNumber((_this.f2).length))];
          let _lhs91 = _1089_v39;
          let _lhs92 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_1089_v39).length));
          let _lhs93 = _this.f2;
          let _lhs94 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_this.f2).length));
          _lhs91[_lhs92] = _rhs136;
          _1094_v44 = _rhs137;
          _lhs93[_lhs94] = _rhs138;
          _1084_v35 = (_1084_v35).update(!(false) || ((_this).f4), _module.__default.safeDivisionInt((_1086_v37)[_module.__default.safeIndex((_1089_v39)[_module.__default.safeIndex(new BigNumber(915), new BigNumber((_1089_v39).length))], new BigNumber((_1086_v37).length))], (_this).f8));
          _1083_v34 = (_1083_v34).update((_this).f4, (_this.f2)[_module.__default.safeIndex(new BigNumber(349), new BigNumber((_this.f2).length))]);
          let _1095_v45;
          _1095_v45 = _module.D2.create_DC5(_1088_v38);
          _1088_v38 = (_1095_v45).dtor_cf13;
          let _1096_v46;
          let _init36 = ((_1097_v39) => function (_1098_i6) {
            return _dafny.MultiSet.fromElements((_1097_v39)[_module.__default.safeIndex(new BigNumber(915), new BigNumber((_1097_v39).length))]);
          })(_1089_v39);
          let _nw150 = Array((new BigNumber(29)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw150.length); _i0_36++) {
            _nw150[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1096_v46 = _nw150;
          let _nw151 = Array((new BigNumber(20)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1096_v46 = _nw151;
        } else {
          (_this).f7 = !(!(p0) || (_this.f7));
          _1082_v33 = _1082_v33;
          _1088_v38 = _dafny.Seq.Concat(_1088_v38, _1088_v38);
          _1090_v40 = _1090_v40;
          let _1099_v47;
          _1099_v47 = _module.D8.create_DC26(p0, true, false, (_this).f8);
          let _1100_v48;
          _1100_v48 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,p0);
          let _pat_let_tv10 = _1100_v48;
          let _1101_v49;
          let _nw152 = Array((new BigNumber(11)).toNumber());
          _nw152[(_dafny.ZERO).toNumber()] = _1099_v47;
          _nw152[(_dafny.ONE).toNumber()] = function (_pat_let17_0) {
            return function (_1102_dt__update__tmp_h0) {
              return function (_pat_let18_0) {
                return function (_1103_dt__update_hcf48_h0) {
                  return function (_pat_let19_0) {
                    return function (_1104_dt__update_hcf46_h0) {
                      return _module.D8.create_DC26((_1102_dt__update__tmp_h0).dtor_cf45, _1104_dt__update_hcf46_h0, (_1102_dt__update__tmp_h0).dtor_cf47, _1103_dt__update_hcf48_h0);
                    }(_pat_let19_0);
                  }(_this.f7);
                }(_pat_let18_0);
              }(new BigNumber((_pat_let_tv10).length));
            }(_pat_let17_0);
          }(_module.D8.create_DC26(p0, (_this.f2)[_module.__default.safeIndex(new BigNumber(349), new BigNumber((_this.f2).length))], _this.f7, _1082_v33));
          _nw152[(new BigNumber(2)).toNumber()] = _1099_v47;
          _nw152[(new BigNumber(3)).toNumber()] = _module.__default.fm28((_this).f8, (_1086_v37)[_module.__default.safeIndex((_this).f8, new BigNumber((_1086_v37).length))], globalState);
          _nw152[(new BigNumber(4)).toNumber()] = _module.D8.create_DC26(p1, _this.f7, false, _1082_v33);
          _nw152[(new BigNumber(5)).toNumber()] = _1099_v47;
          _nw152[(new BigNumber(6)).toNumber()] = function (_pat_let20_0) {
            return function (_1105_dt__update__tmp_h1) {
              return function (_pat_let21_0) {
                return function (_1106_dt__update_hcf45_h0) {
                  return _module.D8.create_DC26(_1106_dt__update_hcf45_h0, (_1105_dt__update__tmp_h1).dtor_cf46, (_1105_dt__update__tmp_h1).dtor_cf47, (_1105_dt__update__tmp_h1).dtor_cf48);
                }(_pat_let21_0);
              }((_this.f2)[_module.__default.safeIndex(new BigNumber(349), new BigNumber((_this.f2).length))]);
            }(_pat_let20_0);
          }(_1099_v47);
          _nw152[(new BigNumber(7)).toNumber()] = _1099_v47;
          _nw152[(new BigNumber(8)).toNumber()] = _module.D8.create_DC26(p1, true, true, (_this).f8);
          _nw152[(new BigNumber(9)).toNumber()] = _1099_v47;
          _nw152[(new BigNumber(10)).toNumber()] = _module.D8.create_DC26(!((_this).f4), (_this).f4, (_this).fm12(globalState), _1082_v33);
          _1101_v49 = _nw152;
          let _index172 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_1101_v49).length));
          (_1101_v49)[_index172] = _1099_v47;
          let _pat_let_tv11 = _1082_v33;
          let _index173 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_1101_v49).length));
          (_1101_v49)[_index173] = function (_pat_let22_0) {
            return function (_1107_dt__update__tmp_h2) {
              return function (_pat_let23_0) {
                return function (_1108_dt__update_hcf48_h1) {
                  return _module.D8.create_DC26((_1107_dt__update__tmp_h2).dtor_cf45, (_1107_dt__update__tmp_h2).dtor_cf46, (_1107_dt__update__tmp_h2).dtor_cf47, _1108_dt__update_hcf48_h1);
                }(_pat_let23_0);
              }(_pat_let_tv11);
            }(_pat_let22_0);
          }(_module.D8.create_DC26((_this).fm12(globalState), true, p0, _1082_v33));
        }
      }
      let _rhs139 = ((_this).f8).isLessThanOrEqualTo(((_this).f8).multipliedBy((_this).f8));
      let _rhs140 = p0;
      let _lhs95 = _this;
      let _lhs96 = _this;
      _lhs95.f7 = _rhs139;
      _lhs96.f7 = _rhs140;
      let _1109_v50;
      _1109_v50 = _dafny.Seq.UnicodeFromString("jwlocp");
      let _1110_i7;
      _1110_i7 = _dafny.ZERO;
      L5: {
        while (_dafny.areEqual(_1109_v50, _1109_v50)) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1110_i7)) {
              break L5;
            }
            _1110_i7 = (_1110_i7).plus(_dafny.ONE);
            _1109_v50 = _dafny.Seq.Concat(_1109_v50, _1109_v50);
            let _1111_v51;
            _1111_v51 = _dafny.Set.fromElements((_this).f8);
            if ((_1111_v51).IsDisjointFrom(function () {
              let _coll45 = new _dafny.Set();
              for (const _compr_45 of (_module.__default.fm29(globalState)).Elements) {
                let _1112_v52 = _compr_45;
                if ((_module.__default.fm29(globalState)).contains(_1112_v52)) {
                  _coll45.add((_1112_v52).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),(_module.D0.create_DC1(true, true, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length))).dtor_cf1)).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),false)).length))).length)));
                }
              }
              return _coll45;
            }())) {
              (_this).f7 = p0;
              let _1113_v53;
              let _init37 = function (_1114_i8) {
                return (_module.D4.create_DC11(_dafny.Seq.of(_this.f7))).dtor_cf28;
              };
              let _nw153 = Array((new BigNumber(4)).toNumber());
              for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw153.length); _i0_37++) {
                _nw153[_i0_37] = _init37(new BigNumber(_i0_37));
              }
              _1113_v53 = _nw153;
              _1113_v53 = _1113_v53;
              let _1115_v54;
              _1115_v54 = new _dafny.CodePoint('c'.codePointAt(0));
              let _1116_v55;
              _1116_v55 = _dafny.Map.Empty.slice().updateUnsafe(_1115_v54,(_this).f8);
              let _1117_v56;
              _1117_v56 = _dafny.Seq.of((_this).f8);
              let _1118_v57;
              _1118_v57 = _dafny.MultiSet.fromElements(new BigNumber(-168));
              let _1119_v58;
              _1119_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_this.f7);
              _1116_v55 = (_1116_v55).update(_1115_v54, new BigNumber((((p1) ? (_dafny.MultiSet.FromArray(_1117_v56)) : ((_1118_v57).update((_this).f8, _module.__default.abs(new BigNumber((_1119_v58).length)))))).cardinality()));
              (_this).f7 = p0;
              (_this).f7 = (_this).f4;
            } else {
              let _1120_v59;
              _1120_v59 = new _dafny.CodePoint('v'.codePointAt(0));
              _1120_v59 = _1120_v59;
              let _1121_v60;
              let _nw154 = new _module.C1();
              _nw154.__ctor((_this).f3, p0, _this.f2);
              _1121_v60 = _nw154;
              let _1122_v61;
              _1122_v61 = _dafny.MultiSet.fromElements(_1121_v60);
              let _1123_v62;
              _1123_v62 = _dafny.Set.fromElements(_1122_v61);
              let _1124_v63;
              _1124_v63 = _dafny.MultiSet.fromElements(new BigNumber((_1123_v62).length));
              _1111_v51 = (_1111_v51).Intersect(function () {
                let _coll46 = new _dafny.Set();
                for (const _compr_46 of (_1124_v63).Elements) {
                  let _1125_v64 = _compr_46;
                  if ((_1124_v63).contains(_1125_v64)) {
                    _coll46.add((_1125_v64).plus((_dafny.ZERO).minus(new BigNumber(-584))));
                  }
                }
                return _coll46;
              }());
              (_this).f7 = _this.f7;
              let _1126_v65;
              let _init38 = function (_1127_i9) {
                return (_1127_i9).plus((_this).f8);
              };
              let _nw155 = Array((new BigNumber(10)).toNumber());
              for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw155.length); _i0_38++) {
                _nw155[_i0_38] = _init38(new BigNumber(_i0_38));
              }
              _1126_v65 = _nw155;
              let _1128_v66;
              _1128_v66 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wbtfe"), _1109_v50),_1126_v65);
              _1128_v66 = (_1128_v66).Merge(((_1128_v66).update(_1109_v50, _1126_v65)).update(_1109_v50, _1126_v65));
              (_this).f7 = !(((_this).f8).isLessThan(_module.__default.safeDivisionInt((_this).f8, (_this).f8)));
            }
            if (!(p0)) {
              let _1129_v67;
              _1129_v67 = new BigNumber(483);
              let _1130_v68;
              _1130_v68 = _dafny.Seq.of(_1129_v67);
              _1129_v67 = new BigNumber((_1130_v68).length);
              _1129_v67 = _module.__default.safeModuloInt(new BigNumber((_module.__default.fm29(globalState)).length), new BigNumber(471));
              let _1131_v69;
              _1131_v69 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_module.D0.create_DC1((_this).fm12(globalState), p1, _1129_v67));
              let _1132_v70;
              _1132_v70 = _module.D0.create_DC1(!(_this.f7), !(false), (_dafny.ZERO).minus(_1129_v67));
              _1131_v69 = (_1131_v69).update(_module.__default.fm0((_this).f8, globalState), _1132_v70);
              _1129_v67 = new BigNumber((_1109_v50).length);
              let _rhs141 = _1129_v67;
              let _rhs142 = (_this).f8;
              let _rhs143 = ((_this).f8).isLessThan(new BigNumber(967));
              let _lhs97 = _this;
              _1129_v67 = _rhs141;
              _1129_v67 = _rhs142;
              _lhs97.f7 = _rhs143;
            } else {
              let _1133_v71;
              _1133_v71 = new _dafny.CodePoint('q'.codePointAt(0));
              let _1134_v72;
              _1134_v72 = _dafny.MultiSet.fromElements(_1133_v71);
              let _1135_v73;
              _1135_v73 = _module.D4.create_DC13(p1, _1134_v72, !(true));
              let _pat_let_tv12 = _1133_v71;
              let _pat_let_tv13 = p0;
              let _1136_v74;
              let _nw156 = Array((new BigNumber(18)).toNumber());
              _nw156[(_dafny.ZERO).toNumber()] = _module.D4.create_DC13(false, (_dafny.MultiSet.fromElements(_1133_v71, _1133_v71, _1133_v71)).update(_1133_v71, _module.__default.abs(new BigNumber(505))), (_this).f4);
              _nw156[(_dafny.ONE).toNumber()] = _1135_v73;
              _nw156[(new BigNumber(2)).toNumber()] = _1135_v73;
              _nw156[(new BigNumber(3)).toNumber()] = _module.D4.create_DC13(p1, _1134_v72, (_this).f4);
              _nw156[(new BigNumber(4)).toNumber()] = _module.__default.fm30((_this).f8, p1, false, p1, globalState);
              _nw156[(new BigNumber(5)).toNumber()] = _module.__default.fm30((_this).f8, _this.f7, (_this).f4, _this.f7, globalState);
              _nw156[(new BigNumber(6)).toNumber()] = _1135_v73;
              _nw156[(new BigNumber(7)).toNumber()] = _1135_v73;
              _nw156[(new BigNumber(8)).toNumber()] = _1135_v73;
              _nw156[(new BigNumber(9)).toNumber()] = _1135_v73;
              _nw156[(new BigNumber(10)).toNumber()] = function (_pat_let24_0) {
                return function (_1137_dt__update__tmp_h3) {
                  return function (_pat_let25_0) {
                    return function (_1138_dt__update_hcf30_h0) {
                      return function (_pat_let26_0) {
                        return function (_1139_dt__update_hcf29_h0) {
                          return _module.D4.create_DC13(_1139_dt__update_hcf29_h0, _1138_dt__update_hcf30_h0, (_1137_dt__update__tmp_h3).dtor_cf31);
                        }(_pat_let26_0);
                      }(_pat_let_tv13);
                    }(_pat_let25_0);
                  }(_dafny.MultiSet.fromElements(_pat_let_tv12));
                }(_pat_let24_0);
              }(_module.D4.create_DC13(p1, _dafny.MultiSet.fromElements(_1133_v71), p1));
              _nw156[(new BigNumber(11)).toNumber()] = _1135_v73;
              _nw156[(new BigNumber(12)).toNumber()] = _1135_v73;
              _nw156[(new BigNumber(13)).toNumber()] = _1135_v73;
              _nw156[(new BigNumber(14)).toNumber()] = _1135_v73;
              _nw156[(new BigNumber(15)).toNumber()] = _1135_v73;
              _nw156[(new BigNumber(16)).toNumber()] = function (_pat_let27_0) {
                return function (_1140_dt__update__tmp_h4) {
                  return function (_pat_let28_0) {
                    return function (_1141_dt__update_hcf29_h1) {
                      return _module.D4.create_DC13(_1141_dt__update_hcf29_h1, (_1140_dt__update__tmp_h4).dtor_cf30, (_1140_dt__update__tmp_h4).dtor_cf31);
                    }(_pat_let28_0);
                  }((_this).f4);
                }(_pat_let27_0);
              }(_1135_v73);
              _nw156[(new BigNumber(17)).toNumber()] = _1135_v73;
              _1136_v74 = _nw156;
              let _1142_v75;
              _1142_v75 = _dafny.Set.fromElements(_1136_v74);
              _1142_v75 = _1142_v75;
              let _arr30 = _this.f2;
              let _index174 = _module.__default.safeIndex(new BigNumber(293), new BigNumber((_this.f2).length));
              _arr30[_index174] = _module.__default.fm20((_this).f4, globalState);
              let _arr31 = _this.f2;
              let _index175 = _module.__default.safeIndex(new BigNumber(293), new BigNumber((_this.f2).length));
              _arr31[_index175] = !(_this.f7) || (_this.f7);
              let _1143_v76;
              _1143_v76 = new BigNumber(482);
              let _1144_v77;
              let _init39 = ((_1145_v76) => function (_1146_i10) {
                return _module.__default.safeDivisionInt(_1146_i10, _1145_v76);
              })(_1143_v76);
              let _nw157 = Array((new BigNumber(13)).toNumber());
              for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw157.length); _i0_39++) {
                _nw157[_i0_39] = _init39(new BigNumber(_i0_39));
              }
              _1144_v77 = _nw157;
              let _1147_v78;
              _1147_v78 = _dafny.Seq.of((_this.f2)[_module.__default.safeIndex(new BigNumber(293), new BigNumber((_this.f2).length))]);
              let _1148_v79;
              _1148_v79 = _dafny.Map.Empty.slice().updateUnsafe(_1144_v77,_dafny.Seq.of(_dafny.Seq.of((_this.f2)[_module.__default.safeIndex(new BigNumber(293), new BigNumber((_this.f2).length))]), _1147_v78));
              _1143_v76 = (new BigNumber(566)).minus(new BigNumber(((((_1148_v79).contains(_1144_v77)) ? ((_1148_v79).get(_1144_v77)) : (_dafny.Seq.of(_1147_v78, _dafny.Seq.of(p0), _1147_v78)))).length));
              let _1149_v80;
              _1149_v80 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_1143_v76));
              let _arr32 = _this.f2;
              let _index176 = _module.__default.safeIndex(new BigNumber(293), new BigNumber((_this.f2).length));
              _arr32[_index176] = ((_1149_v80).Intersect(_1149_v80)).IsSubsetOf(_1149_v80);
              (_this).f7 = _this.f7;
            }
            let _1150_v81;
            _1150_v81 = _dafny.MultiSet.fromElements((_this).f8, (_this).f8);
            let _1151_v82;
            _1151_v82 = _dafny.MultiSet.fromElements(_this.f7);
            let _1152_v83;
            _1152_v83 = _dafny.Seq.of((((_1151_v82).contains(p0)) ? ((_1151_v82).get(p0)) : ((_this).f8)), (_this).f8);
            let _1153_v84;
            _1153_v84 = _module.D0.create_DC2(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), !(_this.f7), p0, (_dafny.ZERO).minus((_this).f8));
            let _1154_v85;
            _1154_v85 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,(_this).f8);
            let _1155_v86;
            _1155_v86 = _module.D8.create_DC26((_this).f4, true, (_1153_v84).dtor_cf5, (((_1154_v85).contains((_this).f8)) ? ((_1154_v85).get((_this).f8)) : ((_this).f8)));
            let _1156_v87;
            _1156_v87 = _module.D2.create_DC7(_module.__default.fm0((_this).f8, globalState), (_this).f8, _dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_1152_v83)), (_this).f8, (_this).f8);
            let _1157_v88;
            _1157_v88 = _dafny.Map.Empty.slice().updateUnsafe(_1156_v87,(_this).f8);
            let _1158_v89;
            let _nw158 = Array((new BigNumber(26)).toNumber());
            _nw158[(_dafny.ZERO).toNumber()] = (_this).f8;
            _nw158[(_dafny.ONE).toNumber()] = ((_this).f8).plus((_this).f8);
            _nw158[(new BigNumber(2)).toNumber()] = new BigNumber(-705);
            _nw158[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_1150_v81).cardinality()), (_this).f8));
            _nw158[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), function (_1159_i11) {
              return new _dafny.CodePoint('r'.codePointAt(0));
            }), _1109_v50)).length);
            _nw158[(new BigNumber(5)).toNumber()] = (_this).f8;
            _nw158[(new BigNumber(6)).toNumber()] = ((_this.f7) ? ((_this).f8) : (new BigNumber(471)));
            _nw158[(new BigNumber(7)).toNumber()] = (_this).f8;
            _nw158[(new BigNumber(8)).toNumber()] = (_this).f8;
            _nw158[(new BigNumber(9)).toNumber()] = ((_this).f8).multipliedBy((_this).f8);
            _nw158[(new BigNumber(10)).toNumber()] = new BigNumber((_1111_v51).length);
            _nw158[(new BigNumber(11)).toNumber()] = (_1152_v83)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1152_v83).length)), new BigNumber((_1152_v83).length))];
            _nw158[(new BigNumber(12)).toNumber()] = (_this).f8;
            _nw158[(new BigNumber(13)).toNumber()] = (_this).f8;
            _nw158[(new BigNumber(14)).toNumber()] = (_this).f8;
            _nw158[(new BigNumber(15)).toNumber()] = (_this).f8;
            _nw158[(new BigNumber(16)).toNumber()] = (_this).f8;
            _nw158[(new BigNumber(17)).toNumber()] = (((_1151_v82).contains(!(p1))) ? ((_1151_v82).get(!(p1))) : ((_1155_v86).dtor_cf48));
            _nw158[(new BigNumber(18)).toNumber()] = (_this).f8;
            _nw158[(new BigNumber(19)).toNumber()] = ((_this).f8).multipliedBy((_this).f8);
            _nw158[(new BigNumber(20)).toNumber()] = _module.__default.safeModuloInt((_this).f8, (_this).f8);
            _nw158[(new BigNumber(21)).toNumber()] = (_this).f8;
            _nw158[(new BigNumber(22)).toNumber()] = (_this).f8;
            _nw158[(new BigNumber(23)).toNumber()] = (_this).f8;
            _nw158[(new BigNumber(24)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1157_v88).length));
            _nw158[(new BigNumber(25)).toNumber()] = new BigNumber(75);
            _1158_v89 = _nw158;
            let _index177 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_1158_v89).length));
            (_1158_v89)[_index177] = (_this).f8;
            let _index178 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_1158_v89).length));
            (_1158_v89)[_index178] = (_this).f8;
          }
        }
      }
      if (!(true)) {
        let _1160_v90;
        let _init40 = ((_1161_p1, _1162_p0) => function (_1163_i12) {
          return _module.D3.create_DC10(_1161_p1, new _dafny.CodePoint('i'.codePointAt(0)), false, true, _1162_p0);
        })(p1, p0);
        let _nw159 = Array((new BigNumber(27)).toNumber());
        for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw159.length); _i0_40++) {
          _nw159[_i0_40] = _init40(new BigNumber(_i0_40));
        }
        _1160_v90 = _nw159;
        let _1164_v91;
        _1164_v91 = new _dafny.CodePoint('q'.codePointAt(0));
        let _1165_v92;
        _1165_v92 = _module.D3.create_DC10(p1, _1164_v91, (_this).f4, p1, p1);
        let _index179 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_1160_v90).length));
        (_1160_v90)[_index179] = _1165_v92;
        let _index180 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_1160_v90).length));
        (_1160_v90)[_index180] = _1165_v92;
        _1164_v91 = _1164_v91;
        let _1166_v93;
        let _nw160 = new _module.C1();
        _nw160.__ctor((_this).f3, p1, _this.f2);
        _1166_v93 = _nw160;
        _1109_v50 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(448)), ((_1167_v91) => function (_1168_i13) {
          return _1167_v91;
        })(_1164_v91)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(194)), ((_1169_v91) => function (_1170_i14) {
          return _1169_v91;
        })(_1164_v91)));
        let _1171_v94;
        _1171_v94 = new BigNumber(800);
        _1171_v94 = (_1171_v94).plus((_this).f8);
      } else {
        let _1172_v95;
        _1172_v95 = _dafny.Map.Empty.slice().updateUnsafe(((true) ? (_this.f7) : (_this.f7)),p0);
        _1172_v95 = (_1172_v95).update((_this).f4, p1);
        let _1173_v96;
        let _nw161 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
        _1173_v96 = _nw161;
        let _index181 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_1173_v96).length));
        (_1173_v96)[_index181] = (_this).f8;
        let _index182 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_1173_v96).length));
        (_1173_v96)[_index182] = ((_dafny.ZERO).minus(((true) ? ((_this).f8) : ((_this).f8)))).multipliedBy((_this).f8);
        let _1174_v97;
        _1174_v97 = _dafny.Seq.of((_1173_v96)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_1173_v96).length))]);
        if (!(!(_dafny.Seq.IsProperPrefixOf((((_this).f4) ? (_1174_v97) : (_1174_v97)), _dafny.Seq.of(new BigNumber((_1109_v50).length)))))) {
          let _1175_v98;
          _1175_v98 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_this.f7);
          let _1176_v99;
          _1176_v99 = _dafny.Set.fromElements(_module.__default.fm18(_this.f7, _1175_v98, globalState), ((p0) ? (_1109_v50) : (_1109_v50)), _dafny.Seq.Concat(_1109_v50, _dafny.Seq.UnicodeFromString("cspctqgt")));
          _1176_v99 = _1176_v99;
          let _arr33 = _this.f2;
          let _index183 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_this.f2).length));
          _arr33[_index183] = (((_1172_v95).contains(_this.f7)) ? ((_1172_v95).get(_this.f7)) : (_this.f7));
          let _1177_v100;
          _1177_v100 = new _dafny.CodePoint('y'.codePointAt(0));
          let _1178_v101;
          _1178_v101 = _dafny.MultiSet.fromElements(_1177_v100, _1177_v100, _1177_v100, _1177_v100, _1177_v100);
          let _1179_v102;
          _1179_v102 = _module.D4.create_DC13(p1, _1178_v101, p1);
          let _arr34 = _this.f2;
          let _index184 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_this.f2).length));
          _arr34[_index184] = (function (_pat_let29_0) {
            return function (_1180_dt__update__tmp_h5) {
              return function (_pat_let30_0) {
                return function (_1181_dt__update_hcf29_h2) {
                  return _module.D4.create_DC13(_1181_dt__update_hcf29_h2, (_1180_dt__update__tmp_h5).dtor_cf30, (_1180_dt__update__tmp_h5).dtor_cf31);
                }(_pat_let30_0);
              }(true);
            }(_pat_let29_0);
          }(_1179_v102)).dtor_cf29;
          let _1182_v103;
          _1182_v103 = _module.D8.create_DC26(p1, _this.f7, (_this).fm12(globalState), (_1173_v96)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_1173_v96).length))]);
          let _1183_v104;
          _1183_v104 = _dafny.MultiSet.fromElements(_1182_v103, _1182_v103);
          let _1184_v105;
          _1184_v105 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),(_1173_v96)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_1173_v96).length))]);
          let _arr35 = _this.f2;
          let _index185 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_this.f2).length));
          let _index186 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_1173_v96).length));
          let _index187 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_1173_v96).length));
          let _rhs144 = (_this).f4;
          let _rhs145 = (_1173_v96)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_1173_v96).length))];
          let _rhs146 = (((_1183_v104).contains(_1182_v103)) ? ((_1183_v104).get(_1182_v103)) : (((_1173_v96)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_1173_v96).length))]).minus((((_1184_v105).contains(_1177_v100)) ? ((_1184_v105).get(_1177_v100)) : (new BigNumber((_1172_v95).length))))));
          let _lhs98 = _this.f2;
          let _lhs99 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_this.f2).length));
          let _lhs100 = _1173_v96;
          let _lhs101 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_1173_v96).length));
          let _lhs102 = _1173_v96;
          let _lhs103 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_1173_v96).length));
          _lhs98[_lhs99] = _rhs144;
          _lhs100[_lhs101] = _rhs145;
          _lhs102[_lhs103] = _rhs146;
          let _1185_v106;
          let _nw162 = new _module.C0();
          _nw162.__ctor();
          _1185_v106 = _nw162;
          let _index188 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_1173_v96).length));
          (_1173_v96)[_index188] = (_this).f8;
        } else {
          let _index189 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_1173_v96).length));
          (_1173_v96)[_index189] = (((p1) ? ((_1173_v96)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_1173_v96).length))]) : (new BigNumber(487)))).multipliedBy((_this).f8);
          (_this).f7 = (_this).fm12(globalState);
          _1109_v50 = _1109_v50;
          let _1186_v107;
          let _nw163 = Array((new BigNumber(6)).toNumber()).fill([]);
          _1186_v107 = _nw163;
          _1186_v107 = _1186_v107;
          let _arr36 = _this.f2;
          let _index190 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_this.f2).length));
          _arr36[_index190] = (_this).f4;
          let _arr37 = _this.f2;
          let _index191 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_this.f2).length));
          _arr37[_index191] = !(((_this).f8).isLessThanOrEqualTo((_this).f8));
        }
        (_this).f7 = (_this).fm12(globalState);
        let _1187_v108;
        let _nw164 = new _module.C1();
        _nw164.__ctor((_this).f3, p1, _this.f2);
        _1187_v108 = _nw164;
      }
      let _1188_v109;
      let _nw165 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
      _1188_v109 = _nw165;
      let _1189_v110;
      _1189_v110 = _1188_v109;
      _1188_v109 = (_1189_v110);
      let _1190_i15;
      _1190_i15 = _dafny.ZERO;
      L6: {
        while (!(p1) || (p1)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1190_i15)) {
              break L6;
            }
            _1190_i15 = (_1190_i15).plus(_dafny.ONE);
            let _arr38 = _this.f2;
            let _index192 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_this.f2).length));
            _arr38[_index192] = ((p1) ? (!(_this.f7)) : (true));
            let _arr39 = _this.f2;
            let _index193 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_this.f2).length));
            _arr39[_index193] = _this.f7;
            let _1191_v111;
            _1191_v111 = _module.D0.create_DC0(p0);
            let _1192_v112;
            _1192_v112 = _dafny.Seq.of(_1191_v111, _1191_v111);
            _1192_v112 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1192_v112, _1192_v112), _1192_v112);
            let _1193_v113;
            _1193_v113 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,(_this.f2)[_module.__default.safeIndex(new BigNumber(263), new BigNumber((_this.f2).length))]);
            let _1194_v114;
            let _nw166 = Array((new BigNumber(11)).toNumber()).fill(false);
            _1194_v114 = _nw166;
            let _1195_v115;
            let _nw167 = new _module.C1();
            _nw167.__ctor((_this).f3, (_1193_v113).contains(_this.f7), _1194_v114);
            _1195_v115 = _nw167;
            _1195_v115 = _1195_v115;
            let _1196_v116;
            _1196_v116 = _dafny.Set.fromElements((_1195_v115).f4);
            let _1197_v117;
            _1197_v117 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_dafny.Seq.of(new BigNumber((_1196_v116).length)));
            let _1198_v118;
            _1198_v118 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f8),(_this).f8);
            let _1199_v119;
            _1199_v119 = _dafny.Seq.of(p0);
            let _arr40 = _this.f2;
            let _index194 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_this.f2).length));
            _arr40[_index194] = _dafny.Seq.IsProperPrefixOf((((_1197_v117).contains(new BigNumber((_1198_v118).length))) ? ((_1197_v117).get(new BigNumber((_1198_v118).length))) : (_dafny.Seq.of((_this).f8))), _module.__default.fm31((_this).f8, !((_1195_v115).f4), (_this).f8, (_1199_v119)[_module.__default.safeIndex((_this).f8, new BigNumber((_1199_v119).length))], globalState));
          }
        }
      }
      return;
    }
    get f8() {
      let _this = this;
      return _this._f8;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f2 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f2() {
      let _this = this;
      return _this._f2;
    };
    set f2(value) {
      let _this = this;
      _this._f2 = value;
    };
    __ctor(f2) {
      let _this = this;
      (_this)._f2 = f2;
      return;
    }
    fm1(p0, p1, globalState) {
      let _this = this;
      return (((false) ? (function () {
        let _coll47 = new _dafny.Map();
        for (const _compr_47 of (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-814)),new BigNumber(241))).Keys.Elements) {
          let _1200_v0 = _compr_47;
          if ((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-814)),new BigNumber(241))).contains(_1200_v0)) {
            _coll47.push([_module.__default.safeDivisionInt(_1200_v0, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(195)), function (_1201_i0) {
              return new BigNumber(92);
            })).length))).length)),false]);
          }
        }
        return _coll47;
      }()) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(929), (_dafny.ZERO).minus(new BigNumber(-512)), new BigNumber(856))).length),true)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("hyrhotgv")).length),false));
    };
    fm2(globalState) {
      let _this = this;
      return new BigNumber(859);
    };
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let _1202_v0;
      let _nw168 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
      _1202_v0 = _nw168;
      let _1203_v1;
      _1203_v1 = _dafny.Map.Empty.slice().updateUnsafe(_1202_v0,p0);
      let _1204_v2;
      let _nw169 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
      _1204_v2 = _nw169;
      let _1205_v3;
      _1205_v3 = true;
      let _1206_v4;
      let _nw170 = new _module.C2();
      _nw170.__ctor(true, (((_1203_v1).contains(_1202_v0)) ? ((_1203_v1).get(_1202_v0)) : (p0)), _1204_v2, _1205_v3, _this.f2);
      _1206_v4 = _nw170;
      _1206_v4 = _1206_v4;
      r0 = _1205_v3;
      let _1207_v5;
      _1207_v5 = _dafny.Seq.UnicodeFromString("akmdxvj");
      let _1208_v6;
      _1208_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1207_v5);
      let _1209_v7;
      _1209_v7 = new _dafny.CodePoint('c'.codePointAt(0));
      if (!(_dafny.Seq.IsPrefixOf(((false) ? (_1207_v5) : ((((_1208_v6).contains(p0)) ? ((_1208_v6).get(p0)) : (_1207_v5)))), _dafny.Seq.update(_dafny.Seq.UnicodeFromString("ebwdtouqi"), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.UnicodeFromString("ebwdtouqi")).length)), _1209_v7)))) {
        let _1210_v8;
        let _nw171 = Array((new BigNumber(13)).toNumber());
        _nw171[(_dafny.ZERO).toNumber()] = _1207_v5;
        _nw171[(_dafny.ONE).toNumber()] = _1207_v5;
        _nw171[(new BigNumber(2)).toNumber()] = _1207_v5;
        _nw171[(new BigNumber(3)).toNumber()] = _1207_v5;
        _nw171[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-257)), ((_1211_v7) => function (_1212_i0) {
          return _1211_v7;
        })(_1209_v7));
        _nw171[(new BigNumber(5)).toNumber()] = _1207_v5;
        _nw171[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(_1209_v7, _module.__default.fm22(_1207_v5, globalState));
        _nw171[(new BigNumber(7)).toNumber()] = _1207_v5;
        _nw171[(new BigNumber(8)).toNumber()] = _1207_v5;
        _nw171[(new BigNumber(9)).toNumber()] = _1207_v5;
        _nw171[(new BigNumber(10)).toNumber()] = _dafny.Seq.of(_1209_v7, _1209_v7);
        _nw171[(new BigNumber(11)).toNumber()] = _dafny.Seq.of(_1209_v7);
        _nw171[(new BigNumber(12)).toNumber()] = _1207_v5;
        _1210_v8 = _nw171;
        let _1213_v9;
        _1213_v9 = _dafny.Seq.of(p0);
        let _1214_v10;
        _1214_v10 = _dafny.Map.Empty.slice().updateUnsafe((_1213_v9)[_module.__default.safeIndex(p0, new BigNumber((_1213_v9).length))],p0);
        let _1215_v12;
        let _nw172 = Array((new BigNumber(7)).toNumber());
        _nw172[(_dafny.ZERO).toNumber()] = _1214_v10;
        _nw172[(_dafny.ONE).toNumber()] = _1214_v10;
        _nw172[(new BigNumber(2)).toNumber()] = _1214_v10;
        _nw172[(new BigNumber(3)).toNumber()] = _1214_v10;
        _nw172[(new BigNumber(4)).toNumber()] = _1214_v10;
        _nw172[(new BigNumber(5)).toNumber()] = (function () {
          let _coll48 = new _dafny.Map();
          for (const _compr_48 of (_1213_v9).Elements) {
            let _1216_v11 = _compr_48;
            if (_dafny.Seq.contains(_1213_v9, _1216_v11)) {
              _coll48.push([(_1216_v11).plus(p0),p0]);
            }
          }
          return _coll48;
        }()).update(p0, new BigNumber((_1213_v9).length));
        _nw172[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        _1215_v12 = _nw172;
        let _1217_v13;
        let _1218_v14;
        let _out28;
        let _out29;
        let _outcollector12 = (_1206_v4).m1(p0, new _dafny.CodePoint('b'.codePointAt(0)), _1210_v8, _1215_v12, globalState);
        _out28 = _outcollector12[0];
        _out29 = _outcollector12[1];
        _1217_v13 = _out28;
        _1218_v14 = _out29;
        let _1219_v15;
        _1219_v15 = _dafny.Seq.of(_1218_v14);
        _1205_v3 = ((_dafny.Seq.IsProperPrefixOf(_1219_v15, _1219_v15)) ? (false) : ((new BigNumber((_dafny.Set.fromElements(_1209_v7)).length)).isLessThanOrEqualTo(p0)));
        r0 = _1205_v3;
        if (_1205_v3) {
          let _1220_v16;
          _1220_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1218_v14,_1210_v8);
          _1220_v16 = (_1220_v16).update(_1218_v14, _1210_v8);
          let _1221_v17;
          let _nw173 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.of());
          _1221_v17 = _nw173;
          let _index195 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_1221_v17).length));
          (_1221_v17)[_index195] = _module.__default.fm31((_1213_v9)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1207_v5).length)), new BigNumber((_1213_v9).length))], _1205_v3, (_dafny.ZERO).minus(p0), _module.__default.fm20(!(_1205_v3), globalState), globalState);
          let _index196 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_1221_v17).length));
          (_1221_v17)[_index196] = _1213_v9;
          let _1222_v18;
          _1222_v18 = _dafny.Map.Empty.slice().updateUnsafe(!(_1205_v3),_1205_v3);
          let _1223_v19;
          _1223_v19 = _module.D8.create_DC26(_1218_v14, _1218_v14, _1205_v3, new BigNumber((_1222_v18).length));
          _1205_v3 = (_1223_v19).dtor_cf45;
          _1218_v14 = _1218_v14;
          let _1224_v20;
          let _nw174 = Array((new BigNumber(10)).toNumber()).fill([]);
          _1224_v20 = _nw174;
          let _index197 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_1224_v20).length));
          (_1224_v20)[_index197] = _1202_v0;
          let _index198 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_1224_v20).length));
          (_1224_v20)[_index198] = _1202_v0;
        } else {
          let _1225_v21;
          _1225_v21 = new BigNumber(-152);
          _1225_v21 = _1225_v21;
          let _1226_v22;
          let _nw175 = new _module.C1();
          _nw175.__ctor(_1204_v2, false, _this.f2);
          _1226_v22 = _nw175;
          let _1227_v23;
          _1227_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1226_v22,(_1225_v21).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(382)), ((_1228_v7) => function (_1229_i1) {
            return _1228_v7;
          })(_1209_v7))).length)));
          let _1230_v24;
          _1230_v24 = _module.D10.create_DC31(_1226_v22);
          _1227_v23 = (_1227_v23).update((_1230_v24).dtor_cf57, (_module.__default.fm20(_1205_v3, globalState)) && (_1218_v14));
          r0 = false;
          let _1231_v25;
          let _1232_v26;
          let _out30;
          let _out31;
          let _outcollector13 = (_1206_v4).m1(_module.__default.safeDivisionInt(p0, p0), _1209_v7, _1210_v8, _1215_v12, globalState);
          _out30 = _outcollector13[0];
          _out31 = _outcollector13[1];
          _1231_v25 = _out30;
          _1232_v26 = _out31;
          let _rhs147 = _1205_v3;
          let _rhs148 = (p0).multipliedBy(_module.__default.safeDivisionInt((_dafny.ZERO).minus(_1225_v21), p0));
          _1232_v26 = _rhs147;
          _1225_v21 = _rhs148;
        }
        let _1233_v27;
        _1233_v27 = _dafny.Set.fromElements(p0);
        let _1234_v28;
        _1234_v28 = _dafny.Seq.of(_1233_v27);
        let _1235_v29;
        _1235_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1207_v5).length),_1234_v28);
        _1235_v29 = (_1235_v29).update(p0, _dafny.Seq.Concat(_dafny.Seq.of(_1233_v27), _dafny.Seq.of(_1233_v27, _1233_v27)));
      } else {
        let _1236_v30;
        let _nw176 = Array((new BigNumber(28)).toNumber());
        _nw176[(_dafny.ZERO).toNumber()] = _1209_v7;
        _nw176[(_dafny.ONE).toNumber()] = _1209_v7;
        _nw176[(new BigNumber(2)).toNumber()] = (_1207_v5)[_module.__default.safeIndex(p0, new BigNumber((_1207_v5).length))];
        _nw176[(new BigNumber(3)).toNumber()] = ((true) ? (_1209_v7) : (_1209_v7));
        _nw176[(new BigNumber(4)).toNumber()] = (_1207_v5)[_module.__default.safeIndex(p0, new BigNumber((_1207_v5).length))];
        _nw176[(new BigNumber(5)).toNumber()] = _1209_v7;
        _nw176[(new BigNumber(6)).toNumber()] = _1209_v7;
        _nw176[(new BigNumber(7)).toNumber()] = _1209_v7;
        _nw176[(new BigNumber(8)).toNumber()] = _1209_v7;
        _nw176[(new BigNumber(9)).toNumber()] = _1209_v7;
        _nw176[(new BigNumber(10)).toNumber()] = _1209_v7;
        _nw176[(new BigNumber(11)).toNumber()] = _1209_v7;
        _nw176[(new BigNumber(12)).toNumber()] = _1209_v7;
        _nw176[(new BigNumber(13)).toNumber()] = _1209_v7;
        _nw176[(new BigNumber(14)).toNumber()] = _1209_v7;
        _nw176[(new BigNumber(15)).toNumber()] = _1209_v7;
        _nw176[(new BigNumber(16)).toNumber()] = _1209_v7;
        _nw176[(new BigNumber(17)).toNumber()] = ((!(_1205_v3)) ? (_1209_v7) : (_1209_v7));
        _nw176[(new BigNumber(18)).toNumber()] = _1209_v7;
        _nw176[(new BigNumber(19)).toNumber()] = _1209_v7;
        _nw176[(new BigNumber(20)).toNumber()] = new _dafny.CodePoint('d'.codePointAt(0));
        _nw176[(new BigNumber(21)).toNumber()] = _1209_v7;
        _nw176[(new BigNumber(22)).toNumber()] = _module.__default.fm22(_1207_v5, globalState);
        _nw176[(new BigNumber(23)).toNumber()] = _1209_v7;
        _nw176[(new BigNumber(24)).toNumber()] = _1209_v7;
        _nw176[(new BigNumber(25)).toNumber()] = _1209_v7;
        _nw176[(new BigNumber(26)).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
        _nw176[(new BigNumber(27)).toNumber()] = _1209_v7;
        _1236_v30 = _nw176;
        let _nw177 = Array((new BigNumber(27)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1236_v30 = _nw177;
        let _1237_v31;
        _1237_v31 = new BigNumber(-732);
        _1237_v31 = _1237_v31;
        let _arr41 = _this.f2;
        let _index199 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_this.f2).length));
        _arr41[_index199] = _module.__default.fm20(_1205_v3, globalState);
        let _arr42 = _this.f2;
        let _index200 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_this.f2).length));
        _arr42[_index200] = _1205_v3;
        _1207_v5 = _1207_v5;
        let _1238_v32;
        _1238_v32 = _dafny.Seq.of(_1237_v31);
        let _index201 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_1202_v0).length));
        (_1202_v0)[_index201] = (_module.D8.create_DC28(_1205_v3, new BigNumber((_1238_v32).length))).dtor_cf54;
        let _index202 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_1202_v0).length));
        (_1202_v0)[_index202] = (_1238_v32)[_module.__default.safeIndex(new BigNumber(682), new BigNumber((_1238_v32).length))];
      }
      let _1239_v33;
      _1239_v33 = _dafny.Seq.of((true) && (_1205_v3), _1205_v3, !(_1205_v3));
      let _1240_v34;
      _1240_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1205_v3,_1205_v3);
      _1239_v33 = _module.__default.fm32(new BigNumber((_dafny.Seq.of(_1202_v0)).length), !(_1205_v3), _1240_v34, globalState);
      let _1241_v35;
      _1241_v35 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-184)), ((_1242_v7) => function (_1243_i2) {
        return _1242_v7;
      })(_1209_v7)),p0);
      let _1244_v36;
      _1244_v36 = _module.D8.create_DC25(_1241_v35);
      let _source16 = _1244_v36;
      if (_source16.is_DC26) {
        let _1245___mcc_h0 = (_source16).cf45;
        let _1246___mcc_h1 = (_source16).cf46;
        let _1247___mcc_h2 = (_source16).cf47;
        let _1248___mcc_h3 = (_source16).cf48;
        let _1249_cf48 = _1248___mcc_h3;
        let _1250_cf47 = _1247___mcc_h2;
        let _1251_cf46 = _1246___mcc_h1;
        let _1252_cf45 = _1245___mcc_h0;
        let _arr43 = _this.f2;
        let _index203 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_this.f2).length));
        _arr43[_index203] = _1251_cf46;
        let _arr44 = _this.f2;
        let _index204 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_this.f2).length));
        _arr44[_index204] = !(((new BigNumber((_1207_v5).length)).isLessThan(new BigNumber((_1207_v5).length))) || (_1250_cf47));
        let _1253_v37;
        _1253_v37 = _dafny.Seq.of(new BigNumber((_1207_v5).length), new BigNumber((_1207_v5).length), p0);
        let _1254_v38;
        _1254_v38 = _dafny.MultiSet.fromElements(p0, (_this).fm2(globalState));
        let _rhs149 = (((_1205_v3) ? (false) : (false))) && ((_1254_v38).IsProperSubsetOf(_dafny.MultiSet.FromArray(_1253_v37)));
        let _rhs150 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_1207_v5, _1207_v5), _1207_v5), _module.__default.safeIndex(_1249_cf48, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1207_v5, _1207_v5), _1207_v5)).length)), _1209_v7);
        let _rhs151 = _1252_cf45;
        _1205_v3 = _rhs149;
        _1207_v5 = _rhs150;
        _1251_cf46 = _rhs151;
        let _index205 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_1202_v0).length));
        (_1202_v0)[_index205] = _1249_cf48;
        let _index206 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_1202_v0).length));
        (_1202_v0)[_index206] = p0;
        let _1255_v39;
        _1255_v39 = _module.D6.create_DC20(_1251_cf46, _1251_cf46);
        let _1256_v40;
        _1256_v40 = _module.D6.create_DC21(_1255_v39);
        let _1257_v41;
        _1257_v41 = _module.D6.create_DC21(_1255_v39);
        let _pat_let_tv14 = _1205_v3;
        let _pat_let_tv15 = _1252_cf45;
        let _1258_v42;
        _1258_v42 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm14((_1202_v0)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_1202_v0).length))], _1249_cf48, _1251_cf46, (_this.f2)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_this.f2).length))], globalState),function (_pat_let31_0) {
          return function (_1259_dt__update__tmp_h0) {
            return function (_pat_let32_0) {
              return function (_1260_dt__update_hcf40_h0) {
                return _module.D6.create_DC21(_1260_dt__update_hcf40_h0);
              }(_pat_let32_0);
            }(_module.D6.create_DC20(_pat_let_tv14, _pat_let_tv15));
          }(_pat_let31_0);
        }(_1257_v41));
        let _1261_v43;
        _1261_v43 = _module.D0.create_DC0((_this.f2)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_this.f2).length))]);
        _1258_v42 = (_1258_v42).update(_1261_v43, _module.D6.create_DC21(_1256_v40));
      } else if (_source16.is_DC27) {
        let _1262___mcc_h4 = (_source16).cf49;
        let _1263___mcc_h5 = (_source16).cf50;
        let _1264___mcc_h6 = (_source16).cf51;
        let _1265___mcc_h7 = (_source16).cf52;
        let _1266_cf52 = _1265___mcc_h7;
        let _1267_cf51 = _1264___mcc_h6;
        let _1268_cf50 = _1263___mcc_h5;
        let _1269_cf49 = _1262___mcc_h4;
        let _1270_v44;
        let _nw178 = Array((new BigNumber(28)).toNumber());
        _nw178[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-36)), ((_1271_v7) => function (_1272_i3) {
          return _1271_v7;
        })(_1209_v7));
        _nw178[(_dafny.ONE).toNumber()] = _1207_v5;
        _nw178[(new BigNumber(2)).toNumber()] = _1207_v5;
        _nw178[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_1207_v5, _module.__default.safeIndex(p0, new BigNumber((_1207_v5).length)), _1209_v7);
        _nw178[(new BigNumber(4)).toNumber()] = _1207_v5;
        _nw178[(new BigNumber(5)).toNumber()] = _1207_v5;
        _nw178[(new BigNumber(6)).toNumber()] = _1207_v5;
        _nw178[(new BigNumber(7)).toNumber()] = _1207_v5;
        _nw178[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_1207_v5, _module.__default.safeIndex(p0, new BigNumber((_1207_v5).length)), _1209_v7);
        _nw178[(new BigNumber(9)).toNumber()] = _1207_v5;
        _nw178[(new BigNumber(10)).toNumber()] = _1207_v5;
        _nw178[(new BigNumber(11)).toNumber()] = _1207_v5;
        _nw178[(new BigNumber(12)).toNumber()] = _1207_v5;
        _nw178[(new BigNumber(13)).toNumber()] = _1207_v5;
        _nw178[(new BigNumber(14)).toNumber()] = _1207_v5;
        _nw178[(new BigNumber(15)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(152)), ((_1273_v7) => function (_1274_i4) {
          return _1273_v7;
        })(_1209_v7));
        _nw178[(new BigNumber(16)).toNumber()] = _1207_v5;
        _nw178[(new BigNumber(17)).toNumber()] = _1207_v5;
        _nw178[(new BigNumber(18)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(481)), function (_1275_i5) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        });
        _nw178[(new BigNumber(19)).toNumber()] = _1207_v5;
        _nw178[(new BigNumber(20)).toNumber()] = _1207_v5;
        _nw178[(new BigNumber(21)).toNumber()] = _module.__default.fm13(false, _1205_v3, true, globalState);
        _nw178[(new BigNumber(22)).toNumber()] = _1207_v5;
        _nw178[(new BigNumber(23)).toNumber()] = _1207_v5;
        _nw178[(new BigNumber(24)).toNumber()] = _1207_v5;
        _nw178[(new BigNumber(25)).toNumber()] = _1207_v5;
        _nw178[(new BigNumber(26)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(890)), ((_1276_v7) => function (_1277_i6) {
          return _1276_v7;
        })(_1209_v7));
        _nw178[(new BigNumber(27)).toNumber()] = _dafny.Seq.of(_1209_v7);
        _1270_v44 = _nw178;
        let _1278_v45;
        let _nw179 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
        _1278_v45 = _nw179;
        let _1279_v46;
        let _1280_v47;
        let _out32;
        let _out33;
        let _outcollector14 = (_1206_v4).m1(_1266_cf52, _1209_v7, _1270_v44, _1278_v45, globalState);
        _out32 = _outcollector14[0];
        _out33 = _outcollector14[1];
        _1279_v46 = _out32;
        _1280_v47 = _out33;
        let _1281_v48;
        _1281_v48 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("gran")).length),_1205_v3);
        _1207_v5 = _module.__default.fm18(_1205_v3, (_1281_v48).Merge(_1281_v48), globalState);
        let _1282_v49;
        let _nw180 = Array((new BigNumber(11)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1282_v49 = _nw180;
        let _index207 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_1282_v49).length));
        (_1282_v49)[_index207] = new _dafny.CodePoint('x'.codePointAt(0));
        let _index208 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_1282_v49).length));
        (_1282_v49)[_index208] = (((new BigNumber(539)).isLessThan((_dafny.ZERO).minus(p0))) ? (_1209_v7) : ((_1207_v5)[_module.__default.safeIndex(p0, new BigNumber((_1207_v5).length))]));
        _1266_cf52 = _module.__default.safeModuloInt(_1266_cf52, _1267_cf51);
      } else if (_source16.is_DC28) {
        let _1283___mcc_h8 = (_source16).cf53;
        let _1284___mcc_h9 = (_source16).cf54;
        let _1285_cf54 = _1284___mcc_h9;
        let _1286_cf53 = _1283___mcc_h8;
        r0 = _1205_v3;
        let _1287_v50;
        let _nw181 = new _module.C0();
        _nw181.__ctor();
        _1287_v50 = _nw181;
        _1285_cf54 = _module.__default.safeDivisionInt(_1285_cf54, (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(415)), function (_1288_i7) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        })).length)).multipliedBy(p0)));
        let _1289_v51;
        _1289_v51 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(31)), ((_1290_p0) => function (_1291_i8) {
          return _1290_p0;
        })(p0)));
        let _1292_v52;
        _1292_v52 = _dafny.Seq.of(_1285_cf54, _1285_cf54);
        let _1293_v53;
        _1293_v53 = _dafny.Map.Empty.slice().updateUnsafe(_1286_cf53,(_1292_v52)[_module.__default.safeIndex(new BigNumber(752), new BigNumber((_1292_v52).length))]);
        let _1294_v54;
        _1294_v54 = _module.D11.create_DC34(_1293_v53);
        let _1295_v55;
        _1295_v55 = _dafny.Seq.of(new BigNumber(((_1294_v54).dtor_cf62).length), new BigNumber((_1292_v52).length));
        let _1296_v56;
        _1296_v56 = _dafny.Seq.of(_1295_v55);
        let _1297_v57;
        _1297_v57 = _dafny.MultiSet.fromElements(_1285_cf54);
        let _1298_v58;
        _1298_v58 = _module.D10.create_DC33(((_dafny.MultiSet.FromArray(_1296_v56)).update(_1295_v55, _module.__default.abs(new BigNumber((_1297_v57).cardinality())))).IsProperSubsetOf(_1289_v51));
        let _source17 = _1298_v58;
        if (_source17.is_DC32) {
          let _1299___mcc_h12 = (_source17).cf58;
          let _1300___mcc_h13 = (_source17).cf59;
          let _1301___mcc_h14 = (_source17).cf60;
          let _1302_cf60 = _1301___mcc_h14;
          let _1303_cf59 = _1300___mcc_h13;
          let _1304_cf58 = _1299___mcc_h12;
          let _index209 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_1202_v0).length));
          (_1202_v0)[_index209] = _1302_cf60;
          let _index210 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_1202_v0).length));
          (_1202_v0)[_index210] = _1302_cf60;
          _1302_cf60 = _1302_cf60;
          let _1305_v59;
          _1305_v59 = _module.D4.create_DC13(_1205_v3, _dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0))), _1205_v3);
          let _1306_v60;
          _1306_v60 = _module.D4.create_DC13(_1205_v3, (_1305_v59).dtor_cf30, _1205_v3);
          let _1307_v61;
          let _nw182 = new _module.C1();
          _nw182.__ctor(_1204_v2, ((_1305_v59).dtor_cf29) || (_1286_cf53), _1206_v4.f2);
          _1307_v61 = _nw182;
          let _1308_v62;
          let _nw183 = new _module.C0();
          _nw183.__ctor();
          _1308_v62 = _nw183;
        } else if (_source17.is_DC33) {
          let _1309___mcc_h15 = (_source17).cf61;
          let _1310_cf61 = _1309___mcc_h15;
          let _1311_v63;
          _1311_v63 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-460),_1205_v3);
          let _1312_v64;
          _1312_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1311_v63,_1298_v58);
          _1312_v64 = (_1312_v64).update((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1239_v33).length),_1286_cf53)).Merge(_1311_v63), ((_1310_cf61) ? (_1298_v58) : (_1298_v58)));
          let _index211 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_1202_v0).length));
          (_1202_v0)[_index211] = p0;
          let _index212 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_1202_v0).length));
          (_1202_v0)[_index212] = (p0).minus(p0);
          let _index213 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_1202_v0).length));
          (_1202_v0)[_index213] = ((_1202_v0)[_module.__default.safeIndex(new BigNumber(342), new BigNumber((_1202_v0).length))]).plus(p0);
          let _index214 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_1202_v0).length));
          (_1202_v0)[_index214] = (_1202_v0)[_module.__default.safeIndex(new BigNumber(342), new BigNumber((_1202_v0).length))];
        } else {
          let _1313___mcc_h16 = (_source17).cf57;
          let _1314_cf57 = _1313___mcc_h16;
          _1285_cf54 = (_dafny.ZERO).minus((p0).minus(new BigNumber(768)));
          let _1315_v65;
          _1315_v65 = _module.D10.create_DC31(((false) ? (_1314_cf57) : (_1314_cf57)));
          let _rhs152 = ((new BigNumber((_1293_v53).length)).minus(new BigNumber((_1207_v5).length))).plus(p0);
          let _rhs153 = _1315_v65;
          let _rhs154 = p0;
          _1285_cf54 = _rhs152;
          _1315_v65 = _rhs153;
          _1285_cf54 = _rhs154;
          let _1316_v66;
          _1316_v66 = _module.D6.create_DC20(_1286_cf53, !(_1286_cf53));
          let _pat_let_tv16 = _1205_v3;
          let _1317_v67;
          let _nw184 = Array((new BigNumber(8)).toNumber());
          _nw184[(_dafny.ZERO).toNumber()] = _1316_v66;
          _nw184[(_dafny.ONE).toNumber()] = _1316_v66;
          _nw184[(new BigNumber(2)).toNumber()] = _1316_v66;
          _nw184[(new BigNumber(3)).toNumber()] = function (_pat_let33_0) {
            return function (_1318_dt__update__tmp_h1) {
              return function (_pat_let34_0) {
                return function (_1319_dt__update_hcf39_h0) {
                  return _module.D6.create_DC20((_1318_dt__update__tmp_h1).dtor_cf38, _1319_dt__update_hcf39_h0);
                }(_pat_let34_0);
              }(_pat_let_tv16);
            }(_pat_let33_0);
          }(_1316_v66);
          _nw184[(new BigNumber(4)).toNumber()] = _1316_v66;
          _nw184[(new BigNumber(5)).toNumber()] = _1316_v66;
          _nw184[(new BigNumber(6)).toNumber()] = _1316_v66;
          _nw184[(new BigNumber(7)).toNumber()] = _module.D6.create_DC20(_1286_cf53, _1205_v3);
          _1317_v67 = _nw184;
          let _index215 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_1317_v67).length));
          (_1317_v67)[_index215] = _1316_v66;
          let _index216 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_1317_v67).length));
          (_1317_v67)[_index216] = _1316_v66;
          let _1320_v68;
          _1320_v68 = _module.D5.create_DC16();
          let _1321_v69;
          _1321_v69 = _dafny.Seq.of(_1320_v68, _1320_v68);
          let _1322_v70;
          _1322_v70 = _dafny.Seq.of(_dafny.Seq.update(_1207_v5, _module.__default.safeIndex(new BigNumber((_1321_v69).length), new BigNumber((_1207_v5).length)), _1209_v7), _1207_v5, _1207_v5);
          let _1323_v71;
          let _nw185 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Map.Empty);
          _1323_v71 = _nw185;
          let _1324_v72;
          _1324_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1206_v4.f2,p0);
          let _index217 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_1323_v71).length));
          (_1323_v71)[_index217] = (_dafny.Map.Empty.slice().updateUnsafe(_1206_v4.f2,_1285_cf54)).Merge(_1324_v72);
          let _1325_v73;
          _1325_v73 = _dafny.MultiSet.fromElements(_1205_v3, !(_1205_v3), (_1239_v33)[_module.__default.safeIndex(p0, new BigNumber((_1239_v33).length))]);
          let _index218 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_1323_v71).length));
          let _rhs155 = (p0).plus((new BigNumber((_1292_v52).length)).minus((_dafny.ZERO).minus(p0)));
          let _rhs156 = _dafny.Seq.Concat(_1322_v70, _dafny.Seq.Create(_module.__default.abs(new BigNumber(854)), ((_1326_v5) => function (_1327_i9) {
            return _1326_v5;
          })(_1207_v5)));
          let _rhs157 = ((_dafny.MultiSet.fromElements(!(false), (_1287_v50).fm7(_1205_v3, _1285_cf54, _1205_v3, new BigNumber((_1239_v33).length), globalState))).update(_1286_cf53, _module.__default.abs(_1285_cf54))).IsSubsetOf(_1325_v73);
          let _rhs158 = _dafny.Seq.of((_dafny.ZERO).minus(_1285_cf54), p0, new BigNumber(276), (_dafny.ZERO).minus(_1285_cf54));
          let _rhs159 = _1324_v72;
          let _lhs104 = _1323_v71;
          let _lhs105 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_1323_v71).length));
          _1285_cf54 = _rhs155;
          _1322_v70 = _rhs156;
          r0 = _rhs157;
          _1295_v55 = _rhs158;
          _lhs104[_lhs105] = _rhs159;
        }
      } else if (_source16.is_DC25) {
        let _1328___mcc_h10 = (_source16).cf44;
        let _1329_cf44 = _1328___mcc_h10;
        let _1330_v74;
        let _nw186 = new _module.C0();
        _nw186.__ctor();
        _1330_v74 = _nw186;
        let _1331_v75;
        _1331_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1330_v74,_1207_v5);
        _1331_v75 = (_1331_v75).update(_1330_v74, _dafny.Seq.UnicodeFromString("nervdqu"));
        if (_module.__default.fm20(_1205_v3, globalState)) {
          _1209_v7 = _1209_v7;
          let _1332_v76;
          _1332_v76 = new BigNumber(-353);
          _1332_v76 = new BigNumber((_dafny.Seq.Concat(_1207_v5, _1207_v5)).length);
          let _arr45 = _1206_v4.f2;
          let _index219 = _module.__default.safeIndex(new BigNumber(276), new BigNumber((_1206_v4.f2).length));
          _arr45[_index219] = (p0).isLessThanOrEqualTo(new BigNumber((function () {
            let _coll49 = new _dafny.Set();
            for (const _compr_49 of (_dafny.Map.Empty.slice().updateUnsafe(p0,_1205_v3)).Keys.Elements) {
              let _1333_v77 = _compr_49;
              if ((_dafny.Map.Empty.slice().updateUnsafe(p0,_1205_v3)).contains(_1333_v77)) {
                _coll49.add(_module.__default.safeDivisionInt(_1333_v77, new BigNumber(893)));
              }
            }
            return _coll49;
          }()).length));
          let _arr46 = _1206_v4.f2;
          let _index220 = _module.__default.safeIndex(new BigNumber(276), new BigNumber((_1206_v4.f2).length));
          _arr46[_index220] = _1205_v3;
          let _1334_v78;
          _1334_v78 = _dafny.Map.Empty.slice().updateUnsafe(_1207_v5,_dafny.Seq.UnicodeFromString("ockkobn"));
          let _1335_v79;
          _1335_v79 = _module.D10.create_DC33(_1205_v3);
          let _1336_v80;
          _1336_v80 = _dafny.Seq.of(_1335_v79, _1335_v79, _1335_v79, _module.D10.create_DC33((_1206_v4.f2)[_module.__default.safeIndex(new BigNumber(276), new BigNumber((_1206_v4.f2).length))]));
          let _1337_v81;
          _1337_v81 = _dafny.Map.Empty.slice().updateUnsafe((((_1334_v78).contains(_1207_v5)) ? ((_1334_v78).get(_1207_v5)) : (_1207_v5)),_1336_v80);
          let _1338_v82;
          _1338_v82 = _module.D12.create_DC36(_1337_v81);
          _1337_v81 = (_1338_v82).dtor_cf66;
          let _1339_v83;
          let _nw187 = Array((new BigNumber(11)).toNumber()).fill([]);
          _1339_v83 = _nw187;
          let _index221 = _module.__default.safeIndex(new BigNumber(308), new BigNumber((_1339_v83).length));
          (_1339_v83)[_index221] = _this.f2;
          let _index222 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_1202_v0).length));
          (_1202_v0)[_index222] = _1332_v76;
          let _index223 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_1202_v0).length));
          (_1202_v0)[_index223] = _1332_v76;
          let _1340_v84;
          _1340_v84 = _module.D4.create_DC11(_dafny.Seq.of((_1206_v4.f2)[_module.__default.safeIndex(new BigNumber(276), new BigNumber((_1206_v4.f2).length))]));
          let _pat_let_tv17 = _1239_v33;
          let _index224 = _module.__default.safeIndex(new BigNumber(308), new BigNumber((_1339_v83).length));
          let _index225 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_1202_v0).length));
          let _index226 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_1202_v0).length));
          let _rhs160 = _1206_v4.f2;
          let _rhs161 = _1332_v76;
          let _rhs162 = new BigNumber(((function (_pat_let35_0) {
            return function (_1341_dt__update__tmp_h2) {
              return function (_pat_let36_0) {
                return function (_1342_dt__update_hcf28_h0) {
                  return _module.D4.create_DC11(_1342_dt__update_hcf28_h0);
                }(_pat_let36_0);
              }(_pat_let_tv17);
            }(_pat_let35_0);
          }(_1340_v84)).dtor_cf28).length);
          let _rhs163 = _1207_v5;
          let _rhs164 = p0;
          let _lhs106 = _1339_v83;
          let _lhs107 = _module.__default.safeIndex(new BigNumber(308), new BigNumber((_1339_v83).length));
          let _lhs108 = _1202_v0;
          let _lhs109 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_1202_v0).length));
          let _lhs110 = _1202_v0;
          let _lhs111 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_1202_v0).length));
          _lhs106[_lhs107] = _rhs160;
          _lhs108[_lhs109] = _rhs161;
          _1332_v76 = _rhs162;
          _1207_v5 = _rhs163;
          _lhs110[_lhs111] = _rhs164;
        } else {
          (_1206_v4).f2 = _1206_v4.f2;
          let _1343_v85;
          _1343_v85 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1205_v3);
          let _1344_v86;
          _1344_v86 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          _1343_v85 = (_1343_v85).update((((_1344_v86).contains(p0)) ? ((_1344_v86).get(p0)) : (p0)), _1205_v3);
          let _arr47 = _this.f2;
          let _index227 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_this.f2).length));
          _arr47[_index227] = _1205_v3;
          let _arr48 = _this.f2;
          let _index228 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_this.f2).length));
          _arr48[_index228] = _1205_v3;
          let _1345_v87;
          _1345_v87 = new BigNumber(427);
          let _1346_v88;
          _1346_v88 = _dafny.Map.Empty.slice().updateUnsafe((_this.f2)[_module.__default.safeIndex(new BigNumber(878), new BigNumber((_this.f2).length))],_1345_v87);
          let _1347_v89;
          _1347_v89 = _dafny.Seq.of(p0, _1345_v87, _1345_v87, new BigNumber(-288), _1345_v87);
          _1345_v87 = (((_1346_v88).contains(_1205_v3)) ? ((_1346_v88).get(_1205_v3)) : (new BigNumber((_1347_v89).length)));
          let _arr49 = _this.f2;
          let _index229 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_this.f2).length));
          _arr49[_index229] = _1205_v3;
        }
        let _1348_v90;
        let _init41 = ((_1349_v5) => function (_1350_i10) {
          return _1349_v5;
        })(_1207_v5);
        let _nw188 = Array((new BigNumber(20)).toNumber());
        for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw188.length); _i0_41++) {
          _nw188[_i0_41] = _init41(new BigNumber(_i0_41));
        }
        _1348_v90 = _nw188;
        let _index230 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_1348_v90).length));
        (_1348_v90)[_index230] = _dafny.Seq.UnicodeFromString("nrklejhh");
        let _index231 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_1348_v90).length));
        (_1348_v90)[_index231] = _1207_v5;
        r0 = (_1330_v74).fm7(true, p0, !_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(691)), ((_1351_v7) => function (_1352_i11) {
          return _1351_v7;
        })(_1209_v7)), (_1348_v90)[_module.__default.safeIndex(new BigNumber(956), new BigNumber((_1348_v90).length))]), _module.__default.safeDivisionInt(p0, p0), globalState);
      } else {
        let _1353___mcc_h11 = (_source16).cf55;
        let _1354_cf55 = _1353___mcc_h11;
        let _1355_v91;
        _1355_v91 = new BigNumber(730);
        let _1356_v92;
        _1356_v92 = _module.D5.create_DC17(p0);
        _1355_v91 = (_1356_v92).dtor_cf35;
        let _1357_v93;
        _1357_v93 = _dafny.Map.Empty.slice().updateUnsafe(!(_1205_v3),_1209_v7);
        let _1358_v94;
        _1358_v94 = _module.D3.create_DC10(_1205_v3, (((_1357_v93).contains(_1205_v3)) ? ((_1357_v93).get(_1205_v3)) : (_1209_v7)), _1205_v3, _dafny.Seq.IsProperPrefixOf(_1207_v5, _dafny.Seq.UnicodeFromString("ligdce")), _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("drun"), _1207_v5));
        let _1359_v95;
        _1359_v95 = _dafny.MultiSet.fromElements(p0, _1355_v91, _1355_v91);
        _1358_v94 = _module.D3.create_DC10(!((_module.__default.fm0((((_1359_v95).contains(p0)) ? ((_1359_v95).get(p0)) : (_1355_v91)), globalState)).isLessThan(_1355_v91)), _1209_v7, _module.__default.fm20(_1205_v3, globalState), _1205_v3, _module.__default.fm20(_1205_v3, globalState));
        let _arr50 = _1206_v4.f2;
        let _index232 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_1206_v4.f2).length));
        _arr50[_index232] = _module.__default.fm20(false, globalState);
        let _1360_v96;
        _1360_v96 = _dafny.Seq.of(_1202_v0);
        let _1361_v97;
        let _nw189 = Array((new BigNumber(2)).toNumber());
        _nw189[(_dafny.ZERO).toNumber()] = ((_1205_v3) ? (_dafny.Seq.of(_1202_v0, _1202_v0)) : (_1360_v96));
        _nw189[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_1360_v96, _module.__default.safeIndex(p0, new BigNumber((_1360_v96).length)), _1202_v0), _dafny.Seq.of(_1202_v0, _1202_v0, _1202_v0));
        _1361_v97 = _nw189;
        let _index233 = _module.__default.safeIndex(new BigNumber(839), new BigNumber((_1361_v97).length));
        (_1361_v97)[_index233] = _1360_v96;
        let _1362_v98;
        _1362_v98 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1202_v0);
        let _1363_v99;
        _1363_v99 = _dafny.Seq.of(_dafny.Seq.Concat(_1360_v96, _dafny.Seq.of((((_1362_v98).contains(_1355_v91)) ? ((_1362_v98).get(_1355_v91)) : (_1202_v0)))), _1360_v96);
        let _arr51 = _1206_v4.f2;
        let _index234 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_1206_v4.f2).length));
        let _index235 = _module.__default.safeIndex(new BigNumber(839), new BigNumber((_1361_v97).length));
        let _rhs165 = _1205_v3;
        let _rhs166 = (_1363_v99)[_module.__default.safeIndex(_1355_v91, new BigNumber((_1363_v99).length))];
        let _rhs167 = _1240_v34;
        let _lhs112 = _1206_v4.f2;
        let _lhs113 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_1206_v4.f2).length));
        let _lhs114 = _1361_v97;
        let _lhs115 = _module.__default.safeIndex(new BigNumber(839), new BigNumber((_1361_v97).length));
        _lhs112[_lhs113] = _rhs165;
        _lhs114[_lhs115] = _rhs166;
        _1240_v34 = _rhs167;
        let _arr52 = _1206_v4.f2;
        let _index236 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_1206_v4.f2).length));
        _arr52[_index236] = _1205_v3;
      }
      let _1364_v100;
      _1364_v100 = _dafny.MultiSet.fromElements((((p2).contains(_1205_v3)) ? ((p2).get(_1205_v3)) : (_this.f2)));
      let _1365_v101;
      _1365_v101 = _module.D7.create_DC22(_1364_v100);
      let _source18 = _1365_v101;
      if (_source18.is_DC23) {
        let _1366___mcc_h17 = (_source18).cf42;
        let _1367_cf42 = _1366___mcc_h17;
        let _1368_v102;
        _1368_v102 = _dafny.Seq.of(_1239_v33);
        let _1369_v103;
        _1369_v103 = _dafny.MultiSet.fromElements(false);
        let _1370_v104;
        _1370_v104 = _dafny.Seq.of(_1207_v5);
        let _1371_v105;
        _1371_v105 = _dafny.Seq.of(_1367_cf42, _1367_cf42);
        let _1372_v106;
        _1372_v106 = _module.D13.create_DC42((_1368_v102)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_1367_cf42)).length), new BigNumber((_1368_v102).length))], _1369_v103, _1370_v104, _1205_v3, _1371_v105);
        _1367_cf42 = new BigNumber(((_1372_v106).dtor_cf78).cardinality());
        _1367_cf42 = _1367_cf42;
        let _source19 = _module.__default.fm33(_1371_v105, p0, globalState);
        if (_source19.is_DC26) {
          let _1373___mcc_h20 = (_source19).cf45;
          let _1374___mcc_h21 = (_source19).cf46;
          let _1375___mcc_h22 = (_source19).cf47;
          let _1376___mcc_h23 = (_source19).cf48;
          let _1377_cf48 = _1376___mcc_h23;
          let _1378_cf47 = _1375___mcc_h22;
          let _1379_cf46 = _1374___mcc_h21;
          let _1380_cf45 = _1373___mcc_h20;
          let _arr53 = _1206_v4.f2;
          let _index237 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_1206_v4.f2).length));
          _arr53[_index237] = ((_1205_v3) ? (_1380_cf45) : (_1380_cf45));
          let _arr54 = _1206_v4.f2;
          let _index238 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_1206_v4.f2).length));
          let _rhs168 = (_1239_v33)[_module.__default.safeIndex(new BigNumber((_1207_v5).length), new BigNumber((_1239_v33).length))];
          let _rhs169 = !(_1380_cf45) || (_1205_v3);
          let _rhs170 = _1206_v4;
          let _lhs116 = _1206_v4.f2;
          let _lhs117 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_1206_v4.f2).length));
          _1379_cf46 = _rhs168;
          _lhs116[_lhs117] = _rhs169;
          _1206_v4 = _rhs170;
          _1379_cf46 = !(_1205_v3);
          _1380_cf45 = !(false);
          r0 = _1380_cf45;
        } else if (_source19.is_DC27) {
          let _1381___mcc_h24 = (_source19).cf49;
          let _1382___mcc_h25 = (_source19).cf50;
          let _1383___mcc_h26 = (_source19).cf51;
          let _1384___mcc_h27 = (_source19).cf52;
          let _1385_cf52 = _1384___mcc_h27;
          let _1386_cf51 = _1383___mcc_h26;
          let _1387_cf50 = _1382___mcc_h25;
          let _1388_cf49 = _1381___mcc_h24;
          _1388_cf49 = _1388_cf49;
          _1205_v3 = _1205_v3;
          let _rhs171 = new BigNumber((_1371_v105).length);
          let _rhs172 = (_1369_v103).IsProperSubsetOf(_dafny.MultiSet.fromElements(false));
          let _rhs173 = (_dafny.ZERO).minus(_1367_cf42);
          _1385_cf52 = _rhs171;
          _1388_cf49 = _rhs172;
          _1385_cf52 = _rhs173;
          let _1389_v107;
          _1389_v107 = _dafny.MultiSet.fromElements(_1369_v103, _1369_v103);
          let _1390_v108;
          _1390_v108 = _dafny.Map.Empty.slice().updateUnsafe(_1209_v7,_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(_1387_cf50, _1387_cf50, _1388_cf49, _1205_v3)));
          _1389_v107 = ((((_1390_v108).contains(_1209_v7)) ? ((_1390_v108).get(_1209_v7)) : (_dafny.MultiSet.FromArray(_dafny.Seq.of(_1369_v103, _1369_v103))))).Union((_1389_v107).update(_dafny.MultiSet.fromElements(_1387_cf50), _module.__default.abs(_1367_cf42)));
        } else if (_source19.is_DC28) {
          let _1391___mcc_h28 = (_source19).cf53;
          let _1392___mcc_h29 = (_source19).cf54;
          let _1393_cf54 = _1392___mcc_h29;
          let _1394_cf53 = _1391___mcc_h28;
          _1205_v3 = _dafny.Seq.IsPrefixOf(_1239_v33, _dafny.Seq.Concat(_dafny.Seq.of(true, _1205_v3, _1205_v3), _1239_v33));
          let _1395_v109;
          _1395_v109 = _dafny.Set.fromElements(p0, new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_1202_v0,new BigNumber((_1207_v5).length))).Merge(_1203_v1)).length), (_dafny.ZERO).minus(_1367_cf42));
          let _rhs174 = _1205_v3;
          let _rhs175 = _1394_cf53;
          let _rhs176 = _1395_v109;
          let _rhs177 = _1394_cf53;
          _1205_v3 = _rhs174;
          _1394_cf53 = _rhs175;
          _1395_v109 = _rhs176;
          _1205_v3 = _rhs177;
          _1367_cf42 = _1367_cf42;
          let _1396_v110;
          _1396_v110 = _dafny.Map.Empty.slice().updateUnsafe(_1205_v3,_1207_v5);
          let _rhs178 = _1394_cf53;
          let _rhs179 = ((_1393_cf54).multipliedBy(new BigNumber((_1371_v105).length))).isEqualTo(new BigNumber(((_1240_v34).update((_1239_v33)[_module.__default.safeIndex(new BigNumber((_1396_v110).length), new BigNumber((_1239_v33).length))], !(_1394_cf53))).length));
          r0 = _rhs178;
          _1205_v3 = _rhs179;
        } else if (_source19.is_DC25) {
          let _1397___mcc_h30 = (_source19).cf44;
          let _1398_cf44 = _1397___mcc_h30;
          _1202_v0 = _1202_v0;
          _1205_v3 = _1205_v3;
          let _1399_v111;
          _1399_v111 = _module.D3.create_DC10(!(_1205_v3), _1209_v7, _1205_v3, _1205_v3, true);
          let _1400_v112;
          _1400_v112 = _dafny.MultiSet.fromElements(_1399_v111);
          let _1401_v113;
          _1401_v113 = _dafny.Seq.of(_1399_v111, _1399_v111, _1399_v111, _1399_v111, _module.D3.create_DC10(_1205_v3, _1209_v7, _1205_v3, _module.__default.fm20(_1205_v3, globalState), _1205_v3));
          let _1402_v114;
          let _nw190 = Array((new BigNumber(3)).toNumber());
          _nw190[(_dafny.ZERO).toNumber()] = _1400_v112;
          _nw190[(_dafny.ONE).toNumber()] = (_1400_v112).Difference(_1400_v112);
          _nw190[(new BigNumber(2)).toNumber()] = (_dafny.MultiSet.FromArray(_1401_v113)).Union(_1400_v112);
          _1402_v114 = _nw190;
          let _index239 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_1402_v114).length));
          (_1402_v114)[_index239] = _1400_v112;
          let _index240 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_1402_v114).length));
          (_1402_v114)[_index240] = (_1400_v112).Union(_1400_v112);
          let _1403_v115;
          _1403_v115 = _module.D4.create_DC11(_1239_v33);
          let _1404_v116;
          _1404_v116 = _module.D5.create_DC15(_1403_v115, _1205_v3);
          let _1405_v117;
          _1405_v117 = _dafny.Seq.of(_1404_v116);
          let _rhs180 = !(_1367_cf42).isEqualTo(_module.__default.safeDivisionInt(p0, p0));
          let _rhs181 = !(_1205_v3);
          let _rhs182 = _module.__default.fm20(_dafny.Seq.contains(_1405_v117, _1404_v116), globalState);
          _1205_v3 = _rhs180;
          _1205_v3 = _rhs181;
          r0 = _rhs182;
        } else {
          let _1406___mcc_h31 = (_source19).cf55;
          let _1407_cf55 = _1406___mcc_h31;
          r0 = _1205_v3;
          _1367_cf42 = _module.__default.safeModuloInt(p0, p0);
          let _1408_v118;
          _1408_v118 = _dafny.Map.Empty.slice().updateUnsafe(_1205_v3,p0);
          _1408_v118 = (_1408_v118).update(_1205_v3, _module.__default.safeDivisionInt(p0, new BigNumber((_dafny.Seq.UnicodeFromString("mgnbqbnka")).length)));
          let _1409_v119;
          _1409_v119 = _dafny.MultiSet.fromElements(_1206_v4, _1206_v4);
          let _1410_v120;
          _1410_v120 = _dafny.Set.fromElements(_1409_v119);
          let _1411_v121;
          _1411_v121 = _dafny.Map.Empty.slice().updateUnsafe(_1410_v120,_1369_v103);
          _1411_v121 = (_1411_v121).update(_1410_v120, _dafny.MultiSet.FromArray(((_1205_v3) ? (_dafny.Seq.of(_1205_v3, _1205_v3)) : (_1239_v33))));
        }
        let _rhs183 = _this.f2;
        let _rhs184 = (true) && (_1205_v3);
        let _rhs185 = _1202_v0;
        let _lhs118 = _1206_v4;
        _lhs118.f2 = _rhs183;
        _1205_v3 = _rhs184;
        _1202_v0 = _rhs185;
      } else if (_source18.is_DC24) {
        let _1412___mcc_h18 = (_source18).cf43;
        let _1413_cf43 = _1412___mcc_h18;
        let _1414_v122;
        let _nw191 = new _module.C0();
        _nw191.__ctor();
        _1414_v122 = _nw191;
        let _1415_v123;
        let _nw192 = new _module.C2();
        _nw192.__ctor(_1205_v3, p0, _1204_v2, _1413_cf43, _1206_v4.f2);
        _1415_v123 = _nw192;
        _1415_v123 = _1415_v123;
        let _1416_v124;
        _1416_v124 = _module.D2.create_DC6(_1209_v7, (_1415_v123).f8);
        let _1417_v125;
        _1417_v125 = _dafny.Set.fromElements(_1209_v7, (_1416_v124).dtor_cf14, _1209_v7, new _dafny.CodePoint('b'.codePointAt(0)));
        (_1415_v123).f7 = (_dafny.Set.fromElements((_1207_v5)[_module.__default.safeIndex(p0, new BigNumber((_1207_v5).length))])).IsSubsetOf(_1417_v125);
        (_1415_v123).f7 = !((_1413_cf43) === (_1413_cf43));
      } else {
        let _1418___mcc_h19 = (_source18).cf41;
        let _1419_cf41 = _1418___mcc_h19;
        _1205_v3 = _1205_v3;
        let _1420_v126;
        let _nw193 = new _module.C2();
        _nw193.__ctor(_1205_v3, (_this).fm2(globalState), _1204_v2, _1205_v3, _1206_v4.f2);
        _1420_v126 = _nw193;
        let _1421_v127;
        let _nw194 = Array((new BigNumber(13)).toNumber());
        _nw194[(_dafny.ZERO).toNumber()] = _1420_v126;
        _nw194[(_dafny.ONE).toNumber()] = _1420_v126;
        _nw194[(new BigNumber(2)).toNumber()] = _1420_v126;
        _nw194[(new BigNumber(3)).toNumber()] = _1420_v126;
        _nw194[(new BigNumber(4)).toNumber()] = _1420_v126;
        _nw194[(new BigNumber(5)).toNumber()] = _1420_v126;
        _nw194[(new BigNumber(6)).toNumber()] = _1420_v126;
        _nw194[(new BigNumber(7)).toNumber()] = _1420_v126;
        _nw194[(new BigNumber(8)).toNumber()] = _1420_v126;
        _nw194[(new BigNumber(9)).toNumber()] = _1420_v126;
        _nw194[(new BigNumber(10)).toNumber()] = _1420_v126;
        _nw194[(new BigNumber(11)).toNumber()] = _1420_v126;
        _nw194[(new BigNumber(12)).toNumber()] = _1420_v126;
        _1421_v127 = _nw194;
        let _1422_v128;
        _1422_v128 = _dafny.Seq.of(_module.__default.fm34(_1205_v3, p0, globalState));
        let _1423_v129;
        _1423_v129 = _dafny.Map.Empty.slice().updateUnsafe(_1421_v127,(_1422_v128)[_module.__default.safeIndex(p0, new BigNumber((_1422_v128).length))]);
        let _1424_v130;
        _1424_v130 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1421_v127);
        let _1425_v131;
        _1425_v131 = _dafny.Map.Empty.slice().updateUnsafe(_1209_v7,p0);
        let _1426_v132;
        _1426_v132 = _dafny.Map.Empty.slice().updateUnsafe(p0,new _dafny.CodePoint('p'.codePointAt(0)));
        let _1427_v133;
        _1427_v133 = _dafny.Seq.of(new BigNumber(((_1425_v131).update(_1209_v7, new BigNumber(((_1426_v132).update(p0, _1209_v7)).length))).length));
        let _1428_v134;
        _1428_v134 = _dafny.MultiSet.fromElements(_1427_v133, _1427_v133, _dafny.Seq.of(p0), _1427_v133, _dafny.Seq.Create(_module.__default.abs(new BigNumber(968)), ((_1429_v33) => function (_1430_i12) {
          return new BigNumber((_1429_v33).length);
        })(_1239_v33)));
        _1205_v3 = ((((_1423_v129).contains((((_1424_v130).contains(p0)) ? ((_1424_v130).get(p0)) : (_1421_v127)))) ? ((_1423_v129).get((((_1424_v130).contains(p0)) ? ((_1424_v130).get(p0)) : (_1421_v127)))) : (_dafny.MultiSet.fromElements(_dafny.Seq.of(p0))))).equals(_1428_v134);
        _1420_v126 = _1420_v126;
        let _1431_v135;
        let _nw195 = new _module.C0();
        _nw195.__ctor();
        _1431_v135 = _nw195;
        let _1432_v136;
        _1432_v136 = _dafny.Map.Empty.slice().updateUnsafe(_1205_v3,_1431_v135);
        let _1433_v137;
        _1433_v137 = _dafny.MultiSet.fromElements(new BigNumber((_1432_v136).length));
        let _1434_v138;
        _1434_v138 = _dafny.MultiSet.fromElements(_1433_v137);
        let _1435_v139;
        _1435_v139 = _module.D2.create_DC7(p0, p0, _1434_v138, p0, p0);
        let _1436_v140;
        _1436_v140 = _module.D2.create_DC8(_module.D2.create_DC8(_1435_v139));
        let _pat_let_tv18 = _1209_v7;
        let _pat_let_tv19 = p0;
        let _1437_v141;
        _1437_v141 = _dafny.Map.Empty.slice().updateUnsafe(p0,function (_pat_let37_0) {
          return function (_1438_dt__update__tmp_h3) {
            return function (_pat_let38_0) {
              return function (_1439_dt__update_hcf21_h0) {
                return _module.D2.create_DC8(_1439_dt__update_hcf21_h0);
              }(_pat_let38_0);
            }(_module.D2.create_DC8(_module.D2.create_DC6(_pat_let_tv18, _pat_let_tv19)));
          }(_pat_let37_0);
        }(_1436_v140));
        _1437_v141 = (_1437_v141).update(p0, _1436_v140);
      }
      let _1440_v142;
      _1440_v142 = _module.D4.create_DC12();
      let _pat_let_tv20 = _1240_v34;
      let _pat_let_tv21 = _1205_v3;
      let _pat_let_tv22 = _1240_v34;
      let _pat_let_tv23 = _1205_v3;
      let _pat_let_tv24 = _1205_v3;
      let _pat_let_tv25 = _1205_v3;
      r0 = function (_source20) {
        if (_source20.is_DC12) {
          if ((_pat_let_tv20).contains(_pat_let_tv21)) {
            return (_pat_let_tv22).get(_pat_let_tv23);
          } else {
            return !(_pat_let_tv24);
          }
        } else if (_source20.is_DC13) {
          let _1441___mcc_h32 = (_source20).cf29;
          let _1442___mcc_h33 = (_source20).cf30;
          let _1443___mcc_h34 = (_source20).cf31;
          let _1444_cf31 = _1443___mcc_h34;
          let _1445_cf30 = _1442___mcc_h33;
          let _1446_cf29 = _1441___mcc_h32;
          return false;
        } else {
          let _1447___mcc_h35 = (_source20).cf28;
          let _1448_cf28 = _1447___mcc_h35;
          return _pat_let_tv25;
        }
      }(_1440_v142);
      return r0;
    }
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.MultiSet.Empty;
      let r1 = false;
      let _1449_v0;
      _1449_v0 = new BigNumber(980);
      let _1450_v1;
      _1450_v1 = false;
      let _1451_v2;
      _1451_v2 = _dafny.Seq.of(_1450_v1);
      _1449_v0 = (_1449_v0).multipliedBy(((false) ? (_1449_v0) : (new BigNumber((_1451_v2).length))));
      let _1452_v3;
      _1452_v3 = _module.D4.create_DC11(_1451_v2);
      let _1453_v4;
      _1453_v4 = _module.D5.create_DC15(_1452_v3, _1450_v1);
      let _pat_let_tv26 = _1451_v2;
      let _pat_let_tv27 = _1450_v1;
      let _pat_let_tv28 = p1;
      let _pat_let_tv29 = _1451_v2;
      let _pat_let_tv30 = p1;
      let _pat_let_tv31 = _1450_v1;
      let _pat_let_tv32 = _1450_v1;
      let _pat_let_tv33 = _1450_v1;
      let _pat_let_tv34 = _1451_v2;
      let _pat_let_tv35 = _1450_v1;
      let _pat_let_tv36 = _1451_v2;
      let _pat_let_tv37 = _1450_v1;
      let _pat_let_tv38 = _1450_v1;
      let _pat_let_tv39 = _1450_v1;
      let _pat_let_tv40 = _1450_v1;
      let _pat_let_tv41 = p1;
      let _pat_let_tv42 = _1450_v1;
      let _pat_let_tv43 = _1450_v1;
      let _pat_let_tv44 = _1451_v2;
      let _pat_let_tv45 = _1450_v1;
      let _pat_let_tv46 = p1;
      let _pat_let_tv47 = _1450_v1;
      let _pat_let_tv48 = _1450_v1;
      let _pat_let_tv49 = _1450_v1;
      let _pat_let_tv50 = _1450_v1;
      let _pat_let_tv51 = _1450_v1;
      let _pat_let_tv52 = p1;
      let _pat_let_tv53 = _1450_v1;
      let _pat_let_tv54 = _1450_v1;
      let _pat_let_tv55 = _1450_v1;
      let _pat_let_tv56 = _1450_v1;
      let _pat_let_tv57 = _1450_v1;
      let _pat_let_tv58 = _1451_v2;
      let _pat_let_tv59 = _1449_v0;
      let _pat_let_tv60 = _1451_v2;
      let _pat_let_tv61 = p1;
      let _pat_let_tv62 = _1450_v1;
      let _pat_let_tv63 = _1450_v1;
      let _pat_let_tv64 = _1450_v1;
      let _pat_let_tv65 = _1450_v1;
      let _pat_let_tv66 = _1450_v1;
      let _pat_let_tv67 = _1450_v1;
      let _pat_let_tv68 = _1450_v1;
      let _pat_let_tv69 = _1450_v1;
      let _pat_let_tv70 = p1;
      let _pat_let_tv71 = _1450_v1;
      let _pat_let_tv72 = _1450_v1;
      let _pat_let_tv73 = _1450_v1;
      let _pat_let_tv74 = _1451_v2;
      let _pat_let_tv75 = _1450_v1;
      let _pat_let_tv76 = p1;
      let _pat_let_tv77 = _1450_v1;
      let _pat_let_tv78 = _1450_v1;
      let _pat_let_tv79 = _1450_v1;
      let _pat_let_tv80 = _1450_v1;
      let _pat_let_tv81 = _1450_v1;
      let _pat_let_tv82 = _1450_v1;
      let _pat_let_tv83 = _1451_v2;
      let _pat_let_tv84 = _1450_v1;
      let _pat_let_tv85 = p1;
      let _pat_let_tv86 = _1450_v1;
      let _pat_let_tv87 = _1450_v1;
      let _pat_let_tv88 = _1450_v1;
      let _pat_let_tv89 = _1450_v1;
      let _pat_let_tv90 = _1450_v1;
      let _pat_let_tv91 = _1451_v2;
      let _pat_let_tv92 = _1449_v0;
      let _pat_let_tv93 = _1451_v2;
      let _pat_let_tv94 = _1449_v0;
      let _pat_let_tv95 = _1450_v1;
      let _pat_let_tv96 = p1;
      let _pat_let_tv97 = _1450_v1;
      let _pat_let_tv98 = _1450_v1;
      _1449_v0 = new BigNumber((function (_source21) {
        if (_source21.is_DC15) {
          let _1454___mcc_h0 = (_source21).cf33;
          let _1455___mcc_h1 = (_source21).cf34;
          let _1456_cf34 = _1455___mcc_h1;
          let _1457_cf33 = _1454___mcc_h0;
          return (_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv26,_dafny.MultiSet.fromElements(_module.D3.create_DC10(_pat_let_tv27, _pat_let_tv28, true, _1456_cf34, _1456_cf34)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv29,(_module.D14.create_DC44(_dafny.MultiSet.FromArray(_dafny.Seq.of(_module.D3.create_DC10(true, _pat_let_tv30, _pat_let_tv31, _pat_let_tv32, _pat_let_tv33))))).dtor_cf83));
        } else if (_source21.is_DC16) {
          return function () {
            let _coll50 = new _dafny.Map();
            for (const _compr_50 of (_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv34,_pat_let_tv35)).Keys.Elements) {
              let _1458_v5 = _compr_50;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv36,_pat_let_tv37)).contains(_1458_v5)) {
                _coll50.push([_1458_v5,_dafny.MultiSet.fromElements(_module.D3.create_DC10(_pat_let_tv38, new _dafny.CodePoint('p'.codePointAt(0)), false, true, _pat_let_tv39), _module.D3.create_DC10(_pat_let_tv40, _pat_let_tv41, _pat_let_tv42, false, _pat_let_tv43))]);
              }
            }
            return _coll50;
          }();
        } else if (_source21.is_DC17) {
          let _1459___mcc_h2 = (_source21).cf35;
          let _1460_cf35 = _1459___mcc_h2;
          return _dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv44,_dafny.MultiSet.fromElements(_module.D3.create_DC10(_pat_let_tv45, _pat_let_tv46, _pat_let_tv47, _pat_let_tv48, _pat_let_tv49), _module.D3.create_DC10(_pat_let_tv50, (_module.D3.create_DC10(_pat_let_tv51, _pat_let_tv52, _pat_let_tv53, _pat_let_tv54, !(true))).dtor_cf24, _pat_let_tv55, _pat_let_tv56, _pat_let_tv57), _module.D3.create_DC10((_pat_let_tv58)[_module.__default.safeIndex(_pat_let_tv59, new BigNumber((_pat_let_tv60).length))], _pat_let_tv61, _pat_let_tv62, _pat_let_tv63, _pat_let_tv64), _module.D3.create_DC10(_pat_let_tv65, new _dafny.CodePoint('y'.codePointAt(0)), !(_pat_let_tv66), _pat_let_tv67, _pat_let_tv68), _module.D3.create_DC10(_pat_let_tv69, _pat_let_tv70, _pat_let_tv71, _pat_let_tv72, _pat_let_tv73)));
        } else if (_source21.is_DC14) {
          let _1461___mcc_h3 = (_source21).cf32;
          let _1462_cf32 = _1461___mcc_h3;
          return (_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv74,_dafny.MultiSet.fromElements(_module.D3.create_DC10(_pat_let_tv75, _pat_let_tv76, _pat_let_tv77, _pat_let_tv78, _pat_let_tv79), _module.D3.create_DC10(true, new _dafny.CodePoint('o'.codePointAt(0)), _pat_let_tv80, _pat_let_tv81, _pat_let_tv82)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv83,_dafny.MultiSet.fromElements(_module.D3.create_DC10(_pat_let_tv84, _pat_let_tv85, false, _pat_let_tv86, _pat_let_tv87), _module.D3.create_DC10(_pat_let_tv88, new _dafny.CodePoint('x'.codePointAt(0)), true, _pat_let_tv89, _pat_let_tv90))));
        } else {
          let _1463___mcc_h4 = (_source21).cf36;
          let _1464_cf36 = _1463___mcc_h4;
          return function () {
            let _coll51 = new _dafny.Map();
            for (const _compr_51 of (_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv91,_pat_let_tv92)).Keys.Elements) {
              let _1465_v6 = _compr_51;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv93,_pat_let_tv94)).contains(_1465_v6)) {
                _coll51.push([_1465_v6,_dafny.MultiSet.fromElements(_module.D3.create_DC10(_pat_let_tv95, _pat_let_tv96, _pat_let_tv97, _pat_let_tv98, false))]);
              }
            }
            return _coll51;
          }();
        }
      }(_1453_v4)).length);
      let _1466_i0;
      _1466_i0 = _dafny.ZERO;
      L7: {
        let _pat_let_tv99 = p0;
        let _pat_let_tv100 = _1450_v1;
        let _pat_let_tv101 = _1450_v1;
        while (function (_source22) {
          if (_source22.is_DC26) {
            let _1485___mcc_h5 = (_source22).cf45;
            let _1486___mcc_h6 = (_source22).cf46;
            let _1487___mcc_h7 = (_source22).cf47;
            let _1488___mcc_h8 = (_source22).cf48;
            let _1489_cf48 = _1488___mcc_h8;
            let _1490_cf47 = _1487___mcc_h7;
            let _1491_cf46 = _1486___mcc_h6;
            let _1492_cf45 = _1485___mcc_h5;
            return _1492_cf45;
          } else if (_source22.is_DC27) {
            let _1493___mcc_h9 = (_source22).cf49;
            let _1494___mcc_h10 = (_source22).cf50;
            let _1495___mcc_h11 = (_source22).cf51;
            let _1496___mcc_h12 = (_source22).cf52;
            let _1497_cf52 = _1496___mcc_h12;
            let _1498_cf51 = _1495___mcc_h11;
            let _1499_cf50 = _1494___mcc_h10;
            let _1500_cf49 = _1493___mcc_h9;
            return _1500_cf49;
          } else if (_source22.is_DC28) {
            let _1501___mcc_h13 = (_source22).cf53;
            let _1502___mcc_h14 = (_source22).cf54;
            let _1503_cf54 = _1502___mcc_h14;
            let _1504_cf53 = _1501___mcc_h13;
            return !(new BigNumber(208)).isEqualTo(_pat_let_tv99);
          } else if (_source22.is_DC25) {
            let _1505___mcc_h15 = (_source22).cf44;
            let _1506_cf44 = _1505___mcc_h15;
            return _pat_let_tv100;
          } else {
            let _1507___mcc_h16 = (_source22).cf55;
            let _1508_cf55 = _1507___mcc_h16;
            return _pat_let_tv101;
          }
        }(_module.D8.create_DC28(_1450_v1, _module.__default.fm0(p0, globalState)))) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1466_i0)) {
              break L7;
            }
            _1466_i0 = (_1466_i0).plus(_dafny.ONE);
            r1 = (_1451_v2)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_dafny.ZERO).minus(_1449_v0), _1449_v0), new BigNumber((_1451_v2).length))];
            _1450_v1 = _1450_v1;
            _1450_v1 = !(_1450_v1);
            if (true) {
              let _1467_v7;
              _1467_v7 = new _dafny.CodePoint('y'.codePointAt(0));
              let _1468_v8;
              let _nw196 = new _module.C0();
              _nw196.__ctor();
              _1468_v8 = _nw196;
              let _1469_v9;
              _1469_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1468_v8,_1467_v7);
              _1467_v7 = (((_1469_v9).contains(_1468_v8)) ? ((_1469_v9).get(_1468_v8)) : (new _dafny.CodePoint('l'.codePointAt(0))));
              r1 = _1450_v1;
              _1451_v2 = _1451_v2;
              let _1470_v10;
              let _init42 = ((_1471_v2, _1472_p0) => function (_1473_i1) {
                return _dafny.Map.Empty.slice().updateUnsafe(_1471_v2,_1472_p0);
              })(_1451_v2, p0);
              let _nw197 = Array((new BigNumber(27)).toNumber());
              for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw197.length); _i0_42++) {
                _nw197[_i0_42] = _init42(new BigNumber(_i0_42));
              }
              _1470_v10 = _nw197;
              let _1474_v11;
              _1474_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1451_v2,p0);
              let _index241 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_1470_v10).length));
              (_1470_v10)[_index241] = _1474_v11;
              let _index242 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_1470_v10).length));
              (_1470_v10)[_index242] = _1474_v11;
              let _1475_v12;
              _1475_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1450_v1,_1450_v1);
              _1475_v12 = (_1475_v12).update(_1450_v1, _1450_v1);
            } else {
              let _1476_v13;
              _1476_v13 = _dafny.MultiSet.fromElements(_1450_v1, _1450_v1);
              let _1477_v14;
              _1477_v14 = _dafny.Seq.of(_1449_v0, (((_1476_v13).contains(_1450_v1)) ? ((_1476_v13).get(_1450_v1)) : ((_this).fm2(globalState))), p0, new BigNumber((_1451_v2).length));
              let _1478_v15;
              _1478_v15 = _dafny.Seq.of(_1477_v14);
              let _1479_v16;
              _1479_v16 = _dafny.Set.fromElements((_1478_v15)[_module.__default.safeIndex(p0, new BigNumber((_1478_v15).length))]);
              _1479_v16 = _1479_v16;
              _1449_v0 = (new BigNumber((_dafny.Seq.Concat(_1451_v2, _1451_v2)).length)).minus(p0);
              _1449_v0 = _1449_v0;
              _1449_v0 = (p0).plus(p0);
              let _1480_v17;
              let _init43 = ((_1481_v1, _1482_v14) => function (_1483_i2) {
                return _dafny.Map.Empty.slice().updateUnsafe(_1481_v1,new BigNumber((_1482_v14).length));
              })(_1450_v1, _1477_v14);
              let _nw198 = Array((new BigNumber(14)).toNumber());
              for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw198.length); _i0_43++) {
                _nw198[_i0_43] = _init43(new BigNumber(_i0_43));
              }
              _1480_v17 = _nw198;
              let _1484_v18;
              let _nw199 = new _module.C1();
              _nw199.__ctor(_1480_v17, !(_1450_v1), _this.f2);
              _1484_v18 = _nw199;
            }
          }
        }
      }
      let _1509_v19;
      let _nw200 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
      _1509_v19 = _nw200;
      _1509_v19 = _1509_v19;
      let _1510_v20;
      _1510_v20 = _dafny.Seq.UnicodeFromString("qbtir");
      _1510_v20 = _dafny.Seq.Concat(_1510_v20, _dafny.Seq.Concat(_1510_v20, _dafny.Seq.UnicodeFromString("nxsjpdcw")));
      _1449_v0 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_1449_v0));
      let _1511_v21;
      _1511_v21 = _dafny.MultiSet.fromElements(_this.f2);
      r0 = _1511_v21;
      r1 = ((_1450_v1) ? (true) : (!(false)));
      return [r0, r1];
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f6 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f6) {
      let _this = this;
      (_this)._f6 = f6;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      return (new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6))).length)).multipliedBy(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-550),(_this).f6)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("su")).length),(_module.D0.create_DC1(false, (_this).f6, new BigNumber(972))).dtor_cf2))).length));
    };
    fm6(p0, p1, globalState) {
      let _this = this;
      return ((_dafny.Set.fromElements(_dafny.MultiSet.fromElements(!((_this).f6)))).Intersect(_dafny.Set.fromElements(_dafny.MultiSet.fromElements((_this).f6, (_this).f6)))).IsDisjointFrom((function () {
        let _coll52 = new _dafny.Set();
        for (const _compr_52 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((_this).f6, (_this).f6, (_this).f6, (_this).f6, false),(_this).f6)).Keys.Elements) {
          let _1512_v0 = _compr_52;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((_this).f6, (_this).f6, (_this).f6, (_this).f6, false),(_this).f6)).contains(_1512_v0)) {
            _coll52.add(_1512_v0);
          }
        }
        return _coll52;
      }()).Intersect(_dafny.Set.fromElements(_dafny.MultiSet.fromElements((_this).f6, (_this).f6), _dafny.MultiSet.fromElements((_this).f6, (_this).f6))));
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = _dafny.Seq.UnicodeFromString("");
      let _1513_v0;
      let _nw201 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1513_v0 = _nw201;
      _1513_v0 = _1513_v0;
      let _1514_v1;
      let _init44 = ((_1515_p0) => function (_1516_i1) {
        return _module.__default.safeDivisionInt(_1516_i1, (_dafny.ZERO).minus(_1515_p0));
      })(p0);
      let _nw202 = Array((new BigNumber(21)).toNumber());
      for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw202.length); _i0_44++) {
        _nw202[_i0_44] = _init44(new BigNumber(_i0_44));
      }
      _1514_v1 = _nw202;
      for (const _guard_loop_7 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1514_v1).length))) {
        let _1517_i0 = _guard_loop_7;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1517_i0)) && ((_1517_i0).isLessThan(new BigNumber((_1514_v1).length))))) {
          (_1514_v1)[(_1517_i0)] = (_1517_i0).multipliedBy(p0);
        }
      }
      let _hi8 = p2;
      for (let _1518_i2 = (_dafny.ZERO).minus(p2); _1518_i2.isLessThan(_hi8); _1518_i2 = _1518_i2.plus(_dafny.ONE)) {
        let _1519_v2;
        _1519_v2 = new BigNumber(578);
        let _1520_v3;
        _1520_v3 = _dafny.Seq.UnicodeFromString("s");
        let _rhs186 = p2;
        let _rhs187 = (((_this).f6) ? (_1518_i2) : ((_dafny.ZERO).minus(_1518_i2)));
        let _rhs188 = _1520_v3;
        let _rhs189 = (_dafny.ZERO).minus(p0);
        let _rhs190 = (_this).f6;
        _1519_v2 = _rhs186;
        _1519_v2 = _rhs187;
        r2 = _rhs188;
        _1519_v2 = _rhs189;
        r0 = _rhs190;
        let _1521_v4;
        let _nw203 = new _module.C0();
        _nw203.__ctor();
        _1521_v4 = _nw203;
        let _index243 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_1514_v1).length));
        (_1514_v1)[_index243] = p0;
        let _1522_v5;
        _1522_v5 = _module.D0.create_DC1((_this).f6, (_this).f6, new BigNumber(-38));
        let _1523_v6;
        _1523_v6 = _dafny.Seq.of(_1522_v5, _1522_v5, _module.D0.create_DC1((_this).f6, !(true), _1519_v2), _module.D0.create_DC1((_this).f6, (_this).f6, p0), _module.D0.create_DC1((_this).f6, (_this).f6, _1519_v2));
        let _1524_v7;
        _1524_v7 = _module.D1.create_DC3(_dafny.Seq.of(_1522_v5));
        let _1525_v8;
        _1525_v8 = _dafny.MultiSet.fromElements(_1523_v6, (_1524_v7).dtor_cf8);
        let _1526_v9;
        _1526_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_this).f6)).length),(_this).f6);
        let _1527_v10;
        _1527_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1525_v8,(_1526_v9).update(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(885)), function (_1528_i3) {
          return new _dafny.CodePoint('t'.codePointAt(0));
        }))).cardinality()), (_this).f6));
        let _1529_v11;
        _1529_v11 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f6),false);
        let _1530_v12;
        _1530_v12 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1529_v11).length),(_this).f6), _1526_v9);
        let _index244 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_1514_v1).length));
        (_1514_v1)[_index244] = new BigNumber(((((_1527_v10).contains(_1525_v8)) ? ((_1527_v10).get(_1525_v8)) : ((_1530_v12)[_module.__default.safeIndex(new BigNumber(-8), new BigNumber((_1530_v12).length))]))).length);
        let _1531_v13;
        _1531_v13 = new _dafny.CodePoint('n'.codePointAt(0));
        _1531_v13 = _1531_v13;
      }
      let _1532_v14;
      _1532_v14 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f6);
      if ((new BigNumber((_1532_v14).length)).isLessThan(p0)) {
        let _1533_v15;
        _1533_v15 = new _dafny.CodePoint('t'.codePointAt(0));
        let _1534_v16;
        _1534_v16 = _dafny.Seq.of(p0);
        let _1535_v17;
        _1535_v17 = _dafny.Set.fromElements(p2);
        let _1536_v18;
        _1536_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1535_v17).length),new BigNumber((_dafny.Seq.of(p0)).length));
        let _1537_v19;
        _1537_v19 = _dafny.Seq.of((_1534_v16)[_module.__default.safeIndex((((_1536_v18).contains((_1534_v16)[_module.__default.safeIndex(p0, new BigNumber((_1534_v16).length))])) ? ((_1536_v18).get((_1534_v16)[_module.__default.safeIndex(p0, new BigNumber((_1534_v16).length))])) : (p0)), new BigNumber((_1534_v16).length))], p2);
        let _1538_v20;
        _1538_v20 = _dafny.Seq.of(_1537_v19);
        let _1539_v21;
        let _nw204 = Array((new BigNumber(29)).toNumber());
        _nw204[(_dafny.ZERO).toNumber()] = (_this).f6;
        _nw204[(_dafny.ONE).toNumber()] = ((_this).f6) || (!((_this).f6));
        _nw204[(new BigNumber(2)).toNumber()] = true;
        _nw204[(new BigNumber(3)).toNumber()] = (_this).f6;
        _nw204[(new BigNumber(4)).toNumber()] = (_this).f6;
        _nw204[(new BigNumber(5)).toNumber()] = !((_this).f6);
        _nw204[(new BigNumber(6)).toNumber()] = (p0).isLessThan(p0);
        _nw204[(new BigNumber(7)).toNumber()] = (_this).f6;
        _nw204[(new BigNumber(8)).toNumber()] = (_this).fm6(_dafny.Seq.of(_1533_v15, _1533_v15, _1533_v15, _1533_v15, _1533_v15), (_this).f6, globalState);
        _nw204[(new BigNumber(9)).toNumber()] = (_this).f6;
        _nw204[(new BigNumber(10)).toNumber()] = !((p2).isLessThanOrEqualTo(new BigNumber((_1538_v20).length)));
        _nw204[(new BigNumber(11)).toNumber()] = (((_this).f6) ? ((_this).f6) : ((_this).f6));
        _nw204[(new BigNumber(12)).toNumber()] = (_this).f6;
        _nw204[(new BigNumber(13)).toNumber()] = false;
        _nw204[(new BigNumber(14)).toNumber()] = (_this).f6;
        _nw204[(new BigNumber(15)).toNumber()] = (!(!((_this).f6))) && ((_this).f6);
        _nw204[(new BigNumber(16)).toNumber()] = (_this).f6;
        _nw204[(new BigNumber(17)).toNumber()] = (_this).f6;
        _nw204[(new BigNumber(18)).toNumber()] = (_this).f6;
        _nw204[(new BigNumber(19)).toNumber()] = (_this).f6;
        _nw204[(new BigNumber(20)).toNumber()] = (_this).f6;
        _nw204[(new BigNumber(21)).toNumber()] = !((_this).f6) || ((_this).f6);
        _nw204[(new BigNumber(22)).toNumber()] = (_this).f6;
        _nw204[(new BigNumber(23)).toNumber()] = (_this).f6;
        _nw204[(new BigNumber(24)).toNumber()] = (_this).f6;
        _nw204[(new BigNumber(25)).toNumber()] = (_this).f6;
        _nw204[(new BigNumber(26)).toNumber()] = false;
        _nw204[(new BigNumber(27)).toNumber()] = (_this).f6;
        _nw204[(new BigNumber(28)).toNumber()] = !((_this).f6) || ((_this).f6);
        _1539_v21 = _nw204;
        let _index245 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_1539_v21).length));
        (_1539_v21)[_index245] = (_this).f6;
        let _index246 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_1539_v21).length));
        (_1539_v21)[_index246] = (_this).f6;
        let _1540_v22;
        _1540_v22 = _dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0)), _1533_v15);
        let _index247 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_1539_v21).length));
        let _rhs191 = _1540_v22;
        let _rhs192 = !((((_this).f6) ? (p2) : ((_dafny.ZERO).minus(p2)))).isEqualTo(p0);
        let _lhs119 = _1539_v21;
        let _lhs120 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_1539_v21).length));
        r2 = _rhs191;
        _lhs119[_lhs120] = _rhs192;
        let _1541_v23;
        _1541_v23 = new BigNumber(284);
        let _1542_v24;
        _1542_v24 = _dafny.Set.fromElements(_module.__default.fm8(_1533_v15, globalState));
        _1541_v23 = (new BigNumber((_1542_v24).length)).minus(_module.__default.fm0(_1541_v23, globalState));
        let _index248 = _module.__default.safeIndex(new BigNumber(551), new BigNumber((_1514_v1).length));
        (_1514_v1)[_index248] = new BigNumber((_1540_v22).length);
        let _index249 = _module.__default.safeIndex(new BigNumber(551), new BigNumber((_1514_v1).length));
        (_1514_v1)[_index249] = (_dafny.ZERO).minus((p2).minus((p2).minus(p2)));
        let _1543_v25;
        _1543_v25 = _dafny.MultiSet.fromElements((_1539_v21)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_1539_v21).length))], (_1539_v21)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_1539_v21).length))]);
        _1541_v23 = ((_1514_v1)[_module.__default.safeIndex(new BigNumber(551), new BigNumber((_1514_v1).length))]).minus((p2).multipliedBy(new BigNumber(((_1543_v25).update((_1539_v21)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_1539_v21).length))], _module.__default.abs(p2))).cardinality())));
      } else {
        let _1544_v26;
        _1544_v26 = _dafny.Seq.UnicodeFromString("jtgco");
        let _index250 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_1513_v0).length));
        (_1513_v0)[_index250] = _dafny.Seq.Concat(_1544_v26, _1544_v26);
        let _1545_v27;
        let _nw205 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.of());
        _1545_v27 = _nw205;
        let _1546_v28;
        _1546_v28 = new BigNumber(260);
        let _index251 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_1513_v0).length));
        let _rhs193 = false;
        let _rhs194 = _dafny.Seq.Concat(_1544_v26, _1544_v26);
        let _rhs195 = _1545_v27;
        let _rhs196 = (_dafny.ZERO).minus(p2);
        let _lhs121 = _1513_v0;
        let _lhs122 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_1513_v0).length));
        r0 = _rhs193;
        _lhs121[_lhs122] = _rhs194;
        _1545_v27 = _rhs195;
        _1546_v28 = _rhs196;
        let _1547_v29;
        _1547_v29 = _module.D2.create_DC5(_1544_v26);
        let _1548_v30;
        _1548_v30 = new _dafny.CodePoint('x'.codePointAt(0));
        if ((_this).fm6((_1547_v29).dtor_cf13, _dafny.Seq.contains(_1544_v26, _1548_v30), globalState)) {
          let _1549_v31;
          _1549_v31 = _module.D2.create_DC6(_1548_v30, _1546_v28);
          let _1550_v32;
          _1550_v32 = _module.D2.create_DC8(_module.D2.create_DC5(_1544_v26));
          let _1551_v33;
          _1551_v33 = _dafny.Map.Empty.slice().updateUnsafe(false,(_1513_v0)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_1513_v0).length))]);
          let _1552_v34;
          _1552_v34 = _dafny.Seq.of(_1550_v32, _1550_v32);
          _1549_v31 = _module.__default.fm9(p0, ((_this).f6) && ((_this).f6), new BigNumber((_dafny.Seq.of(_dafny.Seq.of(_1550_v32), _dafny.Seq.of(_module.D2.create_DC8(_module.D2.create_DC6(_1548_v30, new BigNumber((_1551_v33).length)))), _1552_v34)).length), globalState);
          let _index252 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_1513_v0).length));
          (_1513_v0)[_index252] = (_1513_v0)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_1513_v0).length))];
          let _1553_v35;
          let _nw206 = Array((new BigNumber(24)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1553_v35 = _nw206;
          _1553_v35 = _1553_v35;
          _1546_v28 = (p0).plus(p2);
          _1546_v28 = p2;
        } else {
          let _index253 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_1514_v1).length));
          (_1514_v1)[_index253] = _module.__default.safeModuloInt(new BigNumber(854), _1546_v28);
          let _index254 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_1514_v1).length));
          (_1514_v1)[_index254] = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(604)), ((_1554_v30) => function (_1555_i4) {
            return _1554_v30;
          })(_1548_v30))).length)).plus((_dafny.ZERO).minus(_1546_v28));
          r0 = (_this).f6;
          let _index255 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_1514_v1).length));
          (_1514_v1)[_index255] = (new BigNumber(390)).plus((_1514_v1)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_1514_v1).length))]);
          r0 = !(p2).isEqualTo(_1546_v28);
          let _1556_v36;
          let _nw207 = Array((new BigNumber(26)).toNumber()).fill(false);
          _1556_v36 = _nw207;
          let _index256 = _module.__default.safeIndex(new BigNumber(197), new BigNumber((_1556_v36).length));
          (_1556_v36)[_index256] = (p1)[_module.__default.safeIndex(new BigNumber((_1544_v26).length), new BigNumber((p1).length))];
          let _1557_v37;
          _1557_v37 = _dafny.Map.Empty.slice().updateUnsafe(p1,((_this).f6) || (!((_this).f6)));
          let _1558_v38;
          _1558_v38 = _module.D2.create_DC6(_1548_v30, _1546_v28);
          let _1559_v39;
          _1559_v39 = _module.D2.create_DC8(_1558_v38);
          let _pat_let_tv102 = _1558_v38;
          let _index257 = _module.__default.safeIndex(new BigNumber(197), new BigNumber((_1556_v36).length));
          let _index258 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_1514_v1).length));
          let _rhs197 = ((p2).plus(p2)).isLessThanOrEqualTo((_dafny.ZERO).minus(p2));
          let _rhs198 = (((_1532_v14).contains(new BigNumber(915))) ? ((_1532_v14).get(new BigNumber(915))) : ((_this).f6));
          let _rhs199 = (p0).minus((new BigNumber((_dafny.Seq.of(function (_pat_let39_0) {
            return function (_1560_dt__update__tmp_h0) {
              return function (_pat_let40_0) {
                return function (_1561_dt__update_hcf21_h0) {
                  return _module.D2.create_DC8(_1561_dt__update_hcf21_h0);
                }(_pat_let40_0);
              }(_pat_let_tv102);
            }(_pat_let39_0);
          }(_1559_v39), _module.D2.create_DC8(_1558_v38), _1559_v39, _1559_v39, _1559_v39)).length)).minus(p2));
          let _rhs200 = (_1557_v37).update(_dafny.Seq.Concat(p1, p1), ((_this).f6) || ((_this).f6));
          let _lhs123 = _1556_v36;
          let _lhs124 = _module.__default.safeIndex(new BigNumber(197), new BigNumber((_1556_v36).length));
          let _lhs125 = _1514_v1;
          let _lhs126 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_1514_v1).length));
          r0 = _rhs197;
          _lhs123[_lhs124] = _rhs198;
          _lhs125[_lhs126] = _rhs199;
          _1557_v37 = _rhs200;
        }
        _1548_v30 = _module.__default.fm8(_1548_v30, globalState);
        if (true) {
          let _1562_v40;
          let _nw208 = Array((new BigNumber(22)).toNumber()).fill(false);
          _1562_v40 = _nw208;
          (globalState).f1 = ((false) ? (_1562_v40) : (_1562_v40));
          let _index259 = _module.__default.safeIndex(new BigNumber(418), new BigNumber((_1562_v40).length));
          (_1562_v40)[_index259] = (_this).f6;
          let _index260 = _module.__default.safeIndex(new BigNumber(418), new BigNumber((_1562_v40).length));
          (_1562_v40)[_index260] = (_this).f6;
          let _1563_v41;
          _1563_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1546_v28,_1546_v28);
          _1563_v41 = (_1563_v41).update((new BigNumber(-499)).multipliedBy(new BigNumber(((_1513_v0)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_1513_v0).length))]).length)), _1546_v28);
          let _1564_v42;
          _1564_v42 = _dafny.MultiSet.fromElements((_this).f6, (_this).f6);
          let _1565_v43;
          _1565_v43 = _module.D0.create_DC2(p2, (_this).f6, (_this).f6, new BigNumber((_1564_v42).cardinality()));
          let _1566_v44;
          _1566_v44 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
          let _1567_v45;
          _1567_v45 = _dafny.MultiSet.fromElements(new BigNumber(-423), _1546_v28, p2);
          let _1568_v46;
          _1568_v46 = _dafny.Set.fromElements(_module.__default.fm10((_1562_v40)[_module.__default.safeIndex(new BigNumber(418), new BigNumber((_1562_v40).length))], (_dafny.ZERO).minus(_1546_v28), _1565_v43, globalState), (_dafny.MultiSet.fromElements(new BigNumber((_1566_v44).length))).Intersect(_1567_v45), (_1567_v45).Difference(_1567_v45));
          _1568_v46 = (_dafny.Set.fromElements(_1567_v45)).Union(_1568_v46);
          let _index261 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_1513_v0).length));
          (_1513_v0)[_index261] = _dafny.Seq.Concat((_1513_v0)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_1513_v0).length))], _dafny.Seq.Concat((_1513_v0)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_1513_v0).length))], _module.__default.fm11(p0, (_1562_v40)[_module.__default.safeIndex(new BigNumber(418), new BigNumber((_1562_v40).length))], _1546_v28, globalState)));
        } else {
          _1544_v26 = _1544_v26;
          let _1569_v47;
          _1569_v47 = _dafny.Map.Empty.slice().updateUnsafe((_1513_v0)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_1513_v0).length))],_1532_v14);
          let _1570_v48;
          _1570_v48 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_1569_v47);
          let _index262 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_1514_v1).length));
          (_1514_v1)[_index262] = new BigNumber((((((_1570_v48).contains(false)) ? ((_1570_v48).get(false)) : (_dafny.Map.Empty.slice().updateUnsafe((_1513_v0)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_1513_v0).length))],_1532_v14)))).Merge(function () {
            let _coll53 = new _dafny.Map();
            for (const _compr_53 of (_dafny.Set.fromElements(_1544_v26)).Elements) {
              let _1571_v49 = _compr_53;
              if ((_dafny.Set.fromElements(_1544_v26)).contains(_1571_v49)) {
                _coll53.push([_1571_v49,_1532_v14]);
              }
            }
            return _coll53;
          }())).length);
          let _index263 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_1514_v1).length));
          (_1514_v1)[_index263] = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(new BigNumber(308), p2, p2), _module.__default.safeIndex(_1546_v28, new BigNumber((_dafny.Seq.of(new BigNumber(308), p2, p2)).length)), p0)).length), _1546_v28);
          let _1572_v50;
          _1572_v50 = _dafny.Set.fromElements((_1513_v0)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_1513_v0).length))], _1544_v26);
          let _1573_v51;
          let _nw209 = Array((new BigNumber(22)).toNumber()).fill(false);
          _1573_v51 = _nw209;
          let _index264 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_1573_v51).length));
          (_1573_v51)[_index264] = !((_this).f6) || ((_this).f6);
          let _1574_v52;
          _1574_v52 = _dafny.Set.fromElements((_this).f6, (_this).f6);
          let _1575_v53;
          _1575_v53 = _dafny.Set.fromElements(_1574_v52, (_1574_v52).Intersect(_1574_v52));
          let _index265 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_1514_v1).length));
          let _index266 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_1573_v51).length));
          let _rhs201 = (_dafny.ZERO).minus(new BigNumber((_1575_v53).length));
          let _rhs202 = _1572_v50;
          let _rhs203 = (_this).f6;
          let _rhs204 = _dafny.Seq.Concat(_1544_v26, _1544_v26);
          let _lhs127 = _1514_v1;
          let _lhs128 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_1514_v1).length));
          let _lhs129 = _1573_v51;
          let _lhs130 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_1573_v51).length));
          _lhs127[_lhs128] = _rhs201;
          _1572_v50 = _rhs202;
          _lhs129[_lhs130] = _rhs203;
          _1544_v26 = _rhs204;
          let _1576_v54;
          _1576_v54 = _dafny.Seq.of((_1514_v1)[_module.__default.safeIndex(new BigNumber(823), new BigNumber((_1514_v1).length))]);
          let _1577_v55;
          _1577_v55 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p1).length),new BigNumber(848));
          let _1578_v56;
          _1578_v56 = _dafny.Set.fromElements(_1577_v55, _1577_v55);
          let _1579_v57;
          _1579_v57 = _dafny.MultiSet.fromElements((((_this).f6) ? ((_this).f6) : (!((_1573_v51)[_module.__default.safeIndex(new BigNumber(836), new BigNumber((_1573_v51).length))]))));
          let _1580_v59;
          _1580_v59 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll54 = new _dafny.Map();
            for (const _compr_54 of _dafny.IntegerRange(new BigNumber(393), new BigNumber(951))) {
              let _1581_v58 = _compr_54;
              if (((new BigNumber(393)).isLessThanOrEqualTo(_1581_v58)) && ((_1581_v58).isLessThan(new BigNumber(951)))) {
                _coll54.push([(_1581_v58).minus((_1514_v1)[_module.__default.safeIndex(new BigNumber(823), new BigNumber((_1514_v1).length))]),(_1514_v1)[_module.__default.safeIndex(new BigNumber(823), new BigNumber((_1514_v1).length))]]);
              }
            }
            return _coll54;
          }(),(_1573_v51)[_module.__default.safeIndex(new BigNumber(836), new BigNumber((_1573_v51).length))]);
          let _index267 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_1573_v51).length));
          let _index268 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_1514_v1).length));
          let _rhs205 = _1576_v54;
          let _rhs206 = (_1578_v56).Difference(function () {
            let _coll55 = new _dafny.Set();
            for (const _compr_55 of (_1580_v59).Keys.Elements) {
              let _1582_v60 = _compr_55;
              if ((_1580_v59).contains(_1582_v60)) {
                _coll55.add(_1582_v60);
              }
            }
            return _coll55;
          }());
          let _rhs207 = (_1579_v57).update((_this).f6, _module.__default.abs(p0));
          let _rhs208 = (_1573_v51)[_module.__default.safeIndex(new BigNumber(836), new BigNumber((_1573_v51).length))];
          let _rhs209 = p0;
          let _lhs131 = _1573_v51;
          let _lhs132 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_1573_v51).length));
          let _lhs133 = _1514_v1;
          let _lhs134 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_1514_v1).length));
          _1576_v54 = _rhs205;
          _1578_v56 = _rhs206;
          _1579_v57 = _rhs207;
          _lhs131[_lhs132] = _rhs208;
          _lhs133[_lhs134] = _rhs209;
          let _1583_v61;
          let _nw210 = new _module.C0();
          _nw210.__ctor();
          _1583_v61 = _nw210;
        }
        r0 = (_this).f6;
      }
      let _1584_v62;
      _1584_v62 = new _dafny.CodePoint('c'.codePointAt(0));
      let _1585_v63;
      _1585_v63 = _dafny.Seq.of(_1584_v62, _1584_v62, _1584_v62, _1584_v62, _1584_v62);
      let _1586_v64;
      _1586_v64 = _module.D2.create_DC5(_1585_v63);
      let _1587_v65;
      _1587_v65 = _dafny.Seq.of(_1585_v63, _dafny.Seq.Create(_module.__default.abs(new BigNumber(925)), ((_1588_v62) => function (_1589_i5) {
        return _1588_v62;
      })(_1584_v62)), _1585_v63);
      let _1590_v66;
      let _nw211 = Array((new BigNumber(10)).toNumber());
      _nw211[(_dafny.ZERO).toNumber()] = _1586_v64;
      _nw211[(_dafny.ONE).toNumber()] = _1586_v64;
      _nw211[(new BigNumber(2)).toNumber()] = _1586_v64;
      _nw211[(new BigNumber(3)).toNumber()] = _1586_v64;
      _nw211[(new BigNumber(4)).toNumber()] = _1586_v64;
      _nw211[(new BigNumber(5)).toNumber()] = _1586_v64;
      _nw211[(new BigNumber(6)).toNumber()] = _1586_v64;
      _nw211[(new BigNumber(7)).toNumber()] = _1586_v64;
      _nw211[(new BigNumber(8)).toNumber()] = _module.D2.create_DC5((_1587_v65)[_module.__default.safeIndex(p0, new BigNumber((_1587_v65).length))]);
      _nw211[(new BigNumber(9)).toNumber()] = _1586_v64;
      _1590_v66 = _nw211;
      let _index269 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_1590_v66).length));
      (_1590_v66)[_index269] = _1586_v64;
      let _1591_v67;
      _1591_v67 = new BigNumber(702);
      let _index270 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_1590_v66).length));
      let _rhs210 = _1586_v64;
      let _rhs211 = p2;
      let _lhs135 = _1590_v66;
      let _lhs136 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_1590_v66).length));
      _lhs135[_lhs136] = _rhs210;
      _1591_v67 = _rhs211;
      let _1592_v68;
      _1592_v68 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(_dafny.ZERO), ((_1593_v62) => function (_1594_i6) {
        return _1593_v62;
      })(_1584_v62)));
      r0 = (_1592_v68).IsSubsetOf(_1592_v68);
      r0 = (_this).f6;
      let _1595_v69;
      let _init45 = function (_1596_i7) {
        return (_this).f6;
      };
      let _nw212 = Array((new BigNumber(20)).toNumber());
      for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw212.length); _i0_45++) {
        _nw212[_i0_45] = _init45(new BigNumber(_i0_45));
      }
      _1595_v69 = _nw212;
      r1 = _1595_v69;
      r2 = (((_this).f6) ? (_dafny.Seq.Concat(_1585_v63, _dafny.Seq.UnicodeFromString("flcs"))) : (_dafny.Seq.UnicodeFromString("tudymb")));
      return [r0, r1, r2];
    }
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _1597_v0;
      _1597_v0 = _dafny.Seq.of((_this).f6);
      let _1598_v1;
      _1598_v1 = _dafny.Set.fromElements(_1597_v0, _1597_v0, _dafny.Seq.Concat(_1597_v0, _1597_v0), _1597_v0);
      _1598_v1 = _1598_v1;
      let _1599_v2;
      _1599_v2 = new BigNumber(482);
      let _hi9 = _1599_v2;
      for (let _1600_i0 = _module.__default.fm0(_1599_v2, globalState); _1600_i0.isLessThan(_hi9); _1600_i0 = _1600_i0.plus(_dafny.ONE)) {
        let _1601_v3;
        let _init46 = ((_1602_p1) => function (_1603_i1) {
          return !(_1602_p1);
        })(p1);
        let _nw213 = Array((new BigNumber(2)).toNumber());
        for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw213.length); _i0_46++) {
          _nw213[_i0_46] = _init46(new BigNumber(_i0_46));
        }
        _1601_v3 = _nw213;
        (globalState).f1 = _1601_v3;
        let _1604_v4;
        _1604_v4 = _dafny.Seq.UnicodeFromString("bhimfjb");
        if (_dafny.areEqual(_dafny.Seq.Concat(_1604_v4, _dafny.Seq.UnicodeFromString("uu")), _1604_v4)) {
          let _1605_v5;
          _1605_v5 = new _dafny.CodePoint('k'.codePointAt(0));
          let _1606_v6;
          _1606_v6 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1597_v0);
          let _1607_v7;
          _1607_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1605_v5,(_this).f6);
          let _1608_v8;
          _1608_v8 = _module.D3.create_DC9(_1607_v7);
          let _1609_v9;
          _1609_v9 = _dafny.Map.Empty.slice().updateUnsafe((_1608_v8).dtor_cf22,(_this).f6);
          let _1610_v10;
          _1610_v10 = _module.D2.create_DC6(_1605_v5, (_dafny.ZERO).minus(new BigNumber((_1609_v9).length)));
          let _1611_v11;
          _1611_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1599_v2,((p1) ? (_module.D2.create_DC6(_1605_v5, new BigNumber((_1606_v6).length))) : (_1610_v10)));
          _1611_v11 = (_1611_v11).update(((!(p1)) ? (_1600_i0) : (_1599_v2)), _1610_v10);
          let _1612_v12;
          _1612_v12 = _dafny.MultiSet.fromElements(true);
          let _1613_v13;
          _1613_v13 = _module.D0.create_DC0((_this).f6);
          let _1614_v14;
          _1614_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1613_v13,_1599_v2);
          let _1615_v15;
          _1615_v15 = _dafny.MultiSet.fromElements(new BigNumber((_1612_v12).cardinality()), new BigNumber((_1604_v4).length), _1600_i0, new BigNumber((_1614_v14).length), new BigNumber(956));
          r1 = (_1600_i0).minus((((_1615_v15).contains(_1599_v2)) ? ((_1615_v15).get(_1599_v2)) : ((_dafny.ZERO).minus(_1600_i0))));
          let _index271 = _module.__default.safeIndex(new BigNumber(60), new BigNumber((_1601_v3).length));
          (_1601_v3)[_index271] = p1;
          let _index272 = _module.__default.safeIndex(new BigNumber(60), new BigNumber((_1601_v3).length));
          (_1601_v3)[_index272] = ((_1599_v2).multipliedBy(_1599_v2)).isLessThanOrEqualTo(_1600_i0);
          let _index273 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((p0).length));
          (p0)[_index273] = new BigNumber(-736);
          let _1616_v16;
          _1616_v16 = _dafny.Seq.of(_1599_v2, _1600_i0);
          let _index274 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((p0).length));
          (p0)[_index274] = (_1616_v16)[_module.__default.safeIndex(_1599_v2, new BigNumber((_1616_v16).length))];
          r0 = p1;
        } else {
          let _1617_v17;
          _1617_v17 = _dafny.MultiSet.fromElements((_this).f6);
          _1617_v17 = ((_1617_v17).update(!((_this).f6), _module.__default.abs(_1600_i0))).Intersect((_1617_v17).update(p1, _module.__default.abs(new BigNumber(491))));
          r1 = (_dafny.ZERO).minus(_1599_v2);
          _1599_v2 = (_1599_v2).minus(_1600_i0);
          _1599_v2 = (new BigNumber((_1617_v17).cardinality())).minus(_1600_i0);
          let _1618_v18;
          _1618_v18 = new _dafny.CodePoint('e'.codePointAt(0));
          let _1619_v19;
          _1619_v19 = _dafny.Seq.of((_1600_i0).minus(_1600_i0), (_dafny.ZERO).minus(_1599_v2), (_dafny.ZERO).minus(_1600_i0));
          let _index275 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_1601_v3).length));
          (_1601_v3)[_index275] = p1;
          let _index276 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_1601_v3).length));
          let _rhs212 = _1618_v18;
          let _rhs213 = _1619_v19;
          let _rhs214 = p1;
          let _rhs215 = ((p1) ? ((_this).f6) : ((_this).f6));
          let _lhs137 = _1601_v3;
          let _lhs138 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_1601_v3).length));
          _1618_v18 = _rhs212;
          _1619_v19 = _rhs213;
          r0 = _rhs214;
          _lhs137[_lhs138] = _rhs215;
        }
        _1604_v4 = _dafny.Seq.UnicodeFromString("uca");
        let _1620_v20;
        let _nw214 = new _module.C0();
        _nw214.__ctor();
        _1620_v20 = _nw214;
      }
      let _1621_v21;
      _1621_v21 = _dafny.MultiSet.fromElements(_1599_v2, new BigNumber(19), _1599_v2, _1599_v2, _1599_v2);
      let _1622_v22;
      _1622_v22 = _dafny.Seq.of(_1621_v21, _1621_v21);
      if (!((_1622_v22)[_module.__default.safeIndex(_1599_v2, new BigNumber((_1622_v22).length))]).equals(_1621_v21)) {
        _1599_v2 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(_1599_v2, new BigNumber((_dafny.Seq.UnicodeFromString("ppwwf")).length)), _1599_v2);
        let _1623_v23;
        _1623_v23 = _dafny.Seq.of(_1599_v2, new BigNumber(792));
        let _1624_v24;
        _1624_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_1623_v23)).cardinality()),(_this).f6);
        _1624_v24 = (_1624_v24).update(_1599_v2, !(p1));
        _1599_v2 = _1599_v2;
        let _1625_v25;
        let _nw215 = Array((new BigNumber(15)).toNumber()).fill(false);
        _1625_v25 = _nw215;
        let _1626_v26;
        _1626_v26 = _dafny.Seq.of(_1625_v25);
        let _1627_v27;
        _1627_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1599_v2,(_1626_v26)[_module.__default.safeIndex(new BigNumber((_1623_v23).length), new BigNumber((_1626_v26).length))]);
        let _rhs216 = (((_1627_v27).contains(_1599_v2)) ? ((_1627_v27).get(_1599_v2)) : (_1625_v25));
        let _rhs217 = (_1599_v2).plus(new BigNumber((_1623_v23).length));
        let _rhs218 = (_this).f6;
        let _rhs219 = _1625_v25;
        let _lhs139 = globalState;
        let _lhs140 = globalState;
        _lhs139.f1 = _rhs216;
        r1 = _rhs217;
        r0 = _rhs218;
        _lhs140.f1 = _rhs219;
        let _1628_v28;
        let _init47 = ((_1629_v2) => function (_1630_i2) {
          return _dafny.Map.Empty.slice().updateUnsafe(false,_1629_v2);
        })(_1599_v2);
        let _nw216 = Array((new BigNumber(16)).toNumber());
        for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw216.length); _i0_47++) {
          _nw216[_i0_47] = _init47(new BigNumber(_i0_47));
        }
        _1628_v28 = _nw216;
        let _1631_v29;
        let _nw217 = new _module.C1();
        _nw217.__ctor(_1628_v28, (_this).f6, _1625_v25);
        _1631_v29 = _nw217;
        _1631_v29 = _1631_v29;
      } else {
        let _index277 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length));
        (p0)[_index277] = _1599_v2;
        let _1632_v30;
        _1632_v30 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_1597_v0, _1597_v0)).length),_1599_v2);
        let _1633_v31;
        _1633_v31 = _dafny.Set.fromElements(p0, p0);
        let _1634_v32;
        _1634_v32 = _module.D5.create_DC17(new BigNumber((_1633_v31).length));
        let _pat_let_tv103 = _1599_v2;
        let _pat_let_tv104 = _1599_v2;
        let _1635_v33;
        let _nw218 = Array((new BigNumber(13)).toNumber());
        _nw218[(_dafny.ZERO).toNumber()] = _1634_v32;
        _nw218[(_dafny.ONE).toNumber()] = _1634_v32;
        _nw218[(new BigNumber(2)).toNumber()] = _1634_v32;
        _nw218[(new BigNumber(3)).toNumber()] = _1634_v32;
        _nw218[(new BigNumber(4)).toNumber()] = _module.D5.create_DC17(_1599_v2);
        _nw218[(new BigNumber(5)).toNumber()] = _module.D5.create_DC17((_dafny.ZERO).minus(new BigNumber((_1597_v0).length)));
        _nw218[(new BigNumber(6)).toNumber()] = _1634_v32;
        _nw218[(new BigNumber(7)).toNumber()] = _1634_v32;
        _nw218[(new BigNumber(8)).toNumber()] = _module.__default.fm27(_1599_v2, _1599_v2, globalState);
        _nw218[(new BigNumber(9)).toNumber()] = _module.__default.fm27(_1599_v2, _1599_v2, globalState);
        _nw218[(new BigNumber(10)).toNumber()] = _1634_v32;
        _nw218[(new BigNumber(11)).toNumber()] = function (_pat_let41_0) {
          return function (_1636_dt__update__tmp_h0) {
            return function (_pat_let42_0) {
              return function (_1637_dt__update_hcf35_h0) {
                return _module.D5.create_DC17(_1637_dt__update_hcf35_h0);
              }(_pat_let42_0);
            }(_pat_let_tv103);
          }(_pat_let41_0);
        }(_1634_v32);
        _nw218[(new BigNumber(12)).toNumber()] = function (_pat_let43_0) {
          return function (_1638_dt__update__tmp_h1) {
            return function (_pat_let44_0) {
              return function (_1639_dt__update_hcf35_h1) {
                return _module.D5.create_DC17(_1639_dt__update_hcf35_h1);
              }(_pat_let44_0);
            }(_pat_let_tv104);
          }(_pat_let43_0);
        }(_1634_v32);
        _1635_v33 = _nw218;
        let _index278 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_1635_v33).length));
        (_1635_v33)[_index278] = _1634_v32;
        let _index279 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length));
        let _index280 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_1635_v33).length));
        let _rhs220 = _1599_v2;
        let _rhs221 = _1632_v30;
        let _rhs222 = _1634_v32;
        let _lhs141 = p0;
        let _lhs142 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length));
        let _lhs143 = _1635_v33;
        let _lhs144 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_1635_v33).length));
        _lhs141[_lhs142] = _rhs220;
        _1632_v30 = _rhs221;
        _lhs143[_lhs144] = _rhs222;
        if ((_this).f6) {
          let _1640_v34;
          let _nw219 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1640_v34 = _nw219;
          _1640_v34 = _1640_v34;
          let _1641_v35;
          _1641_v35 = new _dafny.CodePoint('k'.codePointAt(0));
          _1641_v35 = _1641_v35;
          let _index281 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length));
          (p0)[_index281] = _1599_v2;
          let _1642_v36;
          _1642_v36 = _dafny.Seq.of(_1599_v2);
          let _1643_v37;
          _1643_v37 = _module.D0.create_DC0(p1);
          let _1644_v38;
          let _nw220 = Array((new BigNumber(15)).toNumber());
          _nw220[(_dafny.ZERO).toNumber()] = !(!(p1));
          _nw220[(_dafny.ONE).toNumber()] = !((_this).f6);
          _nw220[(new BigNumber(2)).toNumber()] = (_this).f6;
          _nw220[(new BigNumber(3)).toNumber()] = (_this).f6;
          _nw220[(new BigNumber(4)).toNumber()] = p1;
          _nw220[(new BigNumber(5)).toNumber()] = (_this).f6;
          _nw220[(new BigNumber(6)).toNumber()] = !(p1) || ((_this).f6);
          _nw220[(new BigNumber(7)).toNumber()] = (new BigNumber((_1642_v36).length)).isLessThan((p0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length))]);
          _nw220[(new BigNumber(8)).toNumber()] = (_this).f6;
          _nw220[(new BigNumber(9)).toNumber()] = (_1643_v37).dtor_cf0;
          _nw220[(new BigNumber(10)).toNumber()] = p1;
          _nw220[(new BigNumber(11)).toNumber()] = !((_this).f6);
          _nw220[(new BigNumber(12)).toNumber()] = !(p1) || ((_this).f6);
          _nw220[(new BigNumber(13)).toNumber()] = (_1621_v21).IsDisjointFrom(_1621_v21);
          _nw220[(new BigNumber(14)).toNumber()] = (_this).f6;
          _1644_v38 = _nw220;
          let _index282 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_1644_v38).length));
          (_1644_v38)[_index282] = (_this).f6;
          let _1645_v39;
          let _nw221 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.of());
          _1645_v39 = _nw221;
          let _1646_v40;
          _1646_v40 = _module.D10.create_DC32(_dafny.MultiSet.fromElements(_1644_v38), (p0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length))], (p0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length))]);
          let _1647_v41;
          _1647_v41 = _dafny.MultiSet.fromElements(_1644_v38, _1644_v38);
          let _pat_let_tv105 = _1647_v41;
          let _1648_v42;
          _1648_v42 = _dafny.MultiSet.fromElements(function (_pat_let45_0) {
            return function (_1649_dt__update__tmp_h2) {
              return function (_pat_let46_0) {
                return function (_1650_dt__update_hcf58_h0) {
                  return _module.D10.create_DC32(_1650_dt__update_hcf58_h0, (_1649_dt__update__tmp_h2).dtor_cf59, (_1649_dt__update__tmp_h2).dtor_cf60);
                }(_pat_let46_0);
              }(_pat_let_tv105);
            }(_pat_let45_0);
          }(_1646_v40));
          let _index283 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_1645_v39).length));
          (_1645_v39)[_index283] = _dafny.Seq.update(_dafny.Seq.update(_1597_v0, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1648_v42).cardinality())), new BigNumber((_1597_v0).length)), p1), _module.__default.safeIndex(_1599_v2, new BigNumber((_dafny.Seq.update(_1597_v0, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1648_v42).cardinality())), new BigNumber((_1597_v0).length)), p1)).length)), p1);
          let _1651_v43;
          _1651_v43 = _dafny.Seq.UnicodeFromString("hpkbj");
          let _1652_v44;
          let _init48 = ((_1653_v2) => function (_1654_i3) {
            return _module.__default.safeModuloInt(_1654_i3, _1653_v2);
          })(_1599_v2);
          let _nw222 = Array((new BigNumber(18)).toNumber());
          for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw222.length); _i0_48++) {
            _nw222[_i0_48] = _init48(new BigNumber(_i0_48));
          }
          _1652_v44 = _nw222;
          let _1655_v45;
          _1655_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_1652_v44);
          let _1656_v46;
          let _nw223 = Array((new BigNumber(18)).toNumber());
          _nw223[(_dafny.ZERO).toNumber()] = _1652_v44;
          _nw223[(_dafny.ONE).toNumber()] = _1652_v44;
          _nw223[(new BigNumber(2)).toNumber()] = _1652_v44;
          _nw223[(new BigNumber(3)).toNumber()] = _1652_v44;
          _nw223[(new BigNumber(4)).toNumber()] = _1652_v44;
          _nw223[(new BigNumber(5)).toNumber()] = _1652_v44;
          _nw223[(new BigNumber(6)).toNumber()] = _1652_v44;
          _nw223[(new BigNumber(7)).toNumber()] = _1652_v44;
          _nw223[(new BigNumber(8)).toNumber()] = ((false) ? (_1652_v44) : (_1652_v44));
          _nw223[(new BigNumber(9)).toNumber()] = _1652_v44;
          _nw223[(new BigNumber(10)).toNumber()] = (((_1655_v45).contains(!(p1))) ? ((_1655_v45).get(!(p1))) : (_1652_v44));
          _nw223[(new BigNumber(11)).toNumber()] = _1652_v44;
          _nw223[(new BigNumber(12)).toNumber()] = ((p1) ? (_1652_v44) : (_1652_v44));
          _nw223[(new BigNumber(13)).toNumber()] = _1652_v44;
          _nw223[(new BigNumber(14)).toNumber()] = _1652_v44;
          _nw223[(new BigNumber(15)).toNumber()] = _1652_v44;
          _nw223[(new BigNumber(16)).toNumber()] = _1652_v44;
          _nw223[(new BigNumber(17)).toNumber()] = _1652_v44;
          _1656_v46 = _nw223;
          let _index284 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_1656_v46).length));
          (_1656_v46)[_index284] = _1652_v44;
          let _1657_v47;
          _1657_v47 = _1652_v44;
          let _index285 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_1644_v38).length));
          let _index286 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length));
          let _index287 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_1645_v39).length));
          let _index288 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_1656_v46).length));
          let _rhs223 = ((_dafny.Seq.IsProperPrefixOf(_1642_v36, _1642_v36)) ? (false) : ((_this).f6));
          let _rhs224 = (p0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length))];
          let _rhs225 = _1597_v0;
          let _rhs226 = _dafny.Seq.UnicodeFromString("ghblvvi");
          let _rhs227 = (_1657_v47);
          let _lhs145 = _1644_v38;
          let _lhs146 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_1644_v38).length));
          let _lhs147 = p0;
          let _lhs148 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length));
          let _lhs149 = _1645_v39;
          let _lhs150 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_1645_v39).length));
          let _lhs151 = _1656_v46;
          let _lhs152 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_1656_v46).length));
          _lhs145[_lhs146] = _rhs223;
          _lhs147[_lhs148] = _rhs224;
          _lhs149[_lhs150] = _rhs225;
          _1651_v43 = _rhs226;
          _lhs151[_lhs152] = _rhs227;
          let _index289 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_1644_v38).length));
          (_1644_v38)[_index289] = (_1597_v0)[_module.__default.safeIndex((p0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length))], new BigNumber((_1597_v0).length))];
        } else {
          r2 = p1;
          let _1658_v48;
          _1658_v48 = _dafny.Seq.UnicodeFromString("ask");
          let _1659_v49;
          _1659_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1658_v48,_1599_v2);
          _1659_v49 = _1659_v49;
          let _1660_v50;
          let _nw224 = Array((new BigNumber(27)).toNumber()).fill(false);
          _1660_v50 = _nw224;
          let _1661_v51;
          let _nw225 = new _module.C3();
          _nw225.__ctor(_1660_v50);
          _1661_v51 = _nw225;
          let _1662_v52;
          _1662_v52 = new _dafny.CodePoint('h'.codePointAt(0));
          _1662_v52 = _1662_v52;
          let _index290 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length));
          (p0)[_index290] = new BigNumber(((((_this).f6) ? (_1658_v48) : (_1658_v48))).length);
        }
        let _1663_v53;
        let _init49 = function (_1664_i4) {
          return (_this).f6;
        };
        let _nw226 = Array((new BigNumber(21)).toNumber());
        for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw226.length); _i0_49++) {
          _nw226[_i0_49] = _init49(new BigNumber(_i0_49));
        }
        _1663_v53 = _nw226;
        let _1665_v54;
        let _nw227 = new _module.C3();
        _nw227.__ctor(_1663_v53);
        _1665_v54 = _nw227;
        let _1666_v55;
        let _init50 = ((_1667_p1, _1668_p0) => function (_1669_i5) {
          return _dafny.Map.Empty.slice().updateUnsafe(_1667_p1,(_1668_p0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_1668_p0).length))]);
        })(p1, p0);
        let _nw228 = Array((new BigNumber(23)).toNumber());
        for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw228.length); _i0_50++) {
          _nw228[_i0_50] = _init50(new BigNumber(_i0_50));
        }
        _1666_v55 = _nw228;
        let _1670_v56;
        let _nw229 = new _module.C1();
        _nw229.__ctor(_1666_v55, p1, _1663_v53);
        _1670_v56 = _nw229;
        let _1671_v57;
        _1671_v57 = _module.D10.create_DC31(_1670_v56);
        let _1672_v58;
        _1672_v58 = _dafny.Seq.of((_1671_v57).dtor_cf57);
        let _1673_v59;
        _1673_v59 = new _dafny.CodePoint('w'.codePointAt(0));
        let _1674_v60;
        _1674_v60 = _dafny.Set.fromElements(p1);
        let _rhs228 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1672_v58, _module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("yilayj"), _module.__default.safeIndex((p0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length))], new BigNumber((_dafny.Seq.UnicodeFromString("yilayj")).length)), _1673_v59))).length), new BigNumber((_1672_v58).length)), _1670_v56), _1672_v58)).length);
        let _rhs229 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_1599_v2));
        let _rhs230 = ((p1) ? (new BigNumber(((_1674_v60).Union(_dafny.Set.fromElements(p1, p1, (_this).f6, false))).length)) : (new BigNumber(((_1621_v21).Difference(_dafny.MultiSet.fromElements(_1599_v2, _1599_v2, _1599_v2, new BigNumber(-508)))).cardinality())));
        let _rhs231 = _1665_v54;
        _1599_v2 = _rhs228;
        r1 = _rhs229;
        _1599_v2 = _rhs230;
        _1665_v54 = _rhs231;
        let _1675_v62;
        _1675_v62 = _dafny.MultiSet.fromElements(function () {
          let _coll56 = new _dafny.Set();
          for (const _compr_56 of _dafny.IntegerRange(new BigNumber(-757), new BigNumber(981))) {
            let _1676_v61 = _compr_56;
            if (((new BigNumber(-757)).isLessThanOrEqualTo(_1676_v61)) && ((_1676_v61).isLessThan(new BigNumber(981)))) {
              _coll56.add((_1676_v61).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(79)), ((_1677_v59) => function (_1678_i6) {
                return _1677_v59;
              })(_1673_v59))).length))));
            }
          }
          return _coll56;
        }());
        _1675_v62 = _1675_v62;
        let _1679_v63;
        _1679_v63 = _dafny.MultiSet.fromElements(_1673_v59, _1673_v59);
        let _1680_v64;
        _1680_v64 = _module.D4.create_DC13(((p0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length))]).isLessThan(_1599_v2), _1679_v63, (p1) && ((_this).f6));
        let _source23 = _1680_v64;
        if (_source23.is_DC12) {
          let _1681_v65;
          let _init51 = ((_1682_p0) => function (_1683_i7) {
            return _module.__default.safeDivisionInt(_1683_i7, (_1682_p0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_1682_p0).length))]);
          })(p0);
          let _nw230 = Array((new BigNumber(5)).toNumber());
          for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw230.length); _i0_51++) {
            _nw230[_i0_51] = _init51(new BigNumber(_i0_51));
          }
          _1681_v65 = _nw230;
          let _1684_v66;
          _1684_v66 = _dafny.Map.Empty.slice().updateUnsafe(true,_1681_v65);
          let _1685_v67;
          _1685_v67 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_1684_v66).length)),p1);
          let _1686_v68;
          _1686_v68 = _dafny.Seq.of(_1685_v67);
          _1686_v68 = _1686_v68;
          let _1687_v69;
          _1687_v69 = _dafny.Map.Empty.slice().updateUnsafe(p1,true);
          _1687_v69 = _dafny.Map.Empty.slice().updateUnsafe(p1,!(false));
          (globalState).f1 = _1663_v53;
          let _1688_v70;
          _1688_v70 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(!(true)),(_this).f6);
          _1688_v70 = _1688_v70;
        } else if (_source23.is_DC13) {
          let _1689___mcc_h0 = (_source23).cf29;
          let _1690___mcc_h1 = (_source23).cf30;
          let _1691___mcc_h2 = (_source23).cf31;
          let _1692_cf31 = _1691___mcc_h2;
          let _1693_cf30 = _1690___mcc_h1;
          let _1694_cf29 = _1689___mcc_h0;
          let _1695_v71;
          let _nw231 = new _module.C2();
          _nw231.__ctor(_1692_cf31, (p0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length))], _1666_v55, !(_1599_v2).isEqualTo(_1599_v2), _1663_v53);
          _1695_v71 = _nw231;
          let _1696_v72;
          _1696_v72 = _dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)), _1673_v59, new _dafny.CodePoint('u'.codePointAt(0)));
          let _1697_v73;
          _1697_v73 = _dafny.Set.fromElements((p0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length))], _1599_v2, _1599_v2, (_1695_v71).f8);
          let _1698_v74;
          _1698_v74 = _module.D12.create_DC37(_1695_v71.f7, _1597_v0, new BigNumber((_1697_v73).length));
          let _1699_v75;
          _1699_v75 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_1698_v74).dtor_cf69, globalState),_1673_v59);
          let _1700_v76;
          let _nw232 = Array((new BigNumber(16)).toNumber());
          _nw232[(_dafny.ZERO).toNumber()] = _1673_v59;
          _nw232[(_dafny.ONE).toNumber()] = (_1696_v72)[_module.__default.safeIndex((_1695_v71).f8, new BigNumber((_1696_v72).length))];
          _nw232[(new BigNumber(2)).toNumber()] = _1673_v59;
          _nw232[(new BigNumber(3)).toNumber()] = _module.__default.fm22(_1696_v72, globalState);
          _nw232[(new BigNumber(4)).toNumber()] = _1673_v59;
          _nw232[(new BigNumber(5)).toNumber()] = _1673_v59;
          _nw232[(new BigNumber(6)).toNumber()] = (((_1699_v75).contains((_1695_v71).f8)) ? ((_1699_v75).get((_1695_v71).f8)) : (_1673_v59));
          _nw232[(new BigNumber(7)).toNumber()] = _1673_v59;
          _nw232[(new BigNumber(8)).toNumber()] = _1673_v59;
          _nw232[(new BigNumber(9)).toNumber()] = _1673_v59;
          _nw232[(new BigNumber(10)).toNumber()] = _1673_v59;
          _nw232[(new BigNumber(11)).toNumber()] = _1673_v59;
          _nw232[(new BigNumber(12)).toNumber()] = ((_1694_cf29) ? (_1673_v59) : (_1673_v59));
          _nw232[(new BigNumber(13)).toNumber()] = (((_1699_v75).contains(new BigNumber(299))) ? ((_1699_v75).get(new BigNumber(299))) : (_1673_v59));
          _nw232[(new BigNumber(14)).toNumber()] = _1673_v59;
          _nw232[(new BigNumber(15)).toNumber()] = _1673_v59;
          _1700_v76 = _nw232;
          let _index291 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_1700_v76).length));
          (_1700_v76)[_index291] = _1673_v59;
          let _index292 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_1700_v76).length));
          let _rhs232 = _1695_v71;
          let _rhs233 = _1673_v59;
          let _lhs153 = _1700_v76;
          let _lhs154 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_1700_v76).length));
          _1695_v71 = _rhs232;
          _lhs153[_lhs154] = _rhs233;
          let _1701_v77;
          let _nw233 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Map.Empty);
          _1701_v77 = _nw233;
          let _1702_v78;
          _1702_v78 = _module.D6.create_DC19(_1701_v77);
          let _1703_v79;
          _1703_v79 = _module.D6.create_DC21(_1702_v78);
          let _1704_v80;
          let _nw234 = Array((new BigNumber(5)).toNumber());
          _nw234[(_dafny.ZERO).toNumber()] = _1703_v79;
          _nw234[(_dafny.ONE).toNumber()] = _1703_v79;
          _nw234[(new BigNumber(2)).toNumber()] = _1703_v79;
          _nw234[(new BigNumber(3)).toNumber()] = _1703_v79;
          _nw234[(new BigNumber(4)).toNumber()] = _1703_v79;
          _1704_v80 = _nw234;
          let _index293 = _module.__default.safeIndex(new BigNumber(953), new BigNumber((_1704_v80).length));
          (_1704_v80)[_index293] = _module.D6.create_DC21(_1702_v78);
          let _index294 = _module.__default.safeIndex(new BigNumber(953), new BigNumber((_1704_v80).length));
          (_1704_v80)[_index294] = _1703_v79;
          let _1705_v81;
          let _nw235 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          _1705_v81 = _nw235;
          _1705_v81 = _1705_v81;
          let _1706_v82;
          _1706_v82 = _dafny.MultiSet.fromElements(_1663_v53);
          let _1707_v83;
          _1707_v83 = _dafny.Seq.of(_module.D7.create_DC22(_1706_v82));
          let _1708_v84;
          _1708_v84 = _dafny.Map.Empty.slice().updateUnsafe(_1670_v56,_module.__default.fm0(_1599_v2, globalState));
          let _index295 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length));
          let _index296 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_1700_v76).length));
          let _rhs234 = (_1695_v71).f8;
          let _rhs235 = new BigNumber((_1708_v84).length);
          let _rhs236 = p1;
          let _rhs237 = (_1700_v76)[_module.__default.safeIndex(new BigNumber(226), new BigNumber((_1700_v76).length))];
          let _rhs238 = _dafny.Seq.Concat(_1707_v83, _1707_v83);
          let _lhs155 = p0;
          let _lhs156 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length));
          let _lhs157 = _1700_v76;
          let _lhs158 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_1700_v76).length));
          _lhs155[_lhs156] = _rhs234;
          _1599_v2 = _rhs235;
          r2 = _rhs236;
          _lhs157[_lhs158] = _rhs237;
          _1707_v83 = _rhs238;
        } else {
          let _1709___mcc_h3 = (_source23).cf28;
          let _1710_cf28 = _1709___mcc_h3;
          let _1711_v85;
          let _nw236 = new _module.C0();
          _nw236.__ctor();
          _1711_v85 = _nw236;
          let _1712_v86;
          let _nw237 = Array((new BigNumber(5)).toNumber());
          _nw237[(_dafny.ZERO).toNumber()] = _1711_v85;
          _nw237[(_dafny.ONE).toNumber()] = _1711_v85;
          _nw237[(new BigNumber(2)).toNumber()] = _1711_v85;
          _nw237[(new BigNumber(3)).toNumber()] = _1711_v85;
          _nw237[(new BigNumber(4)).toNumber()] = _1711_v85;
          _1712_v86 = _nw237;
          let _index297 = _module.__default.safeIndex(new BigNumber(519), new BigNumber((_1712_v86).length));
          (_1712_v86)[_index297] = _1711_v85;
          let _1713_v87;
          _1713_v87 = _dafny.Seq.UnicodeFromString("jkney");
          let _1714_v88;
          _1714_v88 = _dafny.Map.Empty.slice().updateUnsafe(_1713_v87,(_this).f6);
          let _1715_v89;
          _1715_v89 = _dafny.Map.Empty.slice().updateUnsafe((((_1714_v88).contains(_1713_v87)) ? ((_1714_v88).get(_1713_v87)) : (p1)),(p0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length))]);
          let _1716_v90;
          _1716_v90 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f6);
          let _index298 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length));
          let _index299 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length));
          let _index300 = _module.__default.safeIndex(new BigNumber(519), new BigNumber((_1712_v86).length));
          let _rhs239 = _1599_v2;
          let _rhs240 = _1599_v2;
          let _rhs241 = _1711_v85;
          let _rhs242 = (_dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(184))).Merge(_1715_v89);
          let _rhs243 = new BigNumber(((((p1) ? (_1716_v90) : (_1716_v90))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,p1))).length);
          let _lhs159 = p0;
          let _lhs160 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length));
          let _lhs161 = p0;
          let _lhs162 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length));
          let _lhs163 = _1712_v86;
          let _lhs164 = _module.__default.safeIndex(new BigNumber(519), new BigNumber((_1712_v86).length));
          _lhs159[_lhs160] = _rhs239;
          _lhs161[_lhs162] = _rhs240;
          _lhs163[_lhs164] = _rhs241;
          _1715_v89 = _rhs242;
          _1599_v2 = _rhs243;
          let _1717_v91;
          _1717_v91 = _dafny.MultiSet.fromElements((_this).f6, p1, p1);
          let _1718_v92;
          _1718_v92 = _dafny.Map.Empty.slice().updateUnsafe(_1673_v59,_1717_v91);
          let _1719_v93;
          _1719_v93 = _module.D6.create_DC20(p1, (_this).f6);
          let _1720_v94;
          _1720_v94 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm5((((_1718_v92).contains(_1673_v59)) ? ((_1718_v92).get(_1673_v59)) : (_1717_v91)), globalState),_1719_v93);
          _1720_v94 = (_1720_v94).update((p0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length))], _1719_v93);
          let _1721_v95;
          let _nw238 = Array((new BigNumber(23)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1721_v95 = _nw238;
          let _index301 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_1721_v95).length));
          (_1721_v95)[_index301] = _1673_v59;
          let _index302 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_1721_v95).length));
          (_1721_v95)[_index302] = _1673_v59;
          let _1722_v96;
          _1722_v96 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1713_v87).length),((p1) ? (!((_this).f6)) : ((_this).f6)));
          _1722_v96 = (_1722_v96).update((p0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((p0).length))], (_this).f6);
        }
      }
      let _1723_v97;
      let _nw239 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
      _1723_v97 = _nw239;
      _1723_v97 = p0;
      if ((_this).f6) {
        r0 = (_this).f6;
        let _1724_v98;
        _1724_v98 = _dafny.Seq.UnicodeFromString("upfnqt");
        let _1725_v99;
        _1725_v99 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm11(_1599_v2, (_this).f6, new BigNumber((_1724_v98).length), globalState),_1599_v2);
        let _1726_v100;
        _1726_v100 = _module.D8.create_DC25((_1725_v99).Merge(_1725_v99));
        let _1727_v101;
        _1727_v101 = _module.D13.create_DC41((_dafny.ZERO).minus(_1599_v2));
        let _1728_v102;
        _1728_v102 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1599_v2);
        _1726_v100 = _module.__default.fm35(_1727_v101, !(_1728_v102).contains((_this).f6), globalState);
        r2 = (new BigNumber((_1724_v98).length)).isLessThanOrEqualTo((((_1728_v102).contains(p1)) ? ((_1728_v102).get(p1)) : (_1599_v2)));
        let _rhs244 = new BigNumber(368);
        let _rhs245 = _1597_v0;
        r1 = _rhs244;
        _1597_v0 = _rhs245;
        r0 = (_this).f6;
      } else {
        let _1729_v103;
        _1729_v103 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),p0);
        _1729_v103 = (_1729_v103).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f6,_1723_v97));
        let _1730_v104;
        _1730_v104 = _dafny.Seq.of(new BigNumber((_1597_v0).length));
        let _1731_v105;
        let _init52 = ((_1732_p1, _1733_v2) => function (_1734_i8) {
          return _dafny.Map.Empty.slice().updateUnsafe(_1732_p1,_1733_v2);
        })(p1, _1599_v2);
        let _nw240 = Array((new BigNumber(14)).toNumber());
        for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw240.length); _i0_52++) {
          _nw240[_i0_52] = _init52(new BigNumber(_i0_52));
        }
        _1731_v105 = _nw240;
        let _1735_v106;
        let _nw241 = Array((new BigNumber(27)).toNumber());
        _nw241[(_dafny.ZERO).toNumber()] = p1;
        _nw241[(_dafny.ONE).toNumber()] = (_this).f6;
        _nw241[(new BigNumber(2)).toNumber()] = p1;
        _nw241[(new BigNumber(3)).toNumber()] = (_this).f6;
        _nw241[(new BigNumber(4)).toNumber()] = p1;
        _nw241[(new BigNumber(5)).toNumber()] = p1;
        _nw241[(new BigNumber(6)).toNumber()] = (_this).f6;
        _nw241[(new BigNumber(7)).toNumber()] = (_this).f6;
        _nw241[(new BigNumber(8)).toNumber()] = (_this).f6;
        _nw241[(new BigNumber(9)).toNumber()] = (_this).f6;
        _nw241[(new BigNumber(10)).toNumber()] = (_this).f6;
        _nw241[(new BigNumber(11)).toNumber()] = (_this).f6;
        _nw241[(new BigNumber(12)).toNumber()] = p1;
        _nw241[(new BigNumber(13)).toNumber()] = p1;
        _nw241[(new BigNumber(14)).toNumber()] = p1;
        _nw241[(new BigNumber(15)).toNumber()] = (_this).f6;
        _nw241[(new BigNumber(16)).toNumber()] = _module.__default.fm20(p1, globalState);
        _nw241[(new BigNumber(17)).toNumber()] = p1;
        _nw241[(new BigNumber(18)).toNumber()] = (_this).f6;
        _nw241[(new BigNumber(19)).toNumber()] = (_this).f6;
        _nw241[(new BigNumber(20)).toNumber()] = (_this).f6;
        _nw241[(new BigNumber(21)).toNumber()] = p1;
        _nw241[(new BigNumber(22)).toNumber()] = p1;
        _nw241[(new BigNumber(23)).toNumber()] = (_this).f6;
        _nw241[(new BigNumber(24)).toNumber()] = (_this).f6;
        _nw241[(new BigNumber(25)).toNumber()] = (_this).f6;
        _nw241[(new BigNumber(26)).toNumber()] = p1;
        _1735_v106 = _nw241;
        let _1736_v107;
        let _nw242 = new _module.C2();
        _nw242.__ctor(p1, _1599_v2, _1731_v105, (_this).f6, _1735_v106);
        _1736_v107 = _nw242;
        let _1737_v108;
        _1737_v108 = _dafny.Map.Empty.slice().updateUnsafe(((p1) ? (_1599_v2) : ((_1730_v104)[_module.__default.safeIndex(_1599_v2, new BigNumber((_1730_v104).length))])),_1736_v107);
        _1737_v108 = (_1737_v108).update(_1599_v2, _1736_v107);
        let _1738_v109;
        let _1739_v110;
        let _1740_v111;
        let _out34;
        let _out35;
        let _out36;
        let _outcollector15 = (_this).m2(_1599_v2, _dafny.Seq.of((_this).f6, p1, p1), new BigNumber((_1730_v104).length), globalState);
        _out34 = _outcollector15[0];
        _out35 = _outcollector15[1];
        _out36 = _outcollector15[2];
        _1738_v109 = _out34;
        _1739_v110 = _out35;
        _1740_v111 = _out36;
        let _1741_v112;
        let _init53 = ((_1742_v111) => function (_1743_i9) {
          return _1742_v111;
        })(_1740_v111);
        let _nw243 = Array((_dafny.ONE).toNumber());
        for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw243.length); _i0_53++) {
          _nw243[_i0_53] = _init53(new BigNumber(_i0_53));
        }
        _1741_v112 = _nw243;
        let _1744_v113;
        let _nw244 = Array((_dafny.ONE).toNumber());
        _nw244[(_dafny.ZERO).toNumber()] = _module.__default.fm25(p1, globalState);
        _1744_v113 = _nw244;
        let _1745_v114;
        let _1746_v115;
        let _out37;
        let _out38;
        let _outcollector16 = (_1736_v107).m1(new BigNumber(515), _module.__default.fm22(_1740_v111, globalState), _1741_v112, _1744_v113, globalState);
        _out37 = _outcollector16[0];
        _out38 = _outcollector16[1];
        _1745_v114 = _out37;
        _1746_v115 = _out38;
        if ((_this).f6) {
          r0 = !(p1);
          let _1747_v116;
          let _init54 = ((_1748_v115, _1749_p1) => function (_1750_i10) {
            return _dafny.Map.Empty.slice().updateUnsafe(_1748_v115,_1749_p1);
          })(_1746_v115, p1);
          let _nw245 = Array((new BigNumber(3)).toNumber());
          for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw245.length); _i0_54++) {
            _nw245[_i0_54] = _init54(new BigNumber(_i0_54));
          }
          _1747_v116 = _nw245;
          let _1751_v117;
          _1751_v117 = _module.D6.create_DC19(_1747_v116);
          let _1752_v118;
          _1752_v118 = new _dafny.CodePoint('y'.codePointAt(0));
          let _1753_v119;
          _1753_v119 = _dafny.Map.Empty.slice().updateUnsafe(_1752_v118,(_this).f6);
          let _1754_v120;
          _1754_v120 = _dafny.Map.Empty.slice().updateUnsafe(_1738_v109,_1739_v110);
          let _1755_v121;
          _1755_v121 = _dafny.Seq.of((((_1754_v120).contains(_1736_v107.f7)) ? ((_1754_v120).get(_1736_v107.f7)) : (_1739_v110)));
          let _rhs246 = _dafny.Seq.Concat(_1740_v111, _1740_v111);
          let _rhs247 = _1751_v117;
          let _rhs248 = ((_1736_v107.f7) ? ((_dafny.ZERO).minus((_1736_v107).f8)) : (new BigNumber((_1740_v111).length)));
          let _rhs249 = (_1599_v2).plus(new BigNumber(((_1753_v119).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1752_v118,p1))).length));
          let _rhs250 = (_1755_v121)[_module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1599_v2,_1746_v115)).length), (_dafny.ZERO).minus(new BigNumber((_1621_v21).cardinality()))), new BigNumber((_1755_v121).length))];
          let _lhs165 = globalState;
          _1740_v111 = _rhs246;
          _1751_v117 = _rhs247;
          _1599_v2 = _rhs248;
          r1 = _rhs249;
          _lhs165.f1 = _rhs250;
          let _1756_v122;
          _1756_v122 = _dafny.Map.Empty.slice().updateUnsafe(_1730_v104,(_1736_v107).f8);
          _1756_v122 = (_1756_v122).update(_1730_v104, ((_1736_v107).f8).minus((_dafny.ZERO).minus(new BigNumber(-420))));
          let _1757_v123;
          _1757_v123 = _dafny.Set.fromElements((_this).f6);
          let _1758_v124;
          _1758_v124 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1736_v107.f7);
          let _1759_v125;
          _1759_v125 = _dafny.Map.Empty.slice().updateUnsafe(_1758_v124,(_1736_v107).f8);
          let _rhs251 = ((_this).f6) && (p1);
          let _rhs252 = ((_1621_v21).Union(_dafny.MultiSet.FromArray(_1730_v104))).IsSubsetOf(_1621_v21);
          let _rhs253 = !(((new BigNumber((_1740_v111).length)).multipliedBy((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_1757_v123).length))))).isEqualTo(new BigNumber((_1759_v125).length)));
          let _rhs254 = _1746_v115;
          let _lhs166 = _1736_v107;
          r2 = _rhs251;
          _1738_v109 = _rhs252;
          _lhs166.f7 = _rhs253;
          r0 = _rhs254;
          (_1736_v107).f7 = (_1736_v107.f7) && (_1746_v115);
        } else {
          let _1760_v126;
          _1760_v126 = _dafny.Map.Empty.slice().updateUnsafe((_module.D8.create_DC27(_1736_v107.f7, (_this).f6, _1599_v2, (_1736_v107).f8)).dtor_cf51,(_dafny.ZERO).minus((_1736_v107).f8));
          _1760_v126 = (_1760_v126).update((_1736_v107).f8, _1599_v2);
          let _index303 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_1735_v106).length));
          (_1735_v106)[_index303] = !((_1736_v107).f8).isEqualTo(_1599_v2);
          let _index304 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_1735_v106).length));
          (_1735_v106)[_index304] = !(_1736_v107.f7);
          r1 = (((_this).f6) ? ((_1736_v107).f8) : (_1599_v2));
          _1740_v111 = _dafny.Seq.Concat(_1740_v111, _dafny.Seq.UnicodeFromString("ngbnk"));
          let _index305 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_1735_v106).length));
          (_1735_v106)[_index305] = !(!(!(((false) ? (!(p1) || (_1746_v115)) : (((_1736_v107.f7) ? (p1) : (p1)))))));
        }
      }
      if ((_this).f6) {
        if (p1) {
          r0 = p1;
          _1599_v2 = _1599_v2;
          r0 = _module.__default.fm20(p1, globalState);
          let _1761_v127;
          _1761_v127 = _dafny.Seq.of(new BigNumber(-302));
          let _1762_v128;
          _1762_v128 = _dafny.Seq.of(_1761_v127);
          let _1763_v129;
          _1763_v129 = _dafny.Seq.of((_1762_v128)[_module.__default.safeIndex(_1599_v2, new BigNumber((_1762_v128).length))]);
          let _1764_v130;
          _1764_v130 = _dafny.Seq.of(_1761_v127, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-658)), ((_1765_v2) => function (_1766_i11) {
            return _1765_v2;
          })(_1599_v2)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(246)), ((_1767_v2) => function (_1768_i12) {
            return _1767_v2;
          })(_1599_v2)), _1761_v127, (_1762_v128)[_module.__default.safeIndex(_1599_v2, new BigNumber((_1762_v128).length))]);
          let _1769_v131;
          _1769_v131 = new _dafny.CodePoint('a'.codePointAt(0));
          let _1770_v132;
          _1770_v132 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_1769_v131)).length),_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-785)), ((_1771_v127, _1772_v2) => function (_1773_i13) {
            return _dafny.Seq.of(new BigNumber((_1771_v127).length), _1772_v2, new BigNumber(151));
          })(_1761_v127, _1599_v2)), _module.__default.safeIndex(_1599_v2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-785)), ((_1774_v127, _1775_v2) => function (_1776_i13) {
            return _dafny.Seq.of(new BigNumber((_1774_v127).length), _1775_v2, new BigNumber(151));
          })(_1761_v127, _1599_v2))).length)), _1761_v127), _module.__default.safeIndex(_1599_v2, new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-785)), ((_1777_v127, _1778_v2) => function (_1779_i13) {
            return _dafny.Seq.of(new BigNumber((_1777_v127).length), _1778_v2, new BigNumber(151));
          })(_1761_v127, _1599_v2)), _module.__default.safeIndex(_1599_v2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-785)), ((_1780_v127, _1781_v2) => function (_1782_i13) {
            return _dafny.Seq.of(new BigNumber((_1780_v127).length), _1781_v2, new BigNumber(151));
          })(_1761_v127, _1599_v2))).length)), _1761_v127)).length)), _dafny.Seq.of(_1599_v2)));
          let _1783_v133;
          _1783_v133 = _dafny.MultiSet.fromElements((_this).f6, p1);
          let _1784_v134;
          _1784_v134 = _dafny.Seq.UnicodeFromString("orstbaa");
          let _1785_v135;
          _1785_v135 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_1763_v129, _1764_v130), _dafny.Seq.of(_1761_v127, _1761_v127, _1761_v127, _1761_v127), _dafny.Seq.Concat((((_1770_v132).contains((_this).fm5(_1783_v133, globalState))) ? ((_1770_v132).get((_this).fm5(_1783_v133, globalState))) : (_dafny.Seq.update(_dafny.Seq.of(_1761_v127), _module.__default.safeIndex(new BigNumber((_1784_v134).length), new BigNumber((_dafny.Seq.of(_1761_v127)).length)), _1761_v127))), _1762_v128), _1764_v130, _1763_v129);
          _1599_v2 = (((_1785_v135).contains(_1762_v128)) ? ((_1785_v135).get(_1762_v128)) : (_1599_v2));
          let _1786_v136;
          _1786_v136 = _dafny.Map.Empty.slice().updateUnsafe(_1599_v2,_dafny.Set.fromElements(_1769_v131));
          let _1787_v137;
          _1787_v137 = _dafny.Set.fromElements(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), _1769_v131, (_module.D3.create_DC10(p1, _1769_v131, (_this).fm6(_dafny.Seq.of(_1769_v131, _1769_v131), p1, globalState), (_this).f6, p1)).dtor_cf24);
          _1786_v136 = (_1786_v136).update((_1761_v127)[_module.__default.safeIndex(_1599_v2, new BigNumber((_1761_v127).length))], _1787_v137);
        } else {
          let _1788_v138;
          _1788_v138 = new _dafny.CodePoint('q'.codePointAt(0));
          let _1789_v139;
          _1789_v139 = _dafny.Seq.of(_1788_v138, _1788_v138, _1788_v138, _1788_v138, _1788_v138);
          let _1790_v140;
          _1790_v140 = _module.D2.create_DC5(_1789_v139);
          let _1791_v141;
          _1791_v141 = _dafny.Set.fromElements((_1790_v140).dtor_cf13, _dafny.Seq.UnicodeFromString("ev"), _dafny.Seq.UnicodeFromString("jfvgu"), _dafny.Seq.UnicodeFromString("al"), _1789_v139);
          let _1792_v142;
          _1792_v142 = _dafny.Seq.of(_1791_v141);
          let _rhs255 = p1;
          let _rhs256 = (((_this).fm6(_1789_v139, (_this).f6, globalState)) ? ((_1792_v142)[_module.__default.safeIndex(_module.__default.fm0(_1599_v2, globalState), new BigNumber((_1792_v142).length))]) : (_dafny.Set.fromElements(_1789_v139)));
          let _rhs257 = (_1597_v0)[_module.__default.safeIndex(_1599_v2, new BigNumber((_1597_v0).length))];
          let _rhs258 = _module.__default.fm0(_1599_v2, globalState);
          r0 = _rhs255;
          _1791_v141 = _rhs256;
          r0 = _rhs257;
          r1 = _rhs258;
          let _1793_v143;
          let _init55 = function (_1794_i14) {
            return (_this).f6;
          };
          let _nw246 = Array((new BigNumber(6)).toNumber());
          for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw246.length); _i0_55++) {
            _nw246[_i0_55] = _init55(new BigNumber(_i0_55));
          }
          _1793_v143 = _nw246;
          let _1795_v144;
          let _nw247 = new _module.C3();
          _nw247.__ctor(_1793_v143);
          _1795_v144 = _nw247;
          _1788_v138 = _1788_v138;
          let _1796_v145;
          let _nw248 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Set.Empty);
          _1796_v145 = _nw248;
          let _1797_v146;
          _1797_v146 = _dafny.Set.fromElements((_this).f6, p1, p1);
          let _index306 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_1796_v145).length));
          (_1796_v145)[_index306] = _1797_v146;
          let _index307 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_1796_v145).length));
          (_1796_v145)[_index307] = _1797_v146;
          let _nw249 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
          _1723_v97 = _nw249;
        }
        let _1798_v147;
        _1798_v147 = _module.D1.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-27)), ((_1799_p1, _1800_v2) => function (_1801_i15) {
  return _module.D0.create_DC1(_1799_p1, _1799_p1, _1800_v2);
})(p1, _1599_v2)));
        let _1802_v148;
        _1802_v148 = _dafny.Seq.of(_1798_v147, _1798_v147, _1798_v147);
        let _1803_v149;
        _1803_v149 = _dafny.Seq.UnicodeFromString("epdgsg");
        let _rhs259 = p0;
        let _rhs260 = new BigNumber((_dafny.MultiSet.FromArray(_1802_v148)).cardinality());
        let _rhs261 = _1599_v2;
        let _rhs262 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_module.__default.safeModuloInt(_1599_v2, (_dafny.ZERO).minus(_1599_v2)), ((p1) ? (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1599_v2,_1803_v149)).length)) : (_1599_v2))));
        _1723_v97 = _rhs259;
        _1599_v2 = _rhs260;
        r1 = _rhs261;
        r1 = _rhs262;
        let _1804_v150;
        let _nw250 = Array((new BigNumber(28)).toNumber()).fill(false);
        _1804_v150 = _nw250;
        (globalState).f1 = _1804_v150;
        let _1805_v151;
        _1805_v151 = _dafny.Set.fromElements((_this).f6);
        _1805_v151 = _1805_v151;
        _1597_v0 = _dafny.Seq.Concat(_1597_v0, _1597_v0);
      } else {
        let _1806_v152;
        _1806_v152 = new _dafny.CodePoint('j'.codePointAt(0));
        let _1807_v153;
        _1807_v153 = _dafny.Seq.of(_1806_v152);
        let _rhs263 = (_module.__default.safeDivisionInt(_1599_v2, _1599_v2)).minus(_1599_v2);
        let _rhs264 = new BigNumber((_dafny.Seq.Concat(_1807_v153, _dafny.Seq.of(_module.__default.fm22(_1807_v153, globalState)))).length);
        _1599_v2 = _rhs263;
        _1599_v2 = _rhs264;
        r0 = ((_1621_v21).Union(_1621_v21)).equals(_1621_v21);
        let _rhs265 = p0;
        let _rhs266 = (_this).f6;
        let _rhs267 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1807_v153, _1807_v153), _dafny.Seq.Concat(_1807_v153, _dafny.Seq.UnicodeFromString("eyatql")));
        _1723_v97 = _rhs265;
        r0 = _rhs266;
        _1807_v153 = _rhs267;
        r1 = (_1599_v2).plus((_1599_v2).minus(_1599_v2));
        let _1808_v154;
        _1808_v154 = _dafny.Seq.of(_1599_v2, _1599_v2, _1599_v2, _1599_v2);
        _1621_v21 = (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.update(_1808_v154, _module.__default.safeIndex(new BigNumber((_1807_v153).length), new BigNumber((_1808_v154).length)), _1599_v2), _1808_v154))).Union(_1621_v21);
      }
      r0 = ((_1621_v21).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(_1599_v2, _1599_v2)))).IsDisjointFrom(_1621_v21);
      r1 = _1599_v2;
      let _1809_v155;
      _1809_v155 = _module.D0.create_DC0(true);
      let _pat_let_tv106 = p1;
      let _pat_let_tv107 = p1;
      let _pat_let_tv108 = p1;
      let _pat_let_tv109 = p1;
      let _pat_let_tv110 = p1;
      let _pat_let_tv111 = _1599_v2;
      let _pat_let_tv112 = _1599_v2;
      let _pat_let_tv113 = _1599_v2;
      r2 = function (_source24) {
        if (_source24.is_DC1) {
          let _1810___mcc_h4 = (_source24).cf1;
          let _1811___mcc_h5 = (_source24).cf2;
          let _1812___mcc_h6 = (_source24).cf3;
          let _1813_cf3 = _1812___mcc_h6;
          let _1814_cf2 = _1811___mcc_h5;
          let _1815_cf1 = _1810___mcc_h4;
          return (_this).f6;
        } else if (_source24.is_DC2) {
          let _1816___mcc_h7 = (_source24).cf4;
          let _1817___mcc_h8 = (_source24).cf5;
          let _1818___mcc_h9 = (_source24).cf6;
          let _1819___mcc_h10 = (_source24).cf7;
          let _1820_cf7 = _1819___mcc_h10;
          let _1821_cf6 = _1818___mcc_h9;
          let _1822_cf5 = _1817___mcc_h8;
          let _1823_cf4 = _1816___mcc_h7;
          return (_dafny.MultiSet.fromElements(_module.D13.create_DC40(_dafny.MultiSet.fromElements((_this).f6, _pat_let_tv106, _1821_cf6, false)))).IsProperSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.of(_module.D13.create_DC40(_dafny.MultiSet.fromElements(_1822_cf5)), _module.D13.create_DC40(_dafny.MultiSet.fromElements((_this).f6, (_this).f6)), _module.D13.create_DC40(_dafny.MultiSet.fromElements(_pat_let_tv107)), _module.D13.create_DC40(_dafny.MultiSet.fromElements(_pat_let_tv108)), _module.D13.create_DC40(_dafny.MultiSet.fromElements(_1821_cf6, _pat_let_tv109, true, _pat_let_tv110)))));
        } else {
          let _1824___mcc_h11 = (_source24).cf0;
          let _1825_cf0 = _1824___mcc_h11;
          return !(_dafny.areEqual(_dafny.Seq.of(_pat_let_tv111, _pat_let_tv112), _dafny.Seq.of(_pat_let_tv113)));
        }
      }(_1809_v155);
      return [r0, r1, r2];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f2 = [];
      this._f3 = [];
      this._f4 = false;
      this._f5 = false;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f2() {
      let _this = this;
      return _this._f2;
    };
    set f2(value) {
      let _this = this;
      _this._f2 = value;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    __ctor(f5, f2, f3, f4) {
      let _this = this;
      (_this)._f5 = f5;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      return;
    }
    fm1(p0, p1, globalState) {
      let _this = this;
      if ((_this).f4) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(672),(_this).f5);
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(780),(_this).f4);
      }
    };
    fm2(globalState) {
      let _this = this;
      return new BigNumber(661);
    };
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("ns")).length), (((_this).f5) ? (new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(!(false))).length))))).cardinality())) : ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements((_this).f5, (_this).f5)).cardinality()), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f5,new BigNumber(431))).length))).length))).length)))));
    };
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      r0 = (_this).f5;
      let _1826_v0;
      _1826_v0 = _module.D0.create_DC2(new BigNumber(397), (_this).f4, false, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(783)), function (_1827_i1) {
  return new _dafny.CodePoint('i'.codePointAt(0));
})).length));
      let _1828_i0;
      _1828_i0 = _dafny.ZERO;
      L8: {
        while (((((_this).f4) ? (_1826_v0) : (_1826_v0))).dtor_cf6) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1828_i0)) {
              break L8;
            }
            _1828_i0 = (_1828_i0).plus(_dafny.ONE);
            if (_module.__default.fm4(new BigNumber(-585), ((false) ? (function () {
              let _coll57 = new _dafny.Set();
              for (const _compr_57 of _dafny.IntegerRange(new BigNumber(361), new BigNumber(82))) {
                let _1829_v1 = _compr_57;
                if (((new BigNumber(361)).isLessThanOrEqualTo(_1829_v1)) && ((_1829_v1).isLessThan(new BigNumber(82)))) {
                  _coll57.add(_module.__default.safeDivisionInt(_1829_v1, p0));
                }
              }
              return _coll57;
            }()) : (_dafny.Set.fromElements(p0, p0, p0))), globalState)) {
              let _1830_v2;
              _1830_v2 = _module.D0.create_DC0((_this).f5);
              let _1831_v3;
              _1831_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_1830_v2);
              _1831_v3 = _1831_v3;
              let _1832_v4;
              _1832_v4 = _dafny.Seq.UnicodeFromString("mw");
              let _1833_v5;
              _1833_v5 = _dafny.Set.fromElements(new BigNumber((_1832_v4).length));
              _1833_v5 = _1833_v5;
              let _1834_v6;
              _1834_v6 = _dafny.Seq.of((_this).f4);
              let _1835_v7;
              let _nw251 = new _module.C2();
              _nw251.__ctor((_this).f4, new BigNumber((_1834_v6).length), (_this).f3, (_this).f5, _this.f2);
              _1835_v7 = _nw251;
              let _1836_v8;
              let _nw252 = Array((new BigNumber(23)).toNumber()).fill([]);
              _1836_v8 = _nw252;
              let _1837_v9;
              let _nw253 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
              _1837_v9 = _nw253;
              let _1838_v10;
              _1838_v10 = _dafny.Seq.of(_1837_v9);
              let _1839_v11;
              _1839_v11 = (_1838_v10)[_module.__default.safeIndex(p0, new BigNumber((_1838_v10).length))];
              let _index308 = _module.__default.safeIndex(new BigNumber(533), new BigNumber((_1836_v8).length));
              (_1836_v8)[_index308] = _1839_v11;
              let _index309 = _module.__default.safeIndex(new BigNumber(533), new BigNumber((_1836_v8).length));
              (_1836_v8)[_index309] = _1837_v9;
              let _1840_v12;
              let _nw254 = new _module.C3();
              _nw254.__ctor(_this.f2);
              _1840_v12 = _nw254;
            } else {
              let _1841_v13;
              let _init56 = function (_1842_i2) {
                return (_1842_i2).plus(new BigNumber((_dafny.Seq.UnicodeFromString("gwl")).length));
              };
              let _nw255 = Array((new BigNumber(24)).toNumber());
              for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw255.length); _i0_56++) {
                _nw255[_i0_56] = _init56(new BigNumber(_i0_56));
              }
              _1841_v13 = _nw255;
              let _index310 = _module.__default.safeIndex(new BigNumber(197), new BigNumber((_1841_v13).length));
              (_1841_v13)[_index310] = p0;
              let _index311 = _module.__default.safeIndex(new BigNumber(197), new BigNumber((_1841_v13).length));
              (_1841_v13)[_index311] = p0;
              let _1843_v14;
              let _nw256 = new _module.C1();
              _nw256.__ctor((_this).f3, (_module.__default.fm29(globalState)).contains((_1841_v13)[_module.__default.safeIndex(new BigNumber(197), new BigNumber((_1841_v13).length))]), _this.f2);
              _1843_v14 = _nw256;
              _1843_v14 = _1843_v14;
              let _1844_v15;
              _1844_v15 = _dafny.Seq.UnicodeFromString("f");
              let _1845_v16;
              _1845_v16 = _dafny.MultiSet.fromElements(!((_this).f4), (_this).f4, (_this).f5);
              let _1846_v17;
              _1846_v17 = new _dafny.CodePoint('s'.codePointAt(0));
              let _1847_v18;
              let _nw257 = new _module.C2();
              _nw257.__ctor((_this).f4, new BigNumber(403), (_this).f3, _dafny.Seq.IsProperPrefixOf(_1844_v15, _dafny.Seq.update(_1844_v15, _module.__default.safeIndex(new BigNumber((_1845_v16).cardinality()), new BigNumber((_1844_v15).length)), _1846_v17)), _this.f2);
              _1847_v18 = _nw257;
              let _1848_v19;
              _1848_v19 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_1843_v14).fm2(globalState), globalState),p0);
              let _1849_v20;
              _1849_v20 = _dafny.MultiSet.fromElements((_1847_v18).f8);
              _1848_v19 = (_1848_v19).update(new BigNumber(((_1849_v20).Intersect(_1849_v20)).cardinality()), p0);
              let _arr55 = _this.f2;
              let _index312 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_this.f2).length));
              _arr55[_index312] = (_this).f5;
              let _1850_v21;
              _1850_v21 = _dafny.Seq.of(_module.__default.fm11((_1847_v18).f8, !((_this).f5), (_dafny.ZERO).minus((_1847_v18).f8), globalState), _1844_v15, _dafny.Seq.UnicodeFromString("ylns"), _1844_v15);
              let _arr56 = _this.f2;
              let _index313 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_this.f2).length));
              let _index314 = _module.__default.safeIndex(new BigNumber(197), new BigNumber((_1841_v13).length));
              let _rhs268 = (((_this).f4) ? (_1841_v13) : (_1841_v13));
              let _rhs269 = !(_1847_v18.f7);
              let _rhs270 = (_1841_v13)[_module.__default.safeIndex(new BigNumber(197), new BigNumber((_1841_v13).length))];
              let _rhs271 = _dafny.Seq.of(_module.__default.fm13((_this).f5, true, (_this).f4, globalState), _dafny.Seq.Concat(_1844_v15, _dafny.Seq.UnicodeFromString("lpisqxnf")), _1844_v15, _module.__default.fm11(p0, (_this).f4, (_1847_v18).f8, globalState));
              let _lhs167 = _this.f2;
              let _lhs168 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_this.f2).length));
              let _lhs169 = _1841_v13;
              let _lhs170 = _module.__default.safeIndex(new BigNumber(197), new BigNumber((_1841_v13).length));
              _1841_v13 = _rhs268;
              _lhs167[_lhs168] = _rhs269;
              _lhs169[_lhs170] = _rhs270;
              _1850_v21 = _rhs271;
            }
            let _1851_v22;
            _1851_v22 = _dafny.Seq.UnicodeFromString("acgo");
            let _1852_v23;
            _1852_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1851_v22,p2);
            let _1853_v24;
            _1853_v24 = _dafny.Set.fromElements((_this).f4);
            let _1854_v25;
            _1854_v25 = _dafny.Set.fromElements(new BigNumber((_1853_v24).length), p0);
            let _1855_v26;
            _1855_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1854_v25).length),_module.__default.fm20((_this).f5, globalState));
            let _1856_v27;
            _1856_v27 = _dafny.Seq.of(p2);
            let _1857_v28;
            _1857_v28 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_1856_v27)[_module.__default.safeIndex(p0, new BigNumber((_1856_v27).length))]);
            _1852_v23 = (_1852_v23).update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(482)), function (_1858_i3) {
              return new _dafny.CodePoint('a'.codePointAt(0));
            }), _module.__default.fm18((_this).f4, _1855_v26, globalState)), (((_1857_v28).contains((_this).f5)) ? ((_1857_v28).get((_this).f5)) : (p2)));
            r0 = (((_1855_v26).contains(p0)) ? ((_1855_v26).get(p0)) : ((_this).f4));
            r0 = (_this).f5;
          }
        }
      }
      let _1859_v29;
      _1859_v29 = new _dafny.CodePoint('q'.codePointAt(0));
      _1859_v29 = _1859_v29;
      let _1860_v30;
      let _init57 = ((_1861_p0) => function (_1862_i4) {
        return _module.__default.safeDivisionInt(_1862_i4, _1861_p0);
      })(p0);
      let _nw258 = Array((new BigNumber(16)).toNumber());
      for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw258.length); _i0_57++) {
        _nw258[_i0_57] = _init57(new BigNumber(_i0_57));
      }
      _1860_v30 = _nw258;
      let _index315 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length));
      (_1860_v30)[_index315] = p0;
      let _1863_v31;
      _1863_v31 = _dafny.MultiSet.fromElements(p0);
      let _index316 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length));
      (_1860_v30)[_index316] = (_module.__default.safeModuloInt(new BigNumber(326), p0)).plus(new BigNumber(((_1863_v31).Intersect(_1863_v31)).cardinality()));
      let _1864_v32;
      _1864_v32 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f5);
      let _arr57 = _this.f2;
      let _index317 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length));
      _arr57[_index317] = (p0).isLessThan(new BigNumber(65));
      let _1865_v33;
      let _nw259 = Array((new BigNumber(28)).toNumber());
      _1865_v33 = _nw259;
      let _1866_v34;
      let _nw260 = new _module.C1();
      _nw260.__ctor((_this).f3, !((_this).f4), _this.f2);
      _1866_v34 = _nw260;
      let _index318 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_1865_v33).length));
      (_1865_v33)[_index318] = _1866_v34;
      let _arr58 = _this.f2;
      let _index319 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length));
      let _index320 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_1865_v33).length));
      let _rhs272 = (((_this).f5) ? ((_1866_v34).f4) : (!((_1866_v34).f4) || (true)));
      let _rhs273 = _1864_v32;
      let _rhs274 = ((new BigNumber(440)).plus((_dafny.ZERO).minus(p0))).isLessThan((_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))]);
      let _rhs275 = _1866_v34;
      let _rhs276 = _1860_v30;
      let _lhs171 = _this.f2;
      let _lhs172 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length));
      let _lhs173 = _1865_v33;
      let _lhs174 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_1865_v33).length));
      r0 = _rhs272;
      _1864_v32 = _rhs273;
      _lhs171[_lhs172] = _rhs274;
      _lhs173[_lhs174] = _rhs275;
      _1860_v30 = _rhs276;
      if ((_1866_v34).f4) {
        let _index321 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length));
        (_1860_v30)[_index321] = p0;
        if ((((_1866_v34).f4) ? ((_this).f5) : ((_1866_v34).f4))) {
          let _1867_v35;
          _1867_v35 = _dafny.Seq.of(new BigNumber(229), new BigNumber(-848), p0);
          let _1868_v36;
          _1868_v36 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsProperPrefixOf(_1867_v35, _dafny.Seq.of(p0)),(((_1864_v32).contains((_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))])) ? ((_1864_v32).get((_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))])) : ((_1866_v34).f4)));
          _1868_v36 = (_1868_v36).update((p0).isEqualTo((_dafny.ZERO).minus((_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))])), (_1866_v34).f4);
          let _1869_v37;
          let _nw261 = new _module.C0();
          _nw261.__ctor();
          _1869_v37 = _nw261;
          let _1870_v38;
          _1870_v38 = _dafny.Seq.of(_1869_v37);
          let _1871_v39;
          let _nw262 = Array((new BigNumber(6)).toNumber());
          _nw262[(_dafny.ZERO).toNumber()] = _1869_v37;
          _nw262[(_dafny.ONE).toNumber()] = _1869_v37;
          _nw262[(new BigNumber(2)).toNumber()] = _1869_v37;
          _nw262[(new BigNumber(3)).toNumber()] = _1869_v37;
          _nw262[(new BigNumber(4)).toNumber()] = _1869_v37;
          _nw262[(new BigNumber(5)).toNumber()] = (_1870_v38)[_module.__default.safeIndex((_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))], new BigNumber((_1870_v38).length))];
          _1871_v39 = _nw262;
          _1871_v39 = _1871_v39;
          let _1872_v40;
          _1872_v40 = _dafny.Seq.UnicodeFromString("egcrltj");
          let _1873_v41;
          _1873_v41 = _dafny.Seq.of(!((_1866_v34).f4));
          let _1874_v42;
          _1874_v42 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_module.__default.fm32(new BigNumber((_1872_v40).length), (_this.f2)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length))], _1868_v36, globalState), _1873_v41));
          let _1875_v43;
          let _init58 = function (_1876_i5) {
            return _dafny.MultiSet.fromElements((_this).f4);
          };
          let _nw263 = Array((new BigNumber(6)).toNumber());
          for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw263.length); _i0_58++) {
            _nw263[_i0_58] = _init58(new BigNumber(_i0_58));
          }
          _1875_v43 = _nw263;
          let _1877_v44;
          _1877_v44 = _dafny.MultiSet.fromElements(!((_this).f4));
          let _index322 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_1875_v43).length));
          (_1875_v43)[_index322] = _1877_v44;
          let _1878_v45;
          let _init59 = function (_1879_i6) {
            return true;
          };
          let _nw264 = Array((new BigNumber(11)).toNumber());
          for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw264.length); _i0_59++) {
            _nw264[_i0_59] = _init59(new BigNumber(_i0_59));
          }
          _1878_v45 = _nw264;
          let _1880_v46;
          let _nw265 = new _module.C2();
          _nw265.__ctor(false, (_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))], (_1866_v34).f3, (_1866_v34).f4, _1878_v45);
          _1880_v46 = _nw265;
          let _1881_v47;
          _1881_v47 = _dafny.Set.fromElements(_1880_v46);
          let _1882_v48;
          _1882_v48 = _dafny.Map.Empty.slice().updateUnsafe((_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))],_1881_v47);
          let _1883_v49;
          let _nw266 = new _module.C4();
          _nw266.__ctor(_1880_v46.f7);
          _1883_v49 = _nw266;
          let _1884_v50;
          _1884_v50 = _dafny.Seq.of(_1883_v49, _1883_v49);
          let _1885_v51;
          _1885_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1880_v46,(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1884_v50).length),_1881_v47)).update(p0, _1881_v47));
          let _index323 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_1875_v43).length));
          let _index324 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length));
          let _rhs277 = _1874_v42;
          let _rhs278 = (_1877_v44).Difference(_dafny.MultiSet.fromElements(_1880_v46.f7));
          let _rhs279 = (_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))];
          let _rhs280 = (_1882_v48).Merge((((_1885_v51).contains(_1880_v46)) ? ((_1885_v51).get(_1880_v46)) : (_1882_v48)));
          let _lhs175 = _1875_v43;
          let _lhs176 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_1875_v43).length));
          let _lhs177 = _1860_v30;
          let _lhs178 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length));
          _1874_v42 = _rhs277;
          _lhs175[_lhs176] = _rhs278;
          _lhs177[_lhs178] = _rhs279;
          _1882_v48 = _rhs280;
          let _1886_v52;
          let _nw267 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.of());
          _1886_v52 = _nw267;
          let _1887_v53;
          _1887_v53 = _dafny.Seq.of(_dafny.Set.fromElements((_this).f4, (_1866_v34).f4));
          let _index325 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1886_v52).length));
          (_1886_v52)[_index325] = _1887_v53;
          let _1888_v54;
          _1888_v54 = _dafny.Seq.of(_1887_v53);
          let _index326 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1886_v52).length));
          (_1886_v52)[_index326] = (_1888_v54)[_module.__default.safeIndex((_1880_v46).f8, new BigNumber((_1888_v54).length))];
          let _1889_v55;
          _1889_v55 = _dafny.Set.fromElements((_1880_v46).f8);
          let _index327 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length));
          (_1860_v30)[_index327] = new BigNumber(((_1889_v55).Intersect(_1889_v55)).length);
        } else {
          let _1890_v56;
          let _nw268 = new _module.C4();
          _nw268.__ctor(false);
          _1890_v56 = _nw268;
          _1890_v56 = _1890_v56;
          let _index328 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length));
          (_1860_v30)[_index328] = (_this).fm2(globalState);
          r0 = ((_dafny.ZERO).minus((_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))])).isLessThan((_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))]);
          let _1891_v57;
          let _init60 = ((_1892_v34) => function (_1893_i7) {
            return (_1892_v34).f4;
          })(_1866_v34);
          let _nw269 = Array((new BigNumber(21)).toNumber());
          for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw269.length); _i0_60++) {
            _nw269[_i0_60] = _init60(new BigNumber(_i0_60));
          }
          _1891_v57 = _nw269;
          let _1894_v58;
          let _nw270 = new _module.C1();
          _nw270.__ctor((_1866_v34).f3, (_1866_v34).f4, _1891_v57);
          _1894_v58 = _nw270;
          let _1895_v59;
          _1895_v59 = _dafny.Set.fromElements(new BigNumber(-573));
          r0 = _module.__default.fm4(p0, _1895_v59, globalState);
        }
        let _index329 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length));
        (_1860_v30)[_index329] = (_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))];
        let _index330 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length));
        (_1860_v30)[_index330] = (_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))];
        let _1896_v60;
        _1896_v60 = _dafny.Seq.UnicodeFromString("aoltl");
        let _1897_v61;
        _1897_v61 = _dafny.Set.fromElements(p0, (_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))]);
        let _1898_v62;
        _1898_v62 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(960)), ((_1899_v29) => function (_1900_i8) {
          return _1899_v29;
        })(_1859_v29))).length), new BigNumber((_1896_v60).length)),_1897_v61);
        let _1901_v63;
        _1901_v63 = _dafny.Seq.of((_this.f2)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length))], false);
        let _1902_v64;
        _1902_v64 = _module.D12.create_DC37((_this).f5, _1901_v63, p0);
        let _1903_v65;
        _1903_v65 = _module.D3.create_DC10((_1866_v34).f4, _1859_v29, (_1866_v34).f4, (_this).f5, (_this).f4);
        _1898_v62 = (_1898_v62).update((_1902_v64).dtor_cf69, (_dafny.Set.fromElements((_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))], (_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))])).Intersect(_dafny.Set.fromElements((_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))], new BigNumber((_dafny.Seq.of((_1903_v65).dtor_cf27)).length))));
      } else {
        let _1904_v66;
        _1904_v66 = _dafny.Seq.UnicodeFromString("dvygmlw");
        let _1905_v67;
        _1905_v67 = _module.D10.create_DC33((_1866_v34).f4);
        let _1906_v68;
        _1906_v68 = _dafny.Seq.of(_1905_v67);
        let _1907_v69;
        _1907_v69 = _module.D12.create_DC36(_dafny.Map.Empty.slice().updateUnsafe(_1904_v66,_1906_v68));
        let _1908_v70;
        _1908_v70 = _module.D12.create_DC39(_1907_v69);
        let _pat_let_tv114 = _1907_v69;
        let _source25 = function (_pat_let47_0) {
          return function (_1909_dt__update__tmp_h0) {
            return function (_pat_let48_0) {
              return function (_1910_dt__update_hcf74_h0) {
                return _module.D12.create_DC39(_1910_dt__update_hcf74_h0);
              }(_pat_let48_0);
            }(_pat_let_tv114);
          }(_pat_let47_0);
        }(_1908_v70);
        if (_source25.is_DC37) {
          let _1911___mcc_h0 = (_source25).cf67;
          let _1912___mcc_h1 = (_source25).cf68;
          let _1913___mcc_h2 = (_source25).cf69;
          let _1914_cf69 = _1913___mcc_h2;
          let _1915_cf68 = _1912___mcc_h1;
          let _1916_cf67 = _1911___mcc_h0;
          _1916_cf67 = (_this.f2)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length))];
          let _1917_v71;
          _1917_v71 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_dafny.Seq.UnicodeFromString("dj"));
          let _1918_v72;
          _1918_v72 = _dafny.Set.fromElements(_1914_cf69, (_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))], (new BigNumber((_1917_v71).length)).minus(new BigNumber(-994)), (((_1866_v34).f4) ? (_1914_cf69) : ((_dafny.ZERO).minus(_1914_cf69))), p0);
          _1918_v72 = ((_dafny.Set.fromElements(p0, p0)).Difference(function () {
            let _coll58 = new _dafny.Set();
            for (const _compr_58 of (_1863_v31).Elements) {
              let _1919_v73 = _compr_58;
              if ((_1863_v31).contains(_1919_v73)) {
                _coll58.add((_1919_v73).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(93)), function (_1920_i9) {
                  return new _dafny.CodePoint('k'.codePointAt(0));
                })).length)));
              }
            }
            return _coll58;
          }())).Union(_1918_v72);
          let _1921_v74;
          _1921_v74 = _module.D2.create_DC6(new _dafny.CodePoint('a'.codePointAt(0)), _1914_cf69);
          let _index331 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length));
          (_1860_v30)[_index331] = (_1921_v74).dtor_cf15;
          let _1922_v75;
          let _nw271 = new _module.C0();
          _nw271.__ctor();
          _1922_v75 = _nw271;
          let _1923_v76;
          _1923_v76 = _dafny.MultiSet.fromElements((_this).f5);
          let _1924_v77;
          _1924_v77 = _dafny.Seq.of(_1904_v66, _1904_v66);
          let _1925_v78;
          _1925_v78 = _module.D13.create_DC42(_1915_cf68, _1923_v76, _1924_v77, true, _dafny.Seq.Create(_module.__default.abs(new BigNumber(693)), ((_1926_cf69) => function (_1927_i10) {
  return _1926_cf69;
})(_1914_cf69)));
          let _1928_v79;
          _1928_v79 = _dafny.Map.Empty.slice().updateUnsafe(_1859_v29,(_1866_v34).f4);
          let _1929_v80;
          _1929_v80 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,new BigNumber((_1928_v79).length));
          let _pat_let_tv115 = _1866_v34;
          let _arr59 = _this.f2;
          let _index332 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length));
          let _arr60 = _this.f2;
          let _index333 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length));
          let _rhs281 = _module.__default.fm20(_1916_cf67, globalState);
          let _rhs282 = _dafny.Seq.Concat((_1925_v78).dtor_cf77, _1915_cf68);
          let _rhs283 = new BigNumber(((_1929_v80).Merge(_1929_v80)).length);
          let _rhs284 = _1922_v75;
          let _rhs285 = !((function (_pat_let49_0) {
            return function (_1930_dt__update__tmp_h1) {
              return function (_pat_let50_0) {
                return function (_1931_dt__update_hcf61_h0) {
                  return _module.D10.create_DC33(_1931_dt__update_hcf61_h0);
                }(_pat_let50_0);
              }((_pat_let_tv115).f4);
            }(_pat_let49_0);
          }(_1905_v67)).dtor_cf61) || (_1916_cf67);
          let _lhs179 = _this.f2;
          let _lhs180 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length));
          let _lhs181 = _this.f2;
          let _lhs182 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length));
          _lhs179[_lhs180] = _rhs281;
          _1915_cf68 = _rhs282;
          _1914_cf69 = _rhs283;
          _1922_v75 = _rhs284;
          _lhs181[_lhs182] = _rhs285;
        } else if (_source25.is_DC38) {
          let _1932___mcc_h3 = (_source25).cf70;
          let _1933___mcc_h4 = (_source25).cf71;
          let _1934___mcc_h5 = (_source25).cf72;
          let _1935___mcc_h6 = (_source25).cf73;
          let _1936_cf73 = _1935___mcc_h6;
          let _1937_cf72 = _1934___mcc_h5;
          let _1938_cf71 = _1933___mcc_h4;
          let _1939_cf70 = _1932___mcc_h3;
          let _1940_v81;
          let _nw272 = Array((new BigNumber(26)).toNumber()).fill([]);
          _1940_v81 = _nw272;
          let _1941_v82;
          _1941_v82 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_1940_v81);
          let _1942_v83;
          _1942_v83 = _module.D15.create_DC47(_1940_v81);
          let _pat_let_tv116 = _1940_v81;
          let _1943_v84;
          let _nw273 = Array((new BigNumber(29)).toNumber());
          _nw273[(_dafny.ZERO).toNumber()] = _1940_v81;
          _nw273[(_dafny.ONE).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(2)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(3)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(4)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(5)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(6)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(7)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(8)).toNumber()] = (((_1941_v82).contains((_1866_v34).f4)) ? ((_1941_v82).get((_1866_v34).f4)) : (_1940_v81));
          _nw273[(new BigNumber(9)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(10)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(11)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(12)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(13)).toNumber()] = ((_1937_cf72.f7) ? (_1940_v81) : (_1940_v81));
          _nw273[(new BigNumber(14)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(15)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(16)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(17)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(18)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(19)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(20)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(21)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(22)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(23)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(24)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(25)).toNumber()] = (function (_pat_let51_0) {
            return function (_1944_dt__update__tmp_h2) {
              return function (_pat_let52_0) {
                return function (_1945_dt__update_hcf86_h0) {
                  return _module.D15.create_DC47(_1945_dt__update_hcf86_h0);
                }(_pat_let52_0);
              }(_pat_let_tv116);
            }(_pat_let51_0);
          }(_1942_v83)).dtor_cf86;
          _nw273[(new BigNumber(26)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(27)).toNumber()] = _1940_v81;
          _nw273[(new BigNumber(28)).toNumber()] = _1940_v81;
          _1943_v84 = _nw273;
          let _index334 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_1943_v84).length));
          (_1943_v84)[_index334] = _1940_v81;
          let _index335 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_1943_v84).length));
          (_1943_v84)[_index335] = _1940_v81;
          let _1946_v85;
          _1946_v85 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),(_this.f2)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length))]);
          let _1947_v86;
          _1947_v86 = _module.D3.create_DC9(_1946_v85);
          let _1948_v88;
          _1948_v88 = _dafny.MultiSet.fromElements(((_1947_v86).dtor_cf22).Merge(_1946_v85), _dafny.Map.Empty.slice().updateUnsafe(_1859_v29,true), _1946_v85, (function () {
            let _coll59 = new _dafny.Map();
            for (const _compr_59 of (_1946_v85).Keys.Elements) {
              let _1949_v87 = _compr_59;
              if ((_1946_v85).contains(_1949_v87)) {
                _coll59.push([_1949_v87,(_1866_v34).f4]);
              }
            }
            return _coll59;
          }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1859_v29,(_1866_v34).f4)));
          let _1950_v89;
          _1950_v89 = _module.D16.create_DC50(_1948_v88);
          let _arr61 = _this.f2;
          let _index336 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length));
          let _rhs286 = _1937_cf72.f7;
          let _rhs287 = (_1950_v89).dtor_cf93;
          let _rhs288 = (_1866_v34).f4;
          let _rhs289 = !(_module.__default.fm20((_1866_v34).f4, globalState));
          let _rhs290 = _1860_v30;
          let _lhs183 = _this.f2;
          let _lhs184 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length));
          let _lhs185 = _1937_cf72;
          let _lhs186 = _1937_cf72;
          _lhs183[_lhs184] = _rhs286;
          _1948_v88 = _rhs287;
          _lhs185.f7 = _rhs288;
          _lhs186.f7 = _rhs289;
          _1860_v30 = _rhs290;
          _1939_cf70 = (_1939_cf70).Intersect(_1939_cf70);
          let _1951_v90;
          _1951_v90 = _module.D7.create_DC23(_1938_cf71);
          _1936_cf73 = (_1951_v90).dtor_cf42;
        } else if (_source25.is_DC36) {
          let _1952___mcc_h7 = (_source25).cf66;
          let _1953_cf66 = _1952___mcc_h7;
          r0 = !((_this.f2)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length))]);
          let _1954_v91;
          _1954_v91 = _dafny.MultiSet.fromElements(!((_this).f5), !((_1866_v34).f4), (_this).f5);
          let _1955_v92;
          _1955_v92 = _dafny.Map.Empty.slice().updateUnsafe((_1866_v34).f4,new BigNumber((_1904_v66).length));
          let _index337 = _module.__default.safeIndex(new BigNumber(604), new BigNumber(((_this).f3).length));
          ((_this).f3)[_index337] = (_1955_v92).Merge(_1955_v92);
          let _1956_v93;
          _1956_v93 = _dafny.Seq.of(false, (_1866_v34).f4, (_this.f2)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length))], (_1866_v34).f4, (_this.f2)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length))]);
          let _index338 = _module.__default.safeIndex(new BigNumber(604), new BigNumber(((_this).f3).length));
          let _rhs291 = ((_1954_v91).Intersect(_dafny.MultiSet.FromArray(_1956_v93))).Union(_dafny.MultiSet.fromElements((_1866_v34).f4));
          let _rhs292 = (_dafny.Map.Empty.slice().updateUnsafe((_this.f2)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length))],p0)).update((_this.f2)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length))], new BigNumber((_1863_v31).cardinality()));
          let _lhs187 = (_this).f3;
          let _lhs188 = _module.__default.safeIndex(new BigNumber(604), new BigNumber(((_this).f3).length));
          _1954_v91 = _rhs291;
          _lhs187[_lhs188] = _rhs292;
          let _1957_v94;
          _1957_v94 = _dafny.Map.Empty.slice().updateUnsafe((_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))],(p0).plus((_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))]));
          _1957_v94 = _1957_v94;
          let _1958_v95;
          _1958_v95 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_1956_v93);
          _1958_v95 = (_1958_v95).update(true, _1956_v93);
        } else {
          let _1959___mcc_h8 = (_source25).cf74;
          let _1960_cf74 = _1959___mcc_h8;
          let _1961_v96;
          let _init61 = function (_1962_i11) {
            return (_this.f2)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length))];
          };
          let _nw274 = Array((new BigNumber(5)).toNumber());
          for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw274.length); _i0_61++) {
            _nw274[_i0_61] = _init61(new BigNumber(_i0_61));
          }
          _1961_v96 = _nw274;
          let _1963_v97;
          _1963_v97 = _dafny.Map.Empty.slice().updateUnsafe(_1961_v96,_1961_v96);
          let _1964_v98;
          let _nw275 = new _module.C2();
          _nw275.__ctor(!((_1866_v34).f4), (_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))], (_1866_v34).f3, (_this).f5, (((_1963_v97).contains(_1961_v96)) ? ((_1963_v97).get(_1961_v96)) : (_1961_v96)));
          _1964_v98 = _nw275;
          _1859_v29 = (((_this.f2)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length))]) ? (_1859_v29) : (_1859_v29));
          let _1965_v99;
          _1965_v99 = _dafny.Set.fromElements((_1964_v98).f8, (_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))]);
          let _1966_v100;
          _1966_v100 = _module.D15.create_DC48(_1964_v98.f7, _1961_v96, (_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))], _module.__default.fm4(p0, _1965_v99, globalState), _1904_v66);
          let _1967_v101;
          let _nw276 = new _module.C3();
          _nw276.__ctor((_1966_v100).dtor_cf88);
          _1967_v101 = _nw276;
          let _arr62 = _this.f2;
          let _index339 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length));
          _arr62[_index339] = (_this).f4;
        }
        let _index340 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length));
        (_1860_v30)[_index340] = p0;
        let _index341 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length));
        let _rhs293 = _1859_v29;
        let _rhs294 = (_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))];
        let _lhs189 = _1860_v30;
        let _lhs190 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length));
        _1859_v29 = _rhs293;
        _lhs189[_lhs190] = _rhs294;
        let _1968_v102;
        _1968_v102 = _dafny.Map.Empty.slice().updateUnsafe(!((_this.f2)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_this.f2).length))]),(_this).f4);
        let _index342 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length));
        (_1860_v30)[_index342] = new BigNumber((_1968_v102).length);
        let _1969_v103;
        _1969_v103 = _dafny.Seq.of(_1968_v102, _1968_v102, _1968_v102, _module.__default.fm36((_1866_v34).f4, globalState));
        let _1970_v104;
        _1970_v104 = _module.D7.create_DC23((_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))]);
        let _pat_let_tv117 = _1860_v30;
        let _pat_let_tv118 = _1860_v30;
        _1968_v102 = (_1969_v103)[_module.__default.safeIndex((function (_pat_let53_0) {
          return function (_1971_dt__update__tmp_h3) {
            return function (_pat_let54_0) {
              return function (_1972_dt__update_hcf42_h0) {
                return _module.D7.create_DC23(_1972_dt__update_hcf42_h0);
              }(_pat_let54_0);
            }((_pat_let_tv118)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_pat_let_tv117).length))]);
          }(_pat_let53_0);
        }(_1970_v104)).dtor_cf42, new BigNumber((_1969_v103).length))];
      }
      let _1973_v105;
      let _nw277 = Array((new BigNumber(9)).toNumber());
      _1973_v105 = _nw277;
      let _1974_v106;
      _1974_v106 = _dafny.Seq.of(_1973_v105);
      r0 = _dafny.Seq.contains(_dafny.Seq.update(_dafny.Seq.Concat(_1974_v106, _dafny.Seq.of(_1973_v105, _1973_v105)), _module.__default.safeIndex((_1860_v30)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1860_v30).length))], new BigNumber((_dafny.Seq.Concat(_1974_v106, _dafny.Seq.of(_1973_v105, _1973_v105))).length)), _1973_v105), _1973_v105);
      return r0;
    }
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.MultiSet.Empty;
      let r1 = false;
      let _1975_v0;
      let _nw278 = new _module.C2();
      _nw278.__ctor(true, (_dafny.ZERO).minus(p0), (_this).f3, (_this).f4, _this.f2);
      _1975_v0 = _nw278;
      _1975_v0 = _1975_v0;
      r1 = (_this).f5;
      let _1976_v1;
      _1976_v1 = _dafny.Map.Empty.slice().updateUnsafe(((!((_this).f4)) ? (_1975_v0.f2) : (_1975_v0.f2)),((_this).f5) === (true));
      _1976_v1 = _dafny.Map.Empty.slice().updateUnsafe(_1975_v0.f2,_module.__default.fm20((_this).f5, globalState));
      if ((((p0).isEqualTo(p0)) ? ((_this).f5) : ((_this).f4))) {
        r1 = (p0).isLessThanOrEqualTo(new BigNumber(407));
        let _1977_v2;
        _1977_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,(_this).f4);
        let _1978_v3;
        let _nw279 = new _module.C1();
        _nw279.__ctor((_this).f3, _module.__default.fm4(p0, _dafny.Set.fromElements((_dafny.ZERO).minus(p0), p0, new BigNumber((_1977_v2).length)), globalState), _1975_v0.f2);
        _1978_v3 = _nw279;
        let _1979_v4;
        _1979_v4 = new BigNumber(819);
        _1979_v4 = _1979_v4;
        _1979_v4 = p0;
        let _1980_v5;
        let _nw280 = new _module.C4();
        _nw280.__ctor((_this).f5);
        _1980_v5 = _nw280;
        _1980_v5 = _1980_v5;
      } else {
        let _1981_v6;
        let _nw281 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
        _1981_v6 = _nw281;
        let _index343 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_1981_v6).length));
        (_1981_v6)[_index343] = p0;
        let _index344 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_1981_v6).length));
        (_1981_v6)[_index344] = _module.__default.safeModuloInt(p0, _module.__default.safeModuloInt(new BigNumber((_module.__default.fm37(p0, true, p0, globalState)).length), p0));
        let _arr63 = _1975_v0.f2;
        let _index345 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_1975_v0.f2).length));
        _arr63[_index345] = false;
        let _1982_v7;
        _1982_v7 = _dafny.MultiSet.fromElements((_1981_v6)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((_1981_v6).length))]);
        let _arr64 = _1975_v0.f2;
        let _index346 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_1975_v0.f2).length));
        _arr64[_index346] = (_dafny.MultiSet.fromElements((_1981_v6)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((_1981_v6).length))], p0, p0, (_1981_v6)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((_1981_v6).length))])).IsProperSubsetOf(_1982_v7);
        let _1983_v8;
        _1983_v8 = _1981_v6;
        _1981_v6 = (_1983_v8);
        let _1984_v9;
        _1984_v9 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1975_v0.f2)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_1975_v0.f2).length))]);
        let _1985_v10;
        _1985_v10 = _dafny.MultiSet.fromElements((_this).f5);
        let _1986_v11;
        _1986_v11 = _module.D1.create_DC4((_1984_v9).update((_dafny.ZERO).minus((_1981_v6)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((_1981_v6).length))]), (_this).f4), (_1975_v0.f2)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_1975_v0.f2).length))], (_1981_v6)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((_1981_v6).length))], (_dafny.MultiSet.fromElements((_this).f5, (_1975_v0.f2)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_1975_v0.f2).length))])).IsDisjointFrom(_1985_v10));
        let _source26 = _1986_v11;
        if (_source26.is_DC4) {
          let _1987___mcc_h0 = (_source26).cf9;
          let _1988___mcc_h1 = (_source26).cf10;
          let _1989___mcc_h2 = (_source26).cf11;
          let _1990___mcc_h3 = (_source26).cf12;
          let _1991_cf12 = _1990___mcc_h3;
          let _1992_cf11 = _1989___mcc_h2;
          let _1993_cf10 = _1988___mcc_h1;
          let _1994_cf9 = _1987___mcc_h0;
          let _arr65 = _1975_v0.f2;
          let _index347 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_1975_v0.f2).length));
          _arr65[_index347] = (_1984_v9).equals(_1994_cf9);
          let _1995_v12;
          _1995_v12 = _dafny.Seq.UnicodeFromString("bns");
          _1992_cf11 = (((false) ? ((_1975_v0).fm2(globalState)) : (new BigNumber((_1995_v12).length)))).multipliedBy(((_1981_v6)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((_1981_v6).length))]).multipliedBy(new BigNumber(((_1982_v7).update(p0, _module.__default.abs(_1992_cf11))).cardinality())));
          let _1996_v13;
          let _nw282 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Map.Empty);
          _1996_v13 = _nw282;
          let _1997_v14;
          _1997_v14 = _dafny.Seq.of(_1991_cf12);
          let _1998_v15;
          _1998_v15 = _dafny.Seq.of((_1997_v14)[_module.__default.safeIndex(new BigNumber((_1995_v12).length), new BigNumber((_1997_v14).length))]);
          let _1999_v16;
          _1999_v16 = _module.D4.create_DC11(_1997_v14);
          let _2000_v17;
          _2000_v17 = _module.D5.create_DC18(_module.D5.create_DC15(_1999_v16, (_this).f5));
          let _2001_v18;
          _2001_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_2000_v17);
          let _index348 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_1996_v13).length));
          (_1996_v13)[_index348] = _2001_v18;
          let _index349 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_1996_v13).length));
          (_1996_v13)[_index349] = (_2001_v18).Merge(_2001_v18);
          let _2002_v20;
          _2002_v20 = _dafny.Set.fromElements((_1981_v6)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((_1981_v6).length))]);
          let _2003_v21;
          _2003_v21 = _dafny.Seq.of(function () {
            let _coll60 = new _dafny.Set();
            for (const _compr_60 of _dafny.IntegerRange(new BigNumber(569), new BigNumber(153))) {
              let _2004_v19 = _compr_60;
              if (((new BigNumber(569)).isLessThanOrEqualTo(_2004_v19)) && ((_2004_v19).isLessThan(new BigNumber(153)))) {
                _coll60.add((_2004_v19).multipliedBy(p0));
              }
            }
            return _coll60;
          }(), _2002_v20);
          let _2005_v22;
          _2005_v22 = _module.D0.create_DC2(new BigNumber((_1997_v14).length), _1991_cf12, (_this).f5, new BigNumber(((_2003_v21)[_module.__default.safeIndex((_1981_v6)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((_1981_v6).length))], new BigNumber((_2003_v21).length))]).length));
          let _2006_v23;
          let _nw283 = new _module.C2();
          _nw283.__ctor((_2005_v22).dtor_cf6, ((_dafny.ZERO).minus(p0)).plus((_1981_v6)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((_1981_v6).length))]), (_this).f3, !((_1975_v0.f2)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_1975_v0.f2).length))]), _this.f2);
          _2006_v23 = _nw283;
        } else {
          let _2007___mcc_h4 = (_source26).cf8;
          let _2008_cf8 = _2007___mcc_h4;
          let _2009_v24;
          let _nw284 = new _module.C0();
          _nw284.__ctor();
          _2009_v24 = _nw284;
          let _2010_v25;
          _2010_v25 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f5),!((_1975_v0.f2)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_1975_v0.f2).length))]));
          let _2011_v26;
          _2011_v26 = _dafny.Seq.UnicodeFromString("qahpyyt");
          let _2012_v27;
          _2012_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2010_v25).length),new BigNumber((_2011_v26).length));
          let _2013_v28;
          _2013_v28 = _dafny.Set.fromElements((_1981_v6)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((_1981_v6).length))]);
          let _arr66 = _1975_v0.f2;
          let _index350 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_1975_v0.f2).length));
          let _arr67 = _1975_v0.f2;
          let _index351 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_1975_v0.f2).length));
          let _arr68 = _1975_v0.f2;
          let _index352 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_1975_v0.f2).length));
          let _rhs295 = _module.__default.fm4(p0, (_2013_v28).Intersect(_2013_v28), globalState);
          let _rhs296 = false;
          let _rhs297 = (_2009_v24).fm7((_this).f4, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1975_v0.f2)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_1975_v0.f2).length))],(_1975_v0.f2)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_1975_v0.f2).length))])).length), (false) === ((_this).f5), p0, globalState);
          let _rhs298 = (_2012_v27).Merge(_2012_v27);
          let _lhs191 = _1975_v0.f2;
          let _lhs192 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_1975_v0.f2).length));
          let _lhs193 = _1975_v0.f2;
          let _lhs194 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_1975_v0.f2).length));
          let _lhs195 = _1975_v0.f2;
          let _lhs196 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_1975_v0.f2).length));
          _lhs191[_lhs192] = _rhs295;
          _lhs193[_lhs194] = _rhs296;
          _lhs195[_lhs196] = _rhs297;
          _2012_v27 = _rhs298;
          let _index353 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_1981_v6).length));
          (_1981_v6)[_index353] = (_dafny.ZERO).minus(p0);
          let _arr69 = _1975_v0.f2;
          let _index354 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_1975_v0.f2).length));
          _arr69[_index354] = (_this).f5;
        }
        let _2014_v29;
        _2014_v29 = _dafny.Seq.UnicodeFromString("cemk");
        let _2015_v30;
        _2015_v30 = _dafny.Set.fromElements(p1);
        let _2016_v31;
        _2016_v31 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(670)), ((_2017_p0) => function (_2018_i0) {
          return _2017_p0;
        })(p0)),_2015_v30);
        let _2019_v32;
        _2019_v32 = _dafny.Seq.of((_1975_v0.f2)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_1975_v0.f2).length))], (_1975_v0.f2)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_1975_v0.f2).length))]);
        r1 = !_dafny.Seq.contains(_2019_v32, ((((_2016_v31).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(136)), ((_2022_p0) => function (_2023_i1) {
          return _2022_p0;
        })(p0)))) ? ((_2016_v31).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(136)), ((_2020_p0) => function (_2021_i1) {
          return _2020_p0;
        })(p0)))) : (_2015_v30))).IsSubsetOf(_dafny.Set.fromElements(p1, _module.__default.fm22(_2014_v29, globalState))));
      }
      let _hi10 = (_this).fm2(globalState);
      for (let _2024_i2 = p0; _2024_i2.isLessThan(_hi10); _2024_i2 = _2024_i2.plus(_dafny.ONE)) {
        r1 = (_2024_i2).isLessThan(_2024_i2);
        let _2025_v33;
        let _nw285 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _2025_v33 = _nw285;
        _2025_v33 = _2025_v33;
        let _2026_v34;
        let _nw286 = new _module.C4();
        _nw286.__ctor((_this).f5);
        _2026_v34 = _nw286;
        let _2027_v35;
        _2027_v35 = _dafny.MultiSet.fromElements((_this).f5);
        let _2028_v36;
        _2028_v36 = _module.D13.create_DC40((_2027_v35).Union(_dafny.MultiSet.fromElements((_this).f5)));
        _2028_v36 = _2028_v36;
      }
      let _2029_v37;
      _2029_v37 = new BigNumber(586);
      let _rhs299 = ((!(_module.__default.fm20((_this).f4, globalState))) ? (!(true) || (true)) : (_module.__default.fm20(!((_this).f5), globalState)));
      let _rhs300 = p0;
      r1 = _rhs299;
      _2029_v37 = _rhs300;
      let _2030_v38;
      _2030_v38 = _dafny.MultiSet.fromElements(_1975_v0.f2, _this.f2, _this.f2, _this.f2);
      r0 = _2030_v38;
      r1 = (_2029_v37).isLessThanOrEqualTo(_2029_v37);
      return [r0, r1];
    }
    get f5() {
      let _this = this;
      return _this._f5;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
