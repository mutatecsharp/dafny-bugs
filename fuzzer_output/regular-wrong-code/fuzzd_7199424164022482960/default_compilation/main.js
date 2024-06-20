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
      if (true) {
        return new BigNumber(772);
      } else {
        return (new BigNumber(914)).multipliedBy(new BigNumber(333));
      }
    };
    static fm1(p0, globalState) {
      return ((_module.D0.create_DC0(new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(918)), function (_0_i0) {
  return new BigNumber(137);
})).length)), _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(true)).length)))).length), _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("tq"),new BigNumber(-356)), _dafny.MultiSet.fromElements(new BigNumber(-192), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length)))).dtor_cf0).isLessThanOrEqualTo(((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-610))))).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length)));
    };
    static fm2(p0, globalState) {
      return _dafny.MultiSet.fromElements(new BigNumber(392));
    };
    static fm3(globalState) {
      return _module.D0.create_DC0(new BigNumber((_dafny.Seq.UnicodeFromString("m")).length), (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("xnjhhmdfk"),new BigNumber((_dafny.Seq.UnicodeFromString("tslm")).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("lhmcdhe"),new BigNumber(798))), (_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())), new BigNumber((_dafny.Seq.UnicodeFromString("xoksdycy")).length))).Union(_dafny.MultiSet.fromElements(new BigNumber(298))));
    };
    static fm4(p0, p1, p2, p3, globalState) {
      return (_dafny.Set.fromElements(!(!(false)), false, false, true)).Difference((_dafny.Set.fromElements(false, true, false)).Difference(_dafny.Set.fromElements(!(true), false)));
    };
    static fm5(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('l'.codePointAt(0))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(780)), function (_1_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      })), _dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0))));
    };
    static fm6(p0, globalState) {
      return new _dafny.CodePoint('c'.codePointAt(0));
    };
    static fm7(p0, p1, p2, globalState) {
      return ((_dafny.MultiSet.fromElements(true, false)).Difference(_dafny.MultiSet.fromElements(false, true, false))).Union((_dafny.MultiSet.FromArray(_dafny.Seq.of(false, true, true, !(true)))).Intersect(_dafny.MultiSet.fromElements(true)));
    };
    static fm8(globalState) {
      return (((false) ? (_module.D4.create_DC11(function () {
  let _coll0 = new _dafny.Map();
  for (const _compr_0 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(649)), function (_2_i0) {
    return (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(204), new BigNumber(26))).length));
  })).Elements) {
    let _3_v0 = _compr_0;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(649)), function (_2_i0) {
      return (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(204), new BigNumber(26))).length));
    }), _3_v0)) {
      _coll0.push([_module.__default.safeDivisionInt(_3_v0, new BigNumber(308)),true]);
    }
  }
  return _coll0;
}())) : (_module.D4.create_DC11(function () {
  let _coll1 = new _dafny.Map();
  for (const _compr_1 of (function () {
    let _coll2 = new _dafny.Map();
    for (const _compr_2 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("t")).length),new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false)).length))).length))).Keys.Elements) {
      let _4_v2 = _compr_2;
      if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("t")).length),new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false)).length))).length))).contains(_4_v2)) {
        _coll2.push([(_4_v2).plus(new BigNumber((_dafny.Seq.UnicodeFromString("qc")).length)),new BigNumber((_dafny.Seq.UnicodeFromString("mrjysrs")).length)]);
      }
    }
    return _coll2;
  }()).Keys.Elements) {
    let _5_v1 = _compr_1;
    if ((function () {
      let _coll3 = new _dafny.Map();
      for (const _compr_3 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("t")).length),new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false)).length))).length))).Keys.Elements) {
        let _4_v2 = _compr_3;
        if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("t")).length),new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false)).length))).length))).contains(_4_v2)) {
          _coll3.push([(_4_v2).plus(new BigNumber((_dafny.Seq.UnicodeFromString("qc")).length)),new BigNumber((_dafny.Seq.UnicodeFromString("mrjysrs")).length)]);
        }
      }
      return _coll3;
    }()).contains(_5_v1)) {
      _coll1.push([(_5_v1).minus(new BigNumber(432)),true]);
    }
  }
  return _coll1;
}())))).dtor_cf19;
    };
    static m0(p0, p1, globalState) {
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Seq.of();
      let r2 = _dafny.Set.Empty;
      let r3 = false;
      let _6_v0;
      _6_v0 = _dafny.Seq.UnicodeFromString("hkx");
      let _7_v1;
      _7_v1 = _dafny.Seq.of(_6_v0);
      let _8_v2;
      _8_v2 = true;
      let _9_v3;
      _9_v3 = _dafny.Map.Empty.slice().updateUnsafe(_8_v2,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(902)), function (_10_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })).length));
      let _11_v4;
      _11_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(p1)).length),_8_v2);
      let _12_v5;
      _12_v5 = _dafny.Seq.of(_8_v2, _8_v2);
      let _13_v6;
      _13_v6 = _dafny.Map.Empty.slice().updateUnsafe(_6_v0,p1);
      let _14_v7;
      _14_v7 = new _dafny.CodePoint('i'.codePointAt(0));
      let _15_v8;
      _15_v8 = _dafny.Map.Empty.slice().updateUnsafe(_14_v7,true);
      let _16_v9;
      _16_v9 = _dafny.MultiSet.fromElements(new BigNumber((_15_v8).length), new BigNumber(-545), p1, p1);
      let _17_v10;
      _17_v10 = _module.D0.create_DC0(p1, _13_v6, _16_v9);
      let _18_v11;
      _18_v11 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_17_v10, globalState),_6_v0);
      let _19_v12;
      _19_v12 = _dafny.Map.Empty.slice().updateUnsafe((((_11_v4).contains(new BigNumber((_12_v5).length))) ? ((_11_v4).get(new BigNumber((_12_v5).length))) : (_8_v2)),(((_18_v11).contains(_module.__default.fm0(_17_v10, globalState))) ? ((_18_v11).get(_module.__default.fm0(_17_v10, globalState))) : (_6_v0)));
      let _20_v13;
      _20_v13 = _module.D1.create_DC4(_6_v0, _8_v2, _7_v1, _6_v0, _14_v7);
      let _21_v14;
      let _nw0 = Array((new BigNumber(28)).toNumber());
      _nw0[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat((_7_v1)[_module.__default.safeIndex(p1, new BigNumber((_7_v1).length))], _6_v0);
      _nw0[(_dafny.ONE).toNumber()] = _6_v0;
      _nw0[(new BigNumber(2)).toNumber()] = _module.__default.fm5(new BigNumber((_6_v0).length), p1, p1, new BigNumber((_9_v3).length), globalState);
      _nw0[(new BigNumber(3)).toNumber()] = _6_v0;
      _nw0[(new BigNumber(4)).toNumber()] = _6_v0;
      _nw0[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_6_v0, _6_v0);
      _nw0[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(713)), ((_22_v0) => function (_23_i1) {
        return (_22_v0)[_module.__default.safeIndex(new BigNumber(629), new BigNumber((_22_v0).length))];
      })(_6_v0));
      _nw0[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("hrekd");
      _nw0[(new BigNumber(8)).toNumber()] = _6_v0;
      _nw0[(new BigNumber(9)).toNumber()] = _6_v0;
      _nw0[(new BigNumber(10)).toNumber()] = _6_v0;
      _nw0[(new BigNumber(11)).toNumber()] = _6_v0;
      _nw0[(new BigNumber(12)).toNumber()] = (((_19_v12).contains(_8_v2)) ? ((_19_v12).get(_8_v2)) : (_dafny.Seq.update(_6_v0, _module.__default.safeIndex(new BigNumber(385), new BigNumber((_6_v0).length)), (_20_v13).dtor_cf10)));
      _nw0[(new BigNumber(13)).toNumber()] = _6_v0;
      _nw0[(new BigNumber(14)).toNumber()] = _6_v0;
      _nw0[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_6_v0, _6_v0);
      _nw0[(new BigNumber(16)).toNumber()] = _6_v0;
      _nw0[(new BigNumber(17)).toNumber()] = _6_v0;
      _nw0[(new BigNumber(18)).toNumber()] = _6_v0;
      _nw0[(new BigNumber(19)).toNumber()] = _6_v0;
      _nw0[(new BigNumber(20)).toNumber()] = _6_v0;
      _nw0[(new BigNumber(21)).toNumber()] = _6_v0;
      _nw0[(new BigNumber(22)).toNumber()] = _dafny.Seq.UnicodeFromString("mxsg");
      _nw0[(new BigNumber(23)).toNumber()] = _dafny.Seq.Concat(_6_v0, _6_v0);
      _nw0[(new BigNumber(24)).toNumber()] = _6_v0;
      _nw0[(new BigNumber(25)).toNumber()] = _6_v0;
      _nw0[(new BigNumber(26)).toNumber()] = _6_v0;
      _nw0[(new BigNumber(27)).toNumber()] = _6_v0;
      _21_v14 = _nw0;
      _21_v14 = _21_v14;
      let _24_v15;
      let _nw1 = new _module.C0();
      _nw1.__ctor((p1).plus(p1));
      _24_v15 = _nw1;
      let _25_v16;
      _25_v16 = _dafny.Set.fromElements(_8_v2, _8_v2);
      let _26_v17;
      _26_v17 = _module.D1.create_DC3(_25_v16);
      let _pat_let_tv0 = _6_v0;
      let _pat_let_tv1 = _6_v0;
      let _pat_let_tv2 = _8_v2;
      let _pat_let_tv3 = _8_v2;
      let _pat_let_tv4 = _8_v2;
      r3 = function (_source0) {
        if (_source0.is_DC3) {
          let _27___mcc_h0 = (_source0).cf5;
          let _28_cf5 = _27___mcc_h0;
          return _dafny.Seq.IsProperPrefixOf(_pat_let_tv0, _pat_let_tv1);
        } else if (_source0.is_DC4) {
          let _29___mcc_h1 = (_source0).cf6;
          let _30___mcc_h2 = (_source0).cf7;
          let _31___mcc_h3 = (_source0).cf8;
          let _32___mcc_h4 = (_source0).cf9;
          let _33___mcc_h5 = (_source0).cf10;
          let _34_cf10 = _33___mcc_h5;
          let _35_cf9 = _32___mcc_h4;
          let _36_cf8 = _31___mcc_h3;
          let _37_cf7 = _30___mcc_h2;
          let _38_cf6 = _29___mcc_h1;
          return _37_cf7;
        } else if (_source0.is_DC2) {
          let _39___mcc_h6 = (_source0).cf4;
          let _40_cf4 = _39___mcc_h6;
          return _pat_let_tv2;
        } else {
          let _41___mcc_h7 = (_source0).cf11;
          let _42_cf11 = _41___mcc_h7;
          return (_pat_let_tv3) || (_pat_let_tv4);
        }
      }(_26_v17);
      let _43_v18;
      let _init0 = ((_44_v15) => function (_45_i2) {
        return (_45_i2).multipliedBy((_44_v15).f18);
      })(_24_v15);
      let _nw2 = Array((new BigNumber(19)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw2.length); _i0_0++) {
        _nw2[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _43_v18 = _nw2;
      let _init1 = function (_46_i3) {
        return (_46_i3).multipliedBy(new BigNumber(-986));
      };
      let _nw3 = Array((new BigNumber(17)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw3.length); _i0_1++) {
        _nw3[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _43_v18 = _nw3;
      let _hi0 = (_24_v15).f18;
      for (let _47_i4 = (_24_v15).f18; _47_i4.isLessThan(_hi0); _47_i4 = _47_i4.plus(_dafny.ONE)) {
        let _index0 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_43_v18).length));
        (_43_v18)[_index0] = (_24_v15).f18;
        let _index1 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_43_v18).length));
        (_43_v18)[_index1] = ((_8_v2) ? ((_24_v15).f18) : ((p1).multipliedBy(p1)));
        _6_v0 = _dafny.Seq.Concat(_6_v0, _6_v0);
        let _48_v19;
        _48_v19 = _dafny.Seq.of(new BigNumber(95), p1);
        let _49_v20;
        let _nw4 = Array((new BigNumber(21)).toNumber());
        _nw4[(_dafny.ZERO).toNumber()] = _8_v2;
        _nw4[(_dafny.ONE).toNumber()] = _8_v2;
        _nw4[(new BigNumber(2)).toNumber()] = _8_v2;
        _nw4[(new BigNumber(3)).toNumber()] = _8_v2;
        _nw4[(new BigNumber(4)).toNumber()] = _8_v2;
        _nw4[(new BigNumber(5)).toNumber()] = _8_v2;
        _nw4[(new BigNumber(6)).toNumber()] = _8_v2;
        _nw4[(new BigNumber(7)).toNumber()] = _8_v2;
        _nw4[(new BigNumber(8)).toNumber()] = _8_v2;
        _nw4[(new BigNumber(9)).toNumber()] = _8_v2;
        _nw4[(new BigNumber(10)).toNumber()] = _8_v2;
        _nw4[(new BigNumber(11)).toNumber()] = true;
        _nw4[(new BigNumber(12)).toNumber()] = _8_v2;
        _nw4[(new BigNumber(13)).toNumber()] = _8_v2;
        _nw4[(new BigNumber(14)).toNumber()] = _8_v2;
        _nw4[(new BigNumber(15)).toNumber()] = _8_v2;
        _nw4[(new BigNumber(16)).toNumber()] = _8_v2;
        _nw4[(new BigNumber(17)).toNumber()] = _8_v2;
        _nw4[(new BigNumber(18)).toNumber()] = _8_v2;
        _nw4[(new BigNumber(19)).toNumber()] = _8_v2;
        _nw4[(new BigNumber(20)).toNumber()] = _8_v2;
        _49_v20 = _nw4;
        let _50_v21;
        _50_v21 = _dafny.Seq.of(_49_v20, _49_v20, _49_v20, _49_v20);
        _12_v5 = _dafny.Seq.update(_dafny.Seq.Concat(_12_v5, _dafny.Seq.Concat(_12_v5, _12_v5)), _module.__default.safeIndex((_17_v10).dtor_cf0, new BigNumber((_dafny.Seq.Concat(_12_v5, _dafny.Seq.Concat(_12_v5, _12_v5))).length)), (new BigNumber((_48_v19).length)).isLessThan((_dafny.ZERO).minus(new BigNumber((_50_v21).length))));
        let _source1 = _20_v13;
        if (_source1.is_DC3) {
          let _51___mcc_h8 = (_source1).cf5;
          let _52_cf5 = _51___mcc_h8;
          (globalState).f13 = (_47_i4).isLessThan(new BigNumber(721));
          (globalState).f13 = _8_v2;
          (globalState).f10 = (_24_v15).f18;
          r3 = _8_v2;
        } else if (_source1.is_DC4) {
          let _53___mcc_h9 = (_source1).cf6;
          let _54___mcc_h10 = (_source1).cf7;
          let _55___mcc_h11 = (_source1).cf8;
          let _56___mcc_h12 = (_source1).cf9;
          let _57___mcc_h13 = (_source1).cf10;
          let _58_cf10 = _57___mcc_h13;
          let _59_cf9 = _56___mcc_h12;
          let _60_cf8 = _55___mcc_h11;
          let _61_cf7 = _54___mcc_h10;
          let _62_cf6 = _53___mcc_h9;
          let _63_v22;
          _63_v22 = _dafny.Seq.of(_24_v15, _24_v15, _24_v15);
          _63_v22 = _dafny.Seq.Concat(_63_v22, _63_v22);
          let _64_v23;
          _64_v23 = _dafny.Map.Empty.slice().updateUnsafe(_48_v19,(_24_v15).f18);
          let _index2 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_43_v18).length));
          let _rhs0 = _module.__default.safeDivisionInt((_24_v15).f18, (_dafny.ZERO).minus(_module.__default.fm0(_17_v10, globalState)));
          let _rhs1 = _64_v23;
          let _rhs2 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-273)), ((_65_v15, _66_v9, _67_v0) => function (_68_i5) {
            return new BigNumber((_dafny.Set.fromElements((((_66_v9).contains((_65_v15).f18)) ? ((_66_v9).get((_65_v15).f18)) : (new BigNumber((_67_v0).length))))).length);
          })(_24_v15, _16_v9, _6_v0)), _48_v19), _dafny.Seq.Concat(_48_v19, _48_v19))).length);
          let _lhs0 = _43_v18;
          let _lhs1 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_43_v18).length));
          let _lhs2 = globalState;
          _lhs0[_lhs1] = _rhs0;
          _64_v23 = _rhs1;
          _lhs2.f4 = _rhs2;
          let _index3 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_43_v18).length));
          (_43_v18)[_index3] = _47_i4;
          let _69_v24;
          let _init2 = ((_70_v0) => function (_71_i6) {
            return (_70_v0)[_module.__default.safeIndex(new BigNumber(-327), new BigNumber((_70_v0).length))];
          })(_6_v0);
          let _nw5 = Array((new BigNumber(25)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw5.length); _i0_2++) {
            _nw5[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _69_v24 = _nw5;
          let _72_v25;
          _72_v25 = _dafny.MultiSet.fromElements(_69_v24);
          _72_v25 = _72_v25;
        } else if (_source1.is_DC2) {
          let _73___mcc_h14 = (_source1).cf4;
          let _74_cf4 = _73___mcc_h14;
          let _75_v26;
          let _nw6 = new _module.C0();
          _nw6.__ctor((_47_i4).minus(_module.__default.fm0(_17_v10, globalState)));
          _75_v26 = _nw6;
          r3 = (p1).isLessThan((_24_v15).f18);
          let _76_v27;
          _76_v27 = _dafny.Map.Empty.slice().updateUnsafe(_75_v26,_8_v2);
          let _77_v28;
          let _nw7 = Array((new BigNumber(11)).toNumber());
          _nw7[(_dafny.ZERO).toNumber()] = _76_v27;
          _nw7[(_dafny.ONE).toNumber()] = _76_v27;
          _nw7[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_75_v26,!(_8_v2));
          _nw7[(new BigNumber(3)).toNumber()] = _76_v27;
          _nw7[(new BigNumber(4)).toNumber()] = (_76_v27).update(_75_v26, _8_v2);
          _nw7[(new BigNumber(5)).toNumber()] = (_76_v27).Merge(_76_v27);
          _nw7[(new BigNumber(6)).toNumber()] = _76_v27;
          _nw7[(new BigNumber(7)).toNumber()] = _76_v27;
          _nw7[(new BigNumber(8)).toNumber()] = _76_v27;
          _nw7[(new BigNumber(9)).toNumber()] = _76_v27;
          _nw7[(new BigNumber(10)).toNumber()] = _76_v27;
          _77_v28 = _nw7;
          let _78_v29;
          _78_v29 = _dafny.MultiSet.fromElements(_8_v2);
          let _index4 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_77_v28).length));
          (_77_v28)[_index4] = (_dafny.Map.Empty.slice().updateUnsafe(_75_v26,(_12_v5)[_module.__default.safeIndex(new BigNumber((_78_v29).cardinality()), new BigNumber((_12_v5).length))])).Merge(_76_v27);
          let _index5 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_77_v28).length));
          (_77_v28)[_index5] = _76_v27;
          let _79_v30;
          _79_v30 = _dafny.Map.Empty.slice().updateUnsafe((_24_v15).f18,(_24_v15).f18);
          let _80_v31;
          let _nw8 = Array((new BigNumber(14)).toNumber());
          _nw8[(_dafny.ZERO).toNumber()] = _47_i4;
          _nw8[(_dafny.ONE).toNumber()] = (_43_v18)[_module.__default.safeIndex(new BigNumber(580), new BigNumber((_43_v18).length))];
          _nw8[(new BigNumber(2)).toNumber()] = (_75_v26).f18;
          _nw8[(new BigNumber(3)).toNumber()] = (_24_v15).f18;
          _nw8[(new BigNumber(4)).toNumber()] = (_75_v26).f18;
          _nw8[(new BigNumber(5)).toNumber()] = (_43_v18)[_module.__default.safeIndex(new BigNumber(580), new BigNumber((_43_v18).length))];
          _nw8[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_24_v15).f18);
          _nw8[(new BigNumber(7)).toNumber()] = (_43_v18)[_module.__default.safeIndex(new BigNumber(580), new BigNumber((_43_v18).length))];
          _nw8[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus((_24_v15).f18);
          _nw8[(new BigNumber(9)).toNumber()] = (_24_v15).f18;
          _nw8[(new BigNumber(10)).toNumber()] = (_24_v15).f18;
          _nw8[(new BigNumber(11)).toNumber()] = new BigNumber(((_79_v30).update(p1, new BigNumber((_12_v5).length))).length);
          _nw8[(new BigNumber(12)).toNumber()] = (_43_v18)[_module.__default.safeIndex(new BigNumber(580), new BigNumber((_43_v18).length))];
          _nw8[(new BigNumber(13)).toNumber()] = (_24_v15).f18;
          _80_v31 = _nw8;
          let _81_v32;
          _81_v32 = _module.D3.create_DC10(((_8_v2) ? (true) : (_8_v2)), _80_v31, _dafny.Seq.Concat(_6_v0, _6_v0));
          _81_v32 = _81_v32;
        } else {
          let _82___mcc_h15 = (_source1).cf11;
          let _83_cf11 = _82___mcc_h15;
          (globalState).f0 = _47_i4;
          let _index6 = _module.__default.safeIndex(new BigNumber(960), new BigNumber((_49_v20).length));
          (_49_v20)[_index6] = _8_v2;
          let _index7 = _module.__default.safeIndex(new BigNumber(960), new BigNumber((_49_v20).length));
          (_49_v20)[_index7] = _8_v2;
          let _84_v33;
          let _init3 = ((_85_v7) => function (_86_i7) {
            return _85_v7;
          })(_14_v7);
          let _nw9 = Array((new BigNumber(18)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw9.length); _i0_3++) {
            _nw9[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _84_v33 = _nw9;
          _84_v33 = _84_v33;
          let _index8 = _module.__default.safeIndex(new BigNumber(960), new BigNumber((_49_v20).length));
          (_49_v20)[_index8] = !((_49_v20)[_module.__default.safeIndex(new BigNumber(960), new BigNumber((_49_v20).length))]) || ((_49_v20)[_module.__default.safeIndex(new BigNumber(960), new BigNumber((_49_v20).length))]);
        }
      }
      let _pat_let_tv5 = _7_v1;
      let _pat_let_tv6 = _14_v7;
      let _source2 = function (_pat_let0_0) {
        return function (_87_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_88_dt__update_hcf8_h0) {
              return function (_pat_let2_0) {
                return function (_91_dt__update_hcf6_h0) {
                  return _module.D1.create_DC4(_91_dt__update_hcf6_h0, (_87_dt__update__tmp_h0).dtor_cf7, _88_dt__update_hcf8_h0, (_87_dt__update__tmp_h0).dtor_cf9, (_87_dt__update__tmp_h0).dtor_cf10);
                }(_pat_let2_0);
              }(_dafny.Seq.Create(_module.__default.abs(new BigNumber(335)), ((_89_v7) => function (_90_i8) {
                return _89_v7;
              })(_pat_let_tv6)));
            }(_pat_let1_0);
          }(_pat_let_tv5);
        }(_pat_let0_0);
      }(_20_v13);
      if (_source2.is_DC3) {
        let _92___mcc_h16 = (_source2).cf5;
        let _93_cf5 = _92___mcc_h16;
        let _94_v34;
        let _nw10 = Array((new BigNumber(5)).toNumber()).fill(false);
        _94_v34 = _nw10;
        let _95_v35;
        _95_v35 = _module.D2.create_DC7(_94_v34);
        let _96_v36;
        _96_v36 = _module.D2.create_DC8(_95_v35);
        _96_v36 = _96_v36;
        let _index9 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_94_v34).length));
        (_94_v34)[_index9] = !(_8_v2) || (_8_v2);
        let _index10 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_94_v34).length));
        (_94_v34)[_index10] = _8_v2;
        _94_v34 = ((_8_v2) ? (_94_v34) : (_94_v34));
        let _97_v37;
        _97_v37 = _module.D0.create_DC1(_module.__default.fm3(globalState));
        _97_v37 = p0;
      } else if (_source2.is_DC4) {
        let _98___mcc_h17 = (_source2).cf6;
        let _99___mcc_h18 = (_source2).cf7;
        let _100___mcc_h19 = (_source2).cf8;
        let _101___mcc_h20 = (_source2).cf9;
        let _102___mcc_h21 = (_source2).cf10;
        let _103_cf10 = _102___mcc_h21;
        let _104_cf9 = _101___mcc_h20;
        let _105_cf8 = _100___mcc_h19;
        let _106_cf7 = _99___mcc_h18;
        let _107_cf6 = _98___mcc_h17;
        let _108_v38;
        let _init4 = ((_109_v5, _110_v15, _111_v3, _112_p1) => function (_113_i9) {
          return !(_111_v3).contains((_109_v5)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements((_110_v15).f18, (_110_v15).f18, new BigNumber((_111_v3).length))).length),_112_p1)).length), new BigNumber((_109_v5).length))]);
        })(_12_v5, _24_v15, _9_v3, p1);
        let _nw11 = Array((new BigNumber(27)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw11.length); _i0_4++) {
          _nw11[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _108_v38 = _nw11;
        let _rhs3 = _24_v15;
        let _rhs4 = _108_v38;
        _24_v15 = _rhs3;
        _108_v38 = _rhs4;
        let _114_v39;
        _114_v39 = _dafny.MultiSet.fromElements(_106_cf7, _106_cf7);
        let _index11 = _module.__default.safeIndex(new BigNumber(62), new BigNumber((_43_v18).length));
        (_43_v18)[_index11] = new BigNumber((((_8_v2) ? (_114_v39) : (_module.__default.fm7(p1, _12_v5, _8_v2, globalState)))).cardinality());
        let _index12 = _module.__default.safeIndex(new BigNumber(62), new BigNumber((_43_v18).length));
        (_43_v18)[_index12] = p1;
        let _115_v40;
        _115_v40 = _dafny.Map.Empty.slice().updateUnsafe(_24_v15,(_24_v15).f18);
        _115_v40 = ((_115_v40).Merge(_115_v40)).Merge(_115_v40);
        _14_v7 = _14_v7;
      } else if (_source2.is_DC2) {
        let _116___mcc_h22 = (_source2).cf4;
        let _117_cf4 = _116___mcc_h22;
        let _118_v41;
        _118_v41 = _dafny.Map.Empty.slice().updateUnsafe(_8_v2,true);
        _118_v41 = (_118_v41).Merge(_118_v41);
        if (_8_v2) {
          let _119_v42;
          let _nw12 = Array((new BigNumber(15)).toNumber()).fill(false);
          _119_v42 = _nw12;
          let _120_v43;
          _120_v43 = _dafny.Seq.of(_119_v42);
          let _121_v44;
          _121_v44 = _dafny.Map.Empty.slice().updateUnsafe(_24_v15,_120_v43);
          let _122_v45;
          _122_v45 = _dafny.Map.Empty.slice().updateUnsafe(_8_v2,_43_v18);
          let _123_v46;
          _123_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_122_v45).length),_120_v43);
          let _124_v47;
          let _nw13 = Array((new BigNumber(9)).toNumber());
          _nw13[(_dafny.ZERO).toNumber()] = _120_v43;
          _nw13[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_120_v43, (((_121_v44).contains(_24_v15)) ? ((_121_v44).get(_24_v15)) : (_120_v43)));
          _nw13[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_120_v43, _module.__default.safeIndex(p1, new BigNumber((_120_v43).length)), _119_v42), _120_v43);
          _nw13[(new BigNumber(3)).toNumber()] = _120_v43;
          _nw13[(new BigNumber(4)).toNumber()] = (((_123_v46).contains(p1)) ? ((_123_v46).get(p1)) : (_120_v43));
          _nw13[(new BigNumber(5)).toNumber()] = _120_v43;
          _nw13[(new BigNumber(6)).toNumber()] = _120_v43;
          _nw13[(new BigNumber(7)).toNumber()] = _120_v43;
          _nw13[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_120_v43, _120_v43);
          _124_v47 = _nw13;
          let _index13 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_124_v47).length));
          (_124_v47)[_index13] = _dafny.Seq.Concat(_dafny.Seq.of(_119_v42, _119_v42), _120_v43);
          let _index14 = _module.__default.safeIndex(new BigNumber(368), new BigNumber((_43_v18).length));
          (_43_v18)[_index14] = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_25_v16).length)));
          let _125_v48;
          _125_v48 = _module.D3.create_DC10(!(_8_v2) || (false), _43_v18, _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-28)), ((_126_v7) => function (_127_i10) {
  return _126_v7;
})(_14_v7)), _dafny.Seq.UnicodeFromString("i")), _module.__default.safeIndex((_24_v15).f18, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-28)), ((_128_v7) => function (_129_i10) {
  return _128_v7;
})(_14_v7)), _dafny.Seq.UnicodeFromString("i"))).length)), _14_v7));
          let _130_v49;
          _130_v49 = _dafny.Map.Empty.slice().updateUnsafe(_8_v2,_24_v15);
          let _index15 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_124_v47).length));
          let _index16 = _module.__default.safeIndex(new BigNumber(368), new BigNumber((_43_v18).length));
          let _rhs5 = _dafny.Seq.Concat(_120_v43, _120_v43);
          let _rhs6 = _6_v0;
          let _rhs7 = new BigNumber(((_130_v49).Merge(_130_v49)).length);
          let _rhs8 = _125_v48;
          let _lhs3 = _124_v47;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_124_v47).length));
          let _lhs5 = _43_v18;
          let _lhs6 = _module.__default.safeIndex(new BigNumber(368), new BigNumber((_43_v18).length));
          _lhs3[_lhs4] = _rhs5;
          _6_v0 = _rhs6;
          _lhs5[_lhs6] = _rhs7;
          _125_v48 = _rhs8;
          let _131_v50;
          _131_v50 = _dafny.Map.Empty.slice().updateUnsafe((_43_v18)[_module.__default.safeIndex(new BigNumber(368), new BigNumber((_43_v18).length))],(_24_v15).f18);
          _131_v50 = (_131_v50).update((_43_v18)[_module.__default.safeIndex(new BigNumber(368), new BigNumber((_43_v18).length))], (_24_v15).f18);
          (globalState).f0 = p1;
          _119_v42 = _119_v42;
          let _132_v51;
          _132_v51 = _dafny.MultiSet.fromElements(!(_8_v2), true);
          let _133_v52;
          let _nw14 = new _module.C0();
          _nw14.__ctor(new BigNumber(((_132_v51).Intersect(_132_v51)).cardinality()));
          _133_v52 = _nw14;
        } else {
          let _134_v53;
          let _nw15 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Set.Empty);
          _134_v53 = _nw15;
          _134_v53 = _134_v53;
          let _135_v54;
          _135_v54 = _module.D1.create_DC2(_9_v3);
          let _136_v55;
          _136_v55 = _dafny.Map.Empty.slice().updateUnsafe(_135_v54,(_24_v15).f18);
          _136_v55 = _136_v55;
          let _137_v56;
          _137_v56 = _dafny.Seq.of((_24_v15).f18, (_24_v15).f18, (_24_v15).f18);
          let _138_v57;
          _138_v57 = _dafny.Map.Empty.slice().updateUnsafe(_8_v2,_dafny.Seq.Concat(_137_v56, _137_v56));
          r1 = (((_138_v57).contains(_8_v2)) ? ((_138_v57).get(_8_v2)) : (((_8_v2) ? (_137_v56) : (_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_117_cf4).length)))))));
          let _139_v58;
          let _init5 = ((_140_v2) => function (_141_i11) {
            return _140_v2;
          })(_8_v2);
          let _nw16 = Array((new BigNumber(10)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw16.length); _i0_5++) {
            _nw16[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _139_v58 = _nw16;
          let _index17 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_139_v58).length));
          (_139_v58)[_index17] = true;
          let _index18 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_139_v58).length));
          (_139_v58)[_index18] = _8_v2;
          (globalState).f15 = (_24_v15).f18;
        }
        let _142_v59;
        _142_v59 = _dafny.MultiSet.fromElements(false);
        r3 = ((_dafny.MultiSet.fromElements(_8_v2)).update(!(_8_v2), _module.__default.abs(p1))).IsSubsetOf((_142_v59).Difference(_dafny.MultiSet.fromElements(!(_module.__default.fm1(_8_v2, globalState)), _8_v2)));
        let _143_v60;
        let _nw17 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.of());
        _143_v60 = _nw17;
        let _nw18 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.of());
        _143_v60 = _nw18;
      } else {
        let _144___mcc_h23 = (_source2).cf11;
        let _145_cf11 = _144___mcc_h23;
        let _146_v61;
        _146_v61 = _dafny.Seq.of(_12_v5, _12_v5, _12_v5);
        let _147_v62;
        _147_v62 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-890),_14_v7);
        let _148_v63;
        _148_v63 = _dafny.Seq.of(new BigNumber((_module.__default.fm8(globalState)).length), (_24_v15).f18);
        let _149_v64;
        _149_v64 = _dafny.Seq.of(new BigNumber((_148_v63).length));
        let _150_v65;
        _150_v65 = _module.D3.create_DC10(_8_v2, _43_v18, _6_v0);
        let _151_v66;
        let _nw19 = Array((new BigNumber(22)).toNumber());
        _nw19[(_dafny.ZERO).toNumber()] = true;
        _nw19[(_dafny.ONE).toNumber()] = (_20_v13).dtor_cf7;
        _nw19[(new BigNumber(2)).toNumber()] = _8_v2;
        _nw19[(new BigNumber(3)).toNumber()] = _module.__default.fm1(_8_v2, globalState);
        _nw19[(new BigNumber(4)).toNumber()] = _dafny.Seq.IsPrefixOf(_146_v61, _dafny.Seq.of(_12_v5));
        _nw19[(new BigNumber(5)).toNumber()] = !(true);
        _nw19[(new BigNumber(6)).toNumber()] = _8_v2;
        _nw19[(new BigNumber(7)).toNumber()] = _8_v2;
        _nw19[(new BigNumber(8)).toNumber()] = !((_24_v15).f18).isEqualTo((_24_v15).f18);
        _nw19[(new BigNumber(9)).toNumber()] = _module.__default.fm1(_8_v2, globalState);
        _nw19[(new BigNumber(10)).toNumber()] = _8_v2;
        _nw19[(new BigNumber(11)).toNumber()] = (_147_v62).equals(_dafny.Map.Empty.slice().updateUnsafe((_149_v64)[_module.__default.safeIndex(new BigNumber((_6_v0).length), new BigNumber((_149_v64).length))],_14_v7));
        _nw19[(new BigNumber(12)).toNumber()] = !(_8_v2);
        _nw19[(new BigNumber(13)).toNumber()] = true;
        _nw19[(new BigNumber(14)).toNumber()] = (_12_v5)[_module.__default.safeIndex(p1, new BigNumber((_12_v5).length))];
        _nw19[(new BigNumber(15)).toNumber()] = _8_v2;
        _nw19[(new BigNumber(16)).toNumber()] = (p1).isLessThan((_24_v15).f18);
        _nw19[(new BigNumber(17)).toNumber()] = !(_8_v2);
        _nw19[(new BigNumber(18)).toNumber()] = (_150_v65).dtor_cf16;
        _nw19[(new BigNumber(19)).toNumber()] = _8_v2;
        _nw19[(new BigNumber(20)).toNumber()] = _8_v2;
        _nw19[(new BigNumber(21)).toNumber()] = _module.__default.fm1(!(false), globalState);
        _151_v66 = _nw19;
        _151_v66 = _151_v66;
        let _152_v67;
        let _nw20 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.of());
        _152_v67 = _nw20;
        _152_v67 = _152_v67;
        r3 = _8_v2;
        let _153_v68;
        _153_v68 = _dafny.Map.Empty.slice().updateUnsafe(_151_v66,_24_v15);
        _153_v68 = _153_v68;
      }
      let _pat_let_tv7 = _24_v15;
      r0 = (_dafny.ZERO).minus(function (_source3) {
        if (_source3.is_DC0) {
          let _154___mcc_h24 = (_source3).cf0;
          let _155___mcc_h25 = (_source3).cf1;
          let _156___mcc_h26 = (_source3).cf2;
          let _157_cf2 = _156___mcc_h26;
          let _158_cf1 = _155___mcc_h25;
          let _159_cf0 = _154___mcc_h24;
          return (_dafny.ZERO).minus((_159_cf0).multipliedBy(new BigNumber(483)));
        } else {
          let _160___mcc_h27 = (_source3).cf3;
          let _161_cf3 = _160___mcc_h27;
          return (_pat_let_tv7).f18;
        }
      }(p0));
      let _162_v69;
      _162_v69 = _dafny.Seq.of((_24_v15).f18, new BigNumber((_dafny.Seq.UnicodeFromString("jxu")).length));
      r1 = _162_v69;
      let _163_v70;
      _163_v70 = _dafny.Set.fromElements(p1, _module.__default.fm0(_17_v10, globalState));
      r2 = (_163_v70).Union(_163_v70);
      r3 = _8_v2;
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _164_v0;
      _164_v0 = true;
      let _165_v1;
      _165_v1 = _dafny.Set.fromElements(_164_v0);
      let _166_v2;
      let _nw21 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Map.Empty);
      _166_v2 = _nw21;
      let _167_v3;
      _167_v3 = new BigNumber(639);
      let _168_v4;
      _168_v4 = _dafny.MultiSet.fromElements(_167_v3, new BigNumber(-408));
      let _169_v5;
      _169_v5 = _dafny.MultiSet.fromElements(_164_v0, _164_v0);
      let _170_v6;
      _170_v6 = _dafny.Seq.of(new BigNumber((_168_v4).cardinality()), new BigNumber((_169_v5).cardinality()));
      let _171_v7;
      let _init6 = ((_172_v0) => function (_173_i0) {
        return _172_v0;
      })(_164_v0);
      let _nw22 = Array((new BigNumber(11)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw22.length); _i0_6++) {
        _nw22[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _171_v7 = _nw22;
      let _174_v8;
      _174_v8 = _dafny.Seq.UnicodeFromString("mgx");
      let _175_globalState;
      let _nw23 = new _module.GlobalState();
      _nw23.__ctor(new BigNumber(-206), (_165_v1).Intersect(_165_v1), false, false, new BigNumber(633), new BigNumber(769), true, _166_v2, _dafny.Seq.Concat(_dafny.Seq.update(_170_v6, _module.__default.safeIndex(_167_v3, new BigNumber((_170_v6).length)), _167_v3), _170_v6), _168_v4, new BigNumber(-738), new BigNumber(714), new _dafny.CodePoint('r'.codePointAt(0)), false, new BigNumber(760), new BigNumber(654), _171_v7, _174_v8);
      _175_globalState = _nw23;
      let _176_v9;
      _176_v9 = _dafny.Map.Empty.slice().updateUnsafe(_167_v3,_167_v3);
      let _177_v10;
      _177_v10 = _dafny.Map.Empty.slice().updateUnsafe(_164_v0,new BigNumber((_176_v9).length));
      let _178_v11;
      _178_v11 = _dafny.Map.Empty.slice().updateUnsafe(_174_v8,(((_177_v10).contains(_164_v0)) ? ((_177_v10).get(_164_v0)) : (_167_v3)));
      let _179_v12;
      _179_v12 = _module.D0.create_DC0(_167_v3, _178_v11, _168_v4);
      let _source4 = _179_v12;
      if (_source4.is_DC0) {
        let _180___mcc_h0 = (_source4).cf0;
        let _181___mcc_h1 = (_source4).cf1;
        let _182___mcc_h2 = (_source4).cf2;
        let _183_cf2 = _182___mcc_h2;
        let _184_cf1 = _181___mcc_h1;
        let _185_cf0 = _180___mcc_h0;
        let _186_v13;
        _186_v13 = _module.D0.create_DC1(_module.D0.create_DC1(_module.D0.create_DC0(new BigNumber((_170_v6).length), _178_v11, _183_cf2)));
        let _187_v14;
        _187_v14 = _module.D0.create_DC1(_186_v13);
        let _188_v15;
        let _189_v16;
        let _190_v17;
        let _191_v18;
        let _out0;
        let _out1;
        let _out2;
        let _out3;
        let _outcollector0 = _module.__default.m0(_module.D0.create_DC1(_187_v14), _module.__default.fm0(_179_v12, _175_globalState), _175_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _out3 = _outcollector0[3];
        _188_v15 = _out0;
        _189_v16 = _out1;
        _190_v17 = _out2;
        _191_v18 = _out3;
        _164_v0 = !((_module.__default.fm0(_179_v12, _175_globalState)).isEqualTo(((_dafny.ZERO).minus(_167_v3)).minus(_185_cf0)));
        if (_module.__default.fm1(_164_v0, _175_globalState)) {
          let _192_v19;
          _192_v19 = _module.D0.create_DC1(_186_v13);
          let _193_v20;
          let _194_v21;
          let _195_v22;
          let _196_v23;
          let _out4;
          let _out5;
          let _out6;
          let _out7;
          let _outcollector1 = _module.__default.m0(((_164_v0) ? (_192_v19) : (_192_v19)), _module.__default.safeModuloInt(new BigNumber((_174_v8).length), _185_cf0), _175_globalState);
          _out4 = _outcollector1[0];
          _out5 = _outcollector1[1];
          _out6 = _outcollector1[2];
          _out7 = _outcollector1[3];
          _193_v20 = _out4;
          _194_v21 = _out5;
          _195_v22 = _out6;
          _196_v23 = _out7;
          (_175_globalState).f13 = _module.__default.fm1(_196_v23, _175_globalState);
          (_175_globalState).f13 = _191_v18;
          _177_v10 = (_177_v10).update(_dafny.Seq.IsProperPrefixOf(_174_v8, _174_v8), _167_v3);
          (_175_globalState).f0 = _167_v3;
        } else {
          _179_v12 = _179_v12;
          let _197_v24;
          _197_v24 = _module.D0.create_DC1(_186_v13);
          let _198_v25;
          let _199_v26;
          let _200_v27;
          let _201_v28;
          let _out8;
          let _out9;
          let _out10;
          let _out11;
          let _outcollector2 = _module.__default.m0(_197_v24, (new BigNumber((_174_v8).length)).multipliedBy(_167_v3), _175_globalState);
          _out8 = _outcollector2[0];
          _out9 = _outcollector2[1];
          _out10 = _outcollector2[2];
          _out11 = _outcollector2[3];
          _198_v25 = _out8;
          _199_v26 = _out9;
          _200_v27 = _out10;
          _201_v28 = _out11;
          (_175_globalState).f13 = true;
          _167_v3 = _module.__default.safeModuloInt(new BigNumber(472), new BigNumber((_174_v8).length));
          let _202_v29;
          let _nw24 = new _module.C0();
          _nw24.__ctor(_198_v25);
          _202_v29 = _nw24;
        }
        let _203_v30;
        let _nw25 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
        _203_v30 = _nw25;
        let _index19 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_203_v30).length));
        (_203_v30)[_index19] = (_167_v3).minus(_185_cf0);
        let _index20 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_203_v30).length));
        let _rhs9 = _185_cf0;
        let _rhs10 = _167_v3;
        let _rhs11 = !((new BigNumber((_176_v9).length)).isLessThan((_185_cf0).multipliedBy(_188_v15)));
        let _lhs7 = _203_v30;
        let _lhs8 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_203_v30).length));
        _lhs7[_lhs8] = _rhs9;
        _167_v3 = _rhs10;
        _164_v0 = _rhs11;
      } else {
        let _204___mcc_h3 = (_source4).cf3;
        let _205_cf3 = _204___mcc_h3;
        let _source5 = ((false) ? (_179_v12) : (_179_v12));
        if (_source5.is_DC0) {
          let _206___mcc_h4 = (_source5).cf0;
          let _207___mcc_h5 = (_source5).cf1;
          let _208___mcc_h6 = (_source5).cf2;
          let _209_cf2 = _208___mcc_h6;
          let _210_cf1 = _207___mcc_h5;
          let _211_cf0 = _206___mcc_h4;
          let _212_v31;
          _212_v31 = _dafny.Map.Empty.slice().updateUnsafe((((_169_v5).contains(_164_v0)) ? ((_169_v5).get(_164_v0)) : (_167_v3)),_164_v0);
          _164_v0 = _module.__default.fm1((((_212_v31).contains(_211_cf0)) ? ((_212_v31).get(_211_cf0)) : (_164_v0)), _175_globalState);
          (_175_globalState).f15 = (((_165_v1).IsDisjointFrom(_165_v1)) ? ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_164_v0,_164_v0)).length)).plus(new BigNumber((_174_v8).length))) : ((_dafny.ZERO).minus(_167_v3)));
          let _213_v33;
          _213_v33 = _dafny.Map.Empty.slice().updateUnsafe(_174_v8,_dafny.Seq.of(!(_module.__default.fm1(false, _175_globalState))));
          let _214_v34;
          _214_v34 = _module.D0.create_DC0(_211_cf0, function () {
  let _coll4 = new _dafny.Map();
  for (const _compr_4 of (_213_v33).Keys.Elements) {
    let _215_v32 = _compr_4;
    if ((_213_v33).contains(_215_v32)) {
      _coll4.push([_215_v32,(_179_v12).dtor_cf0]);
    }
  }
  return _coll4;
}(), (_168_v4).update((((_168_v4).contains(_211_cf0)) ? ((_168_v4).get(_211_cf0)) : (_167_v3)), _module.__default.abs(_167_v3)));
          let _216_v35;
          _216_v35 = _module.D0.create_DC1(_214_v34);
          let _217_v36;
          _217_v36 = _module.D0.create_DC1(_216_v35);
          let _218_v37;
          _218_v37 = _module.D0.create_DC1(_217_v36);
          let _219_v38;
          _219_v38 = _dafny.Set.fromElements(_174_v8);
          let _220_v39;
          let _221_v40;
          let _222_v41;
          let _223_v42;
          let _out12;
          let _out13;
          let _out14;
          let _out15;
          let _outcollector3 = _module.__default.m0(_218_v37, new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_211_cf0))), _167_v3, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(830)), ((_224_v8, _225_v3) => function (_226_i1) {
            return (_224_v8)[_module.__default.safeIndex(_225_v3, new BigNumber((_224_v8).length))];
          })(_174_v8, _167_v3))).length)), _167_v3, (_dafny.ZERO).minus(new BigNumber((_219_v38).length)))).length), _175_globalState);
          _out12 = _outcollector3[0];
          _out13 = _outcollector3[1];
          _out14 = _outcollector3[2];
          _out15 = _outcollector3[3];
          _220_v39 = _out12;
          _221_v40 = _out13;
          _222_v41 = _out14;
          _223_v42 = _out15;
          let _227_v43;
          let _nw26 = new _module.C0();
          _nw26.__ctor(_211_cf0);
          _227_v43 = _nw26;
          _227_v43 = _227_v43;
        } else {
          let _228___mcc_h7 = (_source5).cf3;
          let _229_cf3 = _228___mcc_h7;
          let _230_v44;
          let _nw27 = new _module.C0();
          _nw27.__ctor((_167_v3).minus(_167_v3));
          _230_v44 = _nw27;
          _230_v44 = _230_v44;
          let _231_v45;
          _231_v45 = _module.D0.create_DC0(_167_v3, _178_v11, _168_v4);
          let _232_v46;
          let _233_v47;
          let _234_v48;
          let _235_v49;
          let _out16;
          let _out17;
          let _out18;
          let _out19;
          let _outcollector4 = _module.__default.m0(_module.D0.create_DC1(_231_v45), (_module.D0.create_DC0(new BigNumber((_170_v6).length), _178_v11, _module.__default.fm2(_164_v0, _175_globalState))).dtor_cf0, _175_globalState);
          _out16 = _outcollector4[0];
          _out17 = _outcollector4[1];
          _out18 = _outcollector4[2];
          _out19 = _outcollector4[3];
          _232_v46 = _out16;
          _233_v47 = _out17;
          _234_v48 = _out18;
          _235_v49 = _out19;
          let _236_v50;
          _236_v50 = _module.D0.create_DC1(_231_v45);
          let _237_v51;
          let _238_v52;
          let _239_v53;
          let _240_v54;
          let _out20;
          let _out21;
          let _out22;
          let _out23;
          let _outcollector5 = _module.__default.m0(_236_v50, new BigNumber((_169_v5).cardinality()), _175_globalState);
          _out20 = _outcollector5[0];
          _out21 = _outcollector5[1];
          _out22 = _outcollector5[2];
          _out23 = _outcollector5[3];
          _237_v51 = _out20;
          _238_v52 = _out21;
          _239_v53 = _out22;
          _240_v54 = _out23;
          let _pat_let_tv8 = _231_v45;
          let _241_v55;
          let _242_v56;
          let _243_v57;
          let _244_v58;
          let _out24;
          let _out25;
          let _out26;
          let _out27;
          let _outcollector6 = _module.__default.m0(function (_pat_let3_0) {
            return function (_245_dt__update__tmp_h0) {
              return function (_pat_let4_0) {
                return function (_246_dt__update_hcf3_h0) {
                  return _module.D0.create_DC1(_246_dt__update_hcf3_h0);
                }(_pat_let4_0);
              }(_pat_let_tv8);
            }(_pat_let3_0);
          }(_236_v50), new BigNumber(((_dafny.Set.fromElements((_230_v44).f18, _232_v46)).Difference(_234_v48)).length), _175_globalState);
          _out24 = _outcollector6[0];
          _out25 = _outcollector6[1];
          _out26 = _outcollector6[2];
          _out27 = _outcollector6[3];
          _241_v55 = _out24;
          _242_v56 = _out25;
          _243_v57 = _out26;
          _244_v58 = _out27;
        }
        let _247_v59;
        _247_v59 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(false), _dafny.MultiSet.fromElements(_164_v0), _169_v5, _169_v5);
        let _248_v60;
        _248_v60 = _dafny.Seq.of(_169_v5);
        _247_v59 = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_248_v60, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-680)), ((_249_v0) => function (_250_i2) {
          return _dafny.MultiSet.FromArray(_dafny.Seq.of(_249_v0, false, _249_v0));
        })(_164_v0))));
        (_175_globalState).f0 = new BigNumber(12);
        (_175_globalState).f10 = _167_v3;
      }
      let _251_v62;
      _251_v62 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_170_v6), _168_v4, _168_v4, _168_v4);
      let _252_v63;
      let _nw28 = new _module.C0();
      _nw28.__ctor(new BigNumber((_168_v4).cardinality()));
      _252_v63 = _nw28;
      let _253_v64;
      _253_v64 = _dafny.Map.Empty.slice().updateUnsafe(_164_v0,_252_v63);
      let _254_v65;
      _254_v65 = _module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(!(_module.__default.fm1(false, _175_globalState)),_167_v3));
      let _255_v66;
      let _nw29 = Array((new BigNumber(14)).toNumber());
      _nw29[(_dafny.ZERO).toNumber()] = _module.__default.fm0(_module.__default.fm3(_175_globalState), _175_globalState);
      _nw29[(_dafny.ONE).toNumber()] = _167_v3;
      _nw29[(new BigNumber(2)).toNumber()] = (_167_v3).plus(_167_v3);
      _nw29[(new BigNumber(3)).toNumber()] = _167_v3;
      _nw29[(new BigNumber(4)).toNumber()] = _167_v3;
      _nw29[(new BigNumber(5)).toNumber()] = new BigNumber(-437);
      _nw29[(new BigNumber(6)).toNumber()] = (_170_v6)[_module.__default.safeIndex(_module.__default.fm0(_module.D0.create_DC0(_167_v3, _dafny.Map.Empty.slice().updateUnsafe(_174_v8,_167_v3), _168_v4), _175_globalState), new BigNumber((_170_v6).length))];
      _nw29[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_251_v62).Elements) {
          let _256_v61 = _compr_5;
          if (_dafny.Seq.contains(_251_v62, _256_v61)) {
            _coll5.push([_256_v61,_167_v3]);
          }
        }
        return _coll5;
      }()).length));
      _nw29[(new BigNumber(8)).toNumber()] = _167_v3;
      _nw29[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_174_v8, _174_v8)).length);
      _nw29[(new BigNumber(10)).toNumber()] = new BigNumber((_253_v64).length);
      _nw29[(new BigNumber(11)).toNumber()] = new BigNumber(((_254_v65).dtor_cf4).length);
      _nw29[(new BigNumber(12)).toNumber()] = (new BigNumber((_177_v10).length)).plus((_252_v63).f18);
      _nw29[(new BigNumber(13)).toNumber()] = _167_v3;
      _255_v66 = _nw29;
      let _index21 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_255_v66).length));
      (_255_v66)[_index21] = ((_252_v63).f18).multipliedBy(new BigNumber((_174_v8).length));
      let _index22 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_255_v66).length));
      (_255_v66)[_index22] = ((_252_v63).f18).multipliedBy(_167_v3);
      let _257_v67;
      _257_v67 = _dafny.Set.fromElements((_252_v63).f18, (_252_v63).f18);
      let _258_v68;
      _258_v68 = _dafny.Seq.of(_164_v0);
      let _index23 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_255_v66).length));
      let _index24 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_255_v66).length));
      let _rhs12 = ((_164_v0) ? (_167_v3) : (new BigNumber((_257_v67).length)));
      let _rhs13 = _module.__default.safeModuloInt(((_252_v63).f18).minus(new BigNumber((_174_v8).length)), new BigNumber((_258_v68).length));
      let _rhs14 = (_252_v63).f18;
      let _lhs9 = _175_globalState;
      let _lhs10 = _255_v66;
      let _lhs11 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_255_v66).length));
      let _lhs12 = _255_v66;
      let _lhs13 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_255_v66).length));
      _lhs9.f10 = _rhs12;
      _lhs10[_lhs11] = _rhs13;
      _lhs12[_lhs13] = _rhs14;
      let _259_v69;
      let _nw30 = new _module.C0();
      _nw30.__ctor(_167_v3);
      _259_v69 = _nw30;
      _168_v4 = _168_v4;
      let _260_v70;
      _260_v70 = _dafny.Map.Empty.slice().updateUnsafe((_252_v63).f18,!(_164_v0));
      if ((((_260_v70).contains(_module.__default.safeModuloInt((_252_v63).f18, new BigNumber(-905)))) ? ((_260_v70).get(_module.__default.safeModuloInt((_252_v63).f18, new BigNumber(-905)))) : (!(_module.__default.fm1(_164_v0, _175_globalState))))) {
        let _261_v71;
        let _init7 = ((_262_v0, _263_v10) => function (_264_i3) {
          return ((_262_v0) ? (_263_v10) : (_263_v10));
        })(_164_v0, _177_v10);
        let _nw31 = Array((new BigNumber(4)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw31.length); _i0_7++) {
          _nw31[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _261_v71 = _nw31;
        let _index25 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_261_v71).length));
        (_261_v71)[_index25] = _177_v10;
        let _index26 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_261_v71).length));
        (_261_v71)[_index26] = _177_v10;
        let _265_v72;
        let _nw32 = Array((new BigNumber(13)).toNumber()).fill(_dafny.MultiSet.Empty);
        _265_v72 = _nw32;
        let _266_v73;
        let _nw33 = Array((new BigNumber(2)).toNumber());
        _nw33[(_dafny.ZERO).toNumber()] = _254_v65;
        _nw33[(_dafny.ONE).toNumber()] = _254_v65;
        _266_v73 = _nw33;
        let _index27 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_265_v72).length));
        (_265_v72)[_index27] = _dafny.MultiSet.fromElements(_266_v73, _266_v73, _266_v73, _266_v73, _266_v73);
        let _267_v74;
        _267_v74 = _dafny.MultiSet.fromElements(_266_v73, _266_v73);
        let _index28 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_265_v72).length));
        (_265_v72)[_index28] = ((_267_v74).Difference(_dafny.MultiSet.fromElements(_266_v73, _266_v73))).Intersect((_267_v74).Intersect(_267_v74));
        let _268_v75;
        let _nw34 = Array((new BigNumber(25)).toNumber());
        _268_v75 = _nw34;
        let _index29 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_268_v75).length));
        (_268_v75)[_index29] = _252_v63;
        let _index30 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_268_v75).length));
        (_268_v75)[_index30] = _252_v63;
        let _269_v76;
        let _nw35 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Map.Empty);
        _269_v76 = _nw35;
        _269_v76 = _269_v76;
        let _270_v77;
        _270_v77 = _module.D1.create_DC3(_165_v1);
        let _271_v78;
        let _nw36 = Array((new BigNumber(16)).toNumber());
        _nw36[(_dafny.ZERO).toNumber()] = _module.D1.create_DC3(_165_v1);
        _nw36[(_dafny.ONE).toNumber()] = _270_v77;
        _nw36[(new BigNumber(2)).toNumber()] = _270_v77;
        _nw36[(new BigNumber(3)).toNumber()] = _270_v77;
        _nw36[(new BigNumber(4)).toNumber()] = _270_v77;
        _nw36[(new BigNumber(5)).toNumber()] = _270_v77;
        _nw36[(new BigNumber(6)).toNumber()] = _module.D1.create_DC3(_165_v1);
        _nw36[(new BigNumber(7)).toNumber()] = _module.D1.create_DC3(_165_v1);
        _nw36[(new BigNumber(8)).toNumber()] = _270_v77;
        _nw36[(new BigNumber(9)).toNumber()] = _270_v77;
        _nw36[(new BigNumber(10)).toNumber()] = ((true) ? (_270_v77) : (_270_v77));
        _nw36[(new BigNumber(11)).toNumber()] = _270_v77;
        _nw36[(new BigNumber(12)).toNumber()] = _270_v77;
        _nw36[(new BigNumber(13)).toNumber()] = _module.D1.create_DC3(_module.__default.fm4((_252_v63).f18, (_259_v69).f18, (_255_v66)[_module.__default.safeIndex(new BigNumber(580), new BigNumber((_255_v66).length))], (_252_v63).f18, _175_globalState));
        _nw36[(new BigNumber(14)).toNumber()] = _module.D1.create_DC3(_165_v1);
        _nw36[(new BigNumber(15)).toNumber()] = _270_v77;
        _271_v78 = _nw36;
        let _index31 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_271_v78).length));
        (_271_v78)[_index31] = _270_v77;
        let _index32 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_271_v78).length));
        (_271_v78)[_index32] = _270_v77;
      } else {
        let _index33 = _module.__default.safeIndex(new BigNumber(826), new BigNumber((_171_v7).length));
        (_171_v7)[_index33] = _164_v0;
        let _index34 = _module.__default.safeIndex(new BigNumber(826), new BigNumber((_171_v7).length));
        (_171_v7)[_index34] = _164_v0;
        let _272_v79;
        _272_v79 = _module.D0.create_DC0((_252_v63).f18, ((_178_v11).update(_174_v8, new BigNumber((_dafny.MultiSet.FromArray(_258_v68)).cardinality()))).update(_174_v8, (_252_v63).f18), _168_v4);
        let _273_v80;
        _273_v80 = _module.D0.create_DC1(_272_v79);
        let _274_v81;
        _274_v81 = _module.D0.create_DC1(_272_v79);
        let _275_v82;
        let _276_v83;
        let _277_v84;
        let _278_v85;
        let _out28;
        let _out29;
        let _out30;
        let _out31;
        let _outcollector7 = _module.__default.m0(_274_v81, ((_164_v0) ? ((_252_v63).f18) : (new BigNumber(-94))), _175_globalState);
        _out28 = _outcollector7[0];
        _out29 = _outcollector7[1];
        _out30 = _outcollector7[2];
        _out31 = _outcollector7[3];
        _275_v82 = _out28;
        _276_v83 = _out29;
        _277_v84 = _out30;
        _278_v85 = _out31;
        let _279_v86;
        let _nw37 = new _module.C0();
        _nw37.__ctor(_167_v3);
        _279_v86 = _nw37;
        let _280_v87;
        _280_v87 = _dafny.Map.Empty.slice().updateUnsafe((_171_v7)[_module.__default.safeIndex(new BigNumber(826), new BigNumber((_171_v7).length))],_274_v81);
        _280_v87 = (_280_v87).update(_278_v85, _module.D0.create_DC1(_272_v79));
        (_175_globalState).f13 = _164_v0;
      }
      _176_v9 = (_176_v9).update((_252_v63).f18, (_252_v63).f18);
      let _281_v88;
      _281_v88 = new _dafny.CodePoint('j'.codePointAt(0));
      let _282_v89;
      _282_v89 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((_259_v69).f18, (_255_v66)[_module.__default.safeIndex(new BigNumber(580), new BigNumber((_255_v66).length))], _167_v3, _167_v3, new BigNumber(-242)),_281_v88);
      let _283_v90;
      _283_v90 = _dafny.Map.Empty.slice().updateUnsafe(_164_v0,_282_v89);
      _282_v89 = (((_283_v90).contains(true)) ? ((_283_v90).get(true)) : (_282_v89));
      (_175_globalState).f4 = (_dafny.ZERO).minus(_167_v3);
      _164_v0 = !(_164_v0);
      let _284_v91;
      let _nw38 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.of());
      _284_v91 = _nw38;
      let _index35 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_284_v91).length));
      (_284_v91)[_index35] = _170_v6;
      let _index36 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_284_v91).length));
      (_284_v91)[_index36] = _dafny.Seq.Concat(_170_v6, _170_v6);
      _171_v7 = _171_v7;
      _164_v0 = _164_v0;
      if (((_255_v66)[_module.__default.safeIndex(new BigNumber(580), new BigNumber((_255_v66).length))]).isEqualTo(new BigNumber((_174_v8).length))) {
        let _285_v92;
        _285_v92 = _dafny.Seq.of(_259_v69, _259_v69, _252_v63, _259_v69);
        let _286_v93;
        _286_v93 = _dafny.Map.Empty.slice().updateUnsafe(_281_v88,_dafny.Seq.IsProperPrefixOf(_285_v92, _dafny.Seq.of(_252_v63, _252_v63)));
        _286_v93 = (_286_v93).update(_281_v88, _164_v0);
        if (_164_v0) {
          let _287_v94;
          _287_v94 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(54),_179_v12);
          let _288_v96;
          _288_v96 = _dafny.Seq.of(_257_v67, _257_v67, function () {
            let _coll6 = new _dafny.Set();
            for (const _compr_6 of (_287_v94).Keys.Elements) {
              let _289_v95 = _compr_6;
              if ((_287_v94).contains(_289_v95)) {
                _coll6.add((_289_v95).plus(new BigNumber(514)));
              }
            }
            return _coll6;
          }());
          _288_v96 = _dafny.Seq.Concat(_288_v96, _288_v96);
          let _290_v97;
          _290_v97 = _module.D1.create_DC3(_165_v1);
          let _291_v98;
          _291_v98 = _dafny.Map.Empty.slice().updateUnsafe(_165_v1,_174_v8);
          let _pat_let_tv9 = _165_v1;
          _164_v0 = !(_291_v98).contains((function (_pat_let5_0) {
            return function (_292_dt__update__tmp_h1) {
              return function (_pat_let6_0) {
                return function (_293_dt__update_hcf5_h0) {
                  return _module.D1.create_DC3(_293_dt__update_hcf5_h0);
                }(_pat_let6_0);
              }(_pat_let_tv9);
            }(_pat_let5_0);
          }(_290_v97)).dtor_cf5);
          let _294_v99;
          _294_v99 = _module.D0.create_DC0((_259_v69).f18, _178_v11, _168_v4);
          let _295_v100;
          _295_v100 = _module.D0.create_DC1(_294_v99);
          let _296_v101;
          let _297_v102;
          let _298_v103;
          let _299_v104;
          let _out32;
          let _out33;
          let _out34;
          let _out35;
          let _outcollector8 = _module.__default.m0(_295_v100, _module.__default.fm0(_179_v12, _175_globalState), _175_globalState);
          _out32 = _outcollector8[0];
          _out33 = _outcollector8[1];
          _out34 = _outcollector8[2];
          _out35 = _outcollector8[3];
          _296_v101 = _out32;
          _297_v102 = _out33;
          _298_v103 = _out34;
          _299_v104 = _out35;
          let _300_v105;
          _300_v105 = _dafny.Map.Empty.slice().updateUnsafe(_296_v101,_285_v92);
          let _301_v106;
          _301_v106 = _dafny.Map.Empty.slice().updateUnsafe(_164_v0,_285_v92);
          let _302_v107;
          _302_v107 = _dafny.MultiSet.fromElements((((_300_v105).contains(_167_v3)) ? ((_300_v105).get(_167_v3)) : ((((_301_v106).contains(_164_v0)) ? ((_301_v106).get(_164_v0)) : (_285_v92)))), _285_v92);
          let _index37 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_255_v66).length));
          (_255_v66)[_index37] = (new BigNumber(((_302_v107).Intersect(_302_v107)).cardinality())).minus(new BigNumber((_dafny.Seq.UnicodeFromString("tprjqefl")).length));
          _285_v92 = _285_v92;
        } else {
          (_175_globalState).f13 = (new BigNumber((_285_v92).length)).isLessThan((_259_v69).f18);
          let _303_v108;
          _303_v108 = _dafny.Seq.of(_174_v8, _174_v8);
          let _304_v109;
          _304_v109 = _module.D1.create_DC4(_dafny.Seq.UnicodeFromString("cs"), _164_v0, _303_v108, _module.__default.fm5(new BigNumber((_168_v4).cardinality()), ((_284_v91)[_module.__default.safeIndex(new BigNumber(747), new BigNumber((_284_v91).length))])[_module.__default.safeIndex(_module.__default.fm0(_179_v12, _175_globalState), new BigNumber(((_284_v91)[_module.__default.safeIndex(new BigNumber(747), new BigNumber((_284_v91).length))]).length))], (_259_v69).f18, (_252_v63).f18, _175_globalState), _281_v88);
          let _index38 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_255_v66).length));
          let _rhs15 = new BigNumber((_dafny.Seq.Concat(_174_v8, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(650)), ((_305_v88) => function (_306_i4) {
            return _305_v88;
          })(_281_v88)), _174_v8))).length);
          let _rhs16 = _259_v69;
          let _rhs17 = _260_v70;
          let _rhs18 = (_304_v109).dtor_cf7;
          let _rhs19 = (_252_v63).f18;
          let _lhs14 = _175_globalState;
          let _lhs15 = _175_globalState;
          let _lhs16 = _255_v66;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_255_v66).length));
          _lhs14.f15 = _rhs15;
          _252_v63 = _rhs16;
          _260_v70 = _rhs17;
          _lhs15.f13 = _rhs18;
          _lhs16[_lhs17] = _rhs19;
          let _307_v110;
          _307_v110 = _module.D2.create_DC6(_259_v69);
          let _308_v111;
          let _nw39 = Array((new BigNumber(29)).toNumber());
          _nw39[(_dafny.ZERO).toNumber()] = _259_v69;
          _nw39[(_dafny.ONE).toNumber()] = ((!(_164_v0)) ? (_259_v69) : ((_307_v110).dtor_cf12));
          _nw39[(new BigNumber(2)).toNumber()] = _259_v69;
          _nw39[(new BigNumber(3)).toNumber()] = _252_v63;
          _nw39[(new BigNumber(4)).toNumber()] = _252_v63;
          _nw39[(new BigNumber(5)).toNumber()] = _252_v63;
          _nw39[(new BigNumber(6)).toNumber()] = _252_v63;
          _nw39[(new BigNumber(7)).toNumber()] = _252_v63;
          _nw39[(new BigNumber(8)).toNumber()] = _252_v63;
          _nw39[(new BigNumber(9)).toNumber()] = _252_v63;
          _nw39[(new BigNumber(10)).toNumber()] = _259_v69;
          _nw39[(new BigNumber(11)).toNumber()] = (_285_v92)[_module.__default.safeIndex((_255_v66)[_module.__default.safeIndex(new BigNumber(580), new BigNumber((_255_v66).length))], new BigNumber((_285_v92).length))];
          _nw39[(new BigNumber(12)).toNumber()] = _252_v63;
          _nw39[(new BigNumber(13)).toNumber()] = _252_v63;
          _nw39[(new BigNumber(14)).toNumber()] = _259_v69;
          _nw39[(new BigNumber(15)).toNumber()] = _259_v69;
          _nw39[(new BigNumber(16)).toNumber()] = _252_v63;
          _nw39[(new BigNumber(17)).toNumber()] = _252_v63;
          _nw39[(new BigNumber(18)).toNumber()] = (_285_v92)[_module.__default.safeIndex((_259_v69).f18, new BigNumber((_285_v92).length))];
          _nw39[(new BigNumber(19)).toNumber()] = _252_v63;
          _nw39[(new BigNumber(20)).toNumber()] = (_285_v92)[_module.__default.safeIndex(new BigNumber(917), new BigNumber((_285_v92).length))];
          _nw39[(new BigNumber(21)).toNumber()] = _259_v69;
          _nw39[(new BigNumber(22)).toNumber()] = _259_v69;
          _nw39[(new BigNumber(23)).toNumber()] = _252_v63;
          _nw39[(new BigNumber(24)).toNumber()] = _259_v69;
          _nw39[(new BigNumber(25)).toNumber()] = _252_v63;
          _nw39[(new BigNumber(26)).toNumber()] = _252_v63;
          _nw39[(new BigNumber(27)).toNumber()] = _259_v69;
          _nw39[(new BigNumber(28)).toNumber()] = _259_v69;
          _308_v111 = _nw39;
          let _index39 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_308_v111).length));
          (_308_v111)[_index39] = _259_v69;
          let _index40 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_308_v111).length));
          (_308_v111)[_index40] = _259_v69;
          _179_v12 = ((_164_v0) ? (_179_v12) : (_179_v12));
          let _index41 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_171_v7).length));
          (_171_v7)[_index41] = !(_164_v0) || (_164_v0);
          let _index42 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_171_v7).length));
          (_171_v7)[_index42] = _164_v0;
        }
        let _309_v112;
        _309_v112 = _module.D0.create_DC0((_255_v66)[_module.__default.safeIndex(new BigNumber(580), new BigNumber((_255_v66).length))], (_178_v11).update(_174_v8, (_252_v63).f18), _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_167_v3), (_259_v69).f18));
        let _310_v113;
        _310_v113 = _module.D0.create_DC1(_309_v112);
        let _311_v114;
        let _312_v115;
        let _313_v116;
        let _314_v117;
        let _out36;
        let _out37;
        let _out38;
        let _out39;
        let _outcollector9 = _module.__default.m0(_module.D0.create_DC1(_310_v113), ((_259_v69).f18).minus((_259_v69).f18), _175_globalState);
        _out36 = _outcollector9[0];
        _out37 = _outcollector9[1];
        _out38 = _outcollector9[2];
        _out39 = _outcollector9[3];
        _311_v114 = _out36;
        _312_v115 = _out37;
        _313_v116 = _out38;
        _314_v117 = _out39;
        _164_v0 = _164_v0;
        let _index43 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_171_v7).length));
        (_171_v7)[_index43] = _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("eqmqr"), _174_v8);
        let _index44 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_171_v7).length));
        (_171_v7)[_index44] = _314_v117;
      } else {
        let _315_v118;
        let _nw40 = new _module.C0();
        _nw40.__ctor((_252_v63).f18);
        _315_v118 = _nw40;
        (_175_globalState).f4 = (_252_v63).f18;
        let _rhs20 = (_252_v63).f18;
        let _rhs21 = _259_v69;
        let _rhs22 = !(_module.__default.fm1(_164_v0, _175_globalState));
        let _lhs18 = _175_globalState;
        let _lhs19 = _175_globalState;
        _lhs18.f4 = _rhs20;
        _259_v69 = _rhs21;
        _lhs19.f13 = _rhs22;
        let _316_v119;
        _316_v119 = _dafny.Map.Empty.slice().updateUnsafe(_281_v88,_module.__default.fm1(_164_v0, _175_globalState));
        let _rhs23 = !(_164_v0);
        let _rhs24 = _164_v0;
        let _rhs25 = _316_v119;
        let _lhs20 = _175_globalState;
        _164_v0 = _rhs23;
        _lhs20.f13 = _rhs24;
        _316_v119 = _rhs25;
        let _317_v120;
        _317_v120 = _dafny.Map.Empty.slice().updateUnsafe((_259_v69).f18,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-148)), ((_318_v88) => function (_319_i5) {
          return _318_v88;
        })(_281_v88)));
        let _320_v121;
        _320_v121 = _module.D3.create_DC9(_317_v120);
        let _321_v123;
        let _nw41 = Array((new BigNumber(7)).toNumber());
        _nw41[(_dafny.ZERO).toNumber()] = _317_v120;
        _nw41[(_dafny.ONE).toNumber()] = ((_164_v0) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(696),_174_v8)) : (_317_v120));
        _nw41[(new BigNumber(2)).toNumber()] = (_320_v121).dtor_cf15;
        _nw41[(new BigNumber(3)).toNumber()] = function () {
          let _coll7 = new _dafny.Map();
          for (const _compr_7 of _dafny.IntegerRange(new BigNumber(905), new BigNumber(714))) {
            let _322_v122 = _compr_7;
            if (((new BigNumber(905)).isLessThanOrEqualTo(_322_v122)) && ((_322_v122).isLessThan(new BigNumber(714)))) {
              _coll7.push([_module.__default.safeDivisionInt(_322_v122, new BigNumber((_258_v68).length)),_174_v8]);
            }
          }
          return _coll7;
        }();
        _nw41[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_255_v66)[_module.__default.safeIndex(new BigNumber(580), new BigNumber((_255_v66).length))],_174_v8);
        _nw41[(new BigNumber(5)).toNumber()] = _317_v120;
        _nw41[(new BigNumber(6)).toNumber()] = _317_v120;
        _321_v123 = _nw41;
        let _index45 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_321_v123).length));
        (_321_v123)[_index45] = _317_v120;
        let _index46 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_321_v123).length));
        (_321_v123)[_index46] = ((_164_v0) ? (_dafny.Map.Empty.slice().updateUnsafe((_259_v69).f18,_174_v8)) : (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_174_v8).length)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(-266)), ((_323_v88) => function (_324_i6) {
          return _323_v88;
        })(_281_v88)))));
      }
      _281_v88 = _module.__default.fm6((_257_v67).IsProperSubsetOf(_257_v67), _175_globalState);
      _177_v10 = (_177_v10).update((_164_v0) === (!(_164_v0)), (_252_v63).f18);
      if (_164_v0) {
        let _index47 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_255_v66).length));
        (_255_v66)[_index47] = _module.__default.safeModuloInt(((_164_v0) ? ((_259_v69).f18) : (new BigNumber(706))), (_252_v63).f18);
        let _rhs26 = (_252_v63).f18;
        let _rhs27 = _module.__default.fm0(_179_v12, _175_globalState);
        let _lhs21 = _175_globalState;
        let _lhs22 = _175_globalState;
        _lhs21.f4 = _rhs26;
        _lhs22.f4 = _rhs27;
        let _325_v124;
        let _nw42 = new _module.C0();
        _nw42.__ctor(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(51)), ((_326_v68) => function (_327_i7) {
          return _326_v68;
        })(_258_v68)), _module.__default.safeIndex((_255_v66)[_module.__default.safeIndex(new BigNumber(580), new BigNumber((_255_v66).length))], new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(51)), ((_328_v68) => function (_329_i7) {
          return _328_v68;
        })(_258_v68))).length)), _258_v68)).length));
        _325_v124 = _nw42;
        _178_v11 = (_178_v11).update(_174_v8, new BigNumber(759));
        (_175_globalState).f4 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt((_259_v69).f18, (_170_v6)[_module.__default.safeIndex((_325_v124).f18, new BigNumber((_170_v6).length))]), _module.__default.fm0(_179_v12, _175_globalState));
      } else {
        let _330_v125;
        _330_v125 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm7(new BigNumber((_dafny.Seq.of((((_260_v70).contains((_259_v69).f18)) ? ((_260_v70).get((_259_v69).f18)) : (_164_v0)), !(_164_v0), _164_v0)).length), _258_v68, _164_v0, _175_globalState)).cardinality()),_174_v8);
        let _331_v126;
        let _nw43 = Array((new BigNumber(19)).toNumber()).fill(_module.D2.Default());
        _331_v126 = _nw43;
        let _332_v127;
        _332_v127 = _dafny.Map.Empty.slice().updateUnsafe(_330_v125,_331_v126);
        _332_v127 = (_332_v127).update(_330_v125, _331_v126);
        let _333_v128;
        let _nw44 = new _module.C0();
        _nw44.__ctor((new BigNumber(868)).plus(new BigNumber(823)));
        _333_v128 = _nw44;
        (_175_globalState).f13 = ((true) ? (_164_v0) : (!(_164_v0)));
        let _334_v129;
        _334_v129 = _dafny.Map.Empty.slice().updateUnsafe(_164_v0,_255_v66);
        let _index48 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_255_v66).length));
        (_255_v66)[_index48] = (((_dafny.ZERO).minus(_167_v3)).plus((_252_v63).f18)).minus(new BigNumber((_334_v129).length));
        (_175_globalState).f0 = (_167_v3).plus((new BigNumber(326)).plus(new BigNumber((_168_v4).cardinality())));
      }
      process.stdout.write(_dafny.toString(_164_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_165_v1).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_167_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v4).equals(_dafny.MultiSet.fromElements(new BigNumber(639), new BigNumber(-408)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v5).equals(_dafny.MultiSet.fromElements(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_170_v6, _dafny.Seq.of(new BigNumber(2), new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v7)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v7)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v7)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v7)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v7)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v7)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v7)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v7)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v7)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v7)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v7)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_174_v8).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_175_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_175_globalState).f1).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_175_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_175_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_175_globalState.f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_175_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_175_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_175_globalState).f8, _dafny.Seq.of(new BigNumber(2), new BigNumber(639), new BigNumber(2), new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_175_globalState).f9).equals(_dafny.MultiSet.fromElements(new BigNumber(639), new BigNumber(-408)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_175_globalState.f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_175_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_175_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_175_globalState.f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_175_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_175_globalState.f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_175_globalState).f16)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_175_globalState).f16)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_175_globalState).f16)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_175_globalState).f16)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_175_globalState).f16)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_175_globalState).f16)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_175_globalState).f16)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_175_globalState).f16)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_175_globalState).f16)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_175_globalState).f16)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_175_globalState).f16)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_175_globalState).f17).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_176_v9).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(639),new BigNumber(639)).updateUnsafe(new BigNumber(2),new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_177_v10).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(639)).updateUnsafe(false,new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_v11).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("mgx"),new BigNumber(759)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_179_v12).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v12).dtor_cf1).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("mgx"),_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v12).dtor_cf2).equals(_dafny.MultiSet.fromElements(new BigNumber(639), new BigNumber(-408)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_251_v62, _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(2), new BigNumber(2)), _dafny.MultiSet.fromElements(new BigNumber(639), new BigNumber(-408)), _dafny.MultiSet.fromElements(new BigNumber(639), new BigNumber(-408)), _dafny.MultiSet.fromElements(new BigNumber(639), new BigNumber(-408))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_252_v63).f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_253_v64).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_254_v65).dtor_cf4).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(639)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_255_v66)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_255_v66)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_255_v66)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_255_v66)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_255_v66)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_255_v66)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_255_v66)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_255_v66)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_255_v66)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_255_v66)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_255_v66)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_255_v66)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_255_v66)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_255_v66)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_257_v67).equals(_dafny.Set.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_258_v68, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_259_v69).f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_260_v70).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(2),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_281_v88));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v89).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(639), new BigNumber(639), new BigNumber(639), _dafny.ZERO, new BigNumber(-242)),new _dafny.CodePoint('j'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_283_v90).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(639), new BigNumber(639), new BigNumber(639), _dafny.ZERO, new BigNumber(-242)),new _dafny.CodePoint('j'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_284_v91)[new BigNumber(9)], _dafny.Seq.of(new BigNumber(2), new BigNumber(2), new BigNumber(2), new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0(cf0, cf1, cf2) {
      let $dt = new D0(0);
      $dt.cf0 = cf0;
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC1(cf3) {
      let $dt = new D0(1);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ", " + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf3) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf0, other.cf0) && _dafny.areEqual(this.cf1, other.cf1) && _dafny.areEqual(this.cf2, other.cf2);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf3, other.cf3);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC0(_dafny.ZERO, _dafny.Map.Empty, _dafny.MultiSet.Empty);
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
    static create_DC3(cf5) {
      let $dt = new D1(0);
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC4(cf6, cf7, cf8, cf9, cf10) {
      let $dt = new D1(1);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC2(cf4) {
      let $dt = new D1(2);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC5(cf11) {
      let $dt = new D1(3);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get is_DC5() { return this.$tag === 3; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + this.cf6.toVerbatimString(true) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + this.cf9.toVerbatimString(true) + ", " + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6) && this.cf7 === other.cf7 && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9) && _dafny.areEqual(this.cf10, other.cf10);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(_dafny.Set.Empty);
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
    static create_DC7(cf13) {
      let $dt = new D2(0);
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC6(cf12) {
      let $dt = new D2(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC8(cf14) {
      let $dt = new D2(2);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf14) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf13 === other.cf13;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf12 === other.cf12;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf14, other.cf14);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7([]);
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
    static create_DC9(cf15) {
      let $dt = new D3(1);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ", " + this.cf18.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf15) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf16 === other.cf16 && this.cf17 === other.cf17 && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf15, other.cf15);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(false, [], _dafny.Seq.UnicodeFromString(""));
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
    static create_DC11(cf19) {
      let $dt = new D4(1);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC12";
      } else if (this.$tag === 1) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf19) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf19, other.cf19);
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

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = _dafny.ZERO;
      this.f4 = _dafny.ZERO;
      this.f10 = _dafny.ZERO;
      this.f13 = false;
      this.f15 = _dafny.ZERO;
      this._f1 = _dafny.Set.Empty;
      this._f2 = false;
      this._f3 = false;
      this._f5 = _dafny.ZERO;
      this._f6 = false;
      this._f7 = [];
      this._f8 = _dafny.Seq.of();
      this._f9 = _dafny.MultiSet.Empty;
      this._f11 = _dafny.ZERO;
      this._f12 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f14 = _dafny.ZERO;
      this._f16 = [];
      this._f17 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this).f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this)._f14 = f14;
      (_this).f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      return;
    }
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f18 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f18) {
      let _this = this;
      (_this)._f18 = f18;
      return;
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
